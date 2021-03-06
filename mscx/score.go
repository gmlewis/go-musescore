/*
Copyright © 2022 Glenn M. Lewis

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package mscx

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"strconv"
)

// ScoreZip represents a MuseScore 3 score in `mscz` (zip'd) format.
type ScoreZip struct {
	// MetaInf
	MuseScore MuseScore `xml:"museScore"`
	// Thumbnails
}

var (
	xmlEndingsToShorten = []string{
		"></bracket>",
		"></controller>",
		"></endSpanner>",
		"></pos>",
		"></program>",
		"></size>",
	}
	xmlEndingsToSplitLines = []string{
		"</Slur>",
		"</System>",
		"</Zerberus>",
	}
)

type UnhandledError struct {
	Type   string
	Name   string
	Offset int64
}

func (u *UnhandledError) Error() string {
	return fmt.Sprintf("unhandled %v: %v at byte offset %v", u.Type, u.Name, u.Offset)
}

// XML renders the embedded MuseScore to XML format.
func (s *ScoreZip) XML() ([]byte, error) {
	b, err := xml.MarshalIndent(s.MuseScore, "", "  ")
	if err != nil {
		return nil, err
	}

	result := append([]byte(xml.Header+"<m"), b[2:]...)
	for _, ending := range xmlEndingsToShorten {
		result = bytes.ReplaceAll(result, []byte(ending), []byte("/>"))
	}
	for _, ending := range xmlEndingsToSplitLines {
		result = bytes.ReplaceAll(result, []byte(ending), []byte("\n"+ending))
	}
	result = bytes.ReplaceAll(result, []byte("&#xA;"), []byte("\n"))
	result = bytes.ReplaceAll(result, []byte("&#39;"), []byte("'"))
	result = bytes.ReplaceAll(result, []byte("  </"), []byte("    </"))
	result = bytes.ReplaceAll(result, []byte("</MuseScore>"), []byte("  </museScore>\n"))
	return result, nil
}

// MuseScore represents MuseScore 3 data in XML.
type MuseScore struct {
	Version string `xml:"version,attr"`

	ProgramVersion  string `xml:"programVersion"`
	ProgramRevision string `xml:"programRevision"`
	Score           Score  `xml:"Score"`
}

// Score represents the XML data of the same name.
type Score struct {
	LayerTag        LayerTag      `xml:"LayerTag"`
	CurrentLayer    int           `xml:"currentLayer"`
	Synthesizer     *Synthesizer  `xml:"Synthesizer"`
	Division        int           `xml:"Division"`
	Style           *Style        `xml:"Style"`
	ShowInvisible   int           `xml:"showInvisible"`
	ShowUnprintable int           `xml:"showUnprintable"`
	ShowFrames      int           `xml:"showFrames"`
	ShowMargins     int           `xml:"showMargins"`
	MetaTags        []*MetaTag    `xml:"metaTag"`
	PageList        *PageList     `xml:"PageList"`
	Part            []*Part       `xml:"Part"`
	Staffs          []*ScoreStaff `xml:"Staff"`
}

// LayerTag represents the XML data of the same name.
type LayerTag struct {
	ID  string `xml:"id,attr"`
	Tag string `xml:"tag,attr"`
}

// Style represents the XML data of the same name.
type Style struct {
	ConcertPitch         int         `xml:"concertPitch,omitempty"`
	PageLayout           *PageLayout `xml:"page-layout"`
	PageWidth            float64     `xml:"pageWidth,omitempty"`
	PageHeight           float64     `xml:"pageHeight,omitempty"`
	PagePrintableWidth   float64     `xml:"pagePrintableWidth,omitempty"`
	UseStandardNoteNames int         `xml:"useStandardNoteNames,omitempty"`
	Spatium              float64     `xml:"Spatium"`
}

// MetaTag represents the XML data of the same name.
type MetaTag struct {
	Name string `xml:"name,attr"`

	Text string `xml:",chardata"`
}

// Part represents the XML data of the same name.
type Part struct {
	Staff      []*PartStaff `xml:"Staff"`
	Show       string       `xml:"show,omitempty"`
	TrackName  string       `xml:"trackName"`
	Instrument *Instrument  `xml:"Instrument"`
}

// PartStaff represents the XML data of the same name.
type PartStaff struct {
	ID string `xml:"id,attr"`

	StaffType     StaffType `xml:"StaffType"`
	StaffElements []any
}

func (p *PartStaff) MarshalXML(encoder *xml.Encoder, start xml.StartElement) error {
	se := xml.StartElement{
		Name: xml.Name{Local: "Staff"},
		Attr: []xml.Attr{
			{
				Name:  xml.Name{Local: "id"},
				Value: p.ID,
			},
		},
	}
	if err := encoder.EncodeToken(se); err != nil {
		return fmt.Errorf("PartStaff.MarshalXML: %w", err)
	}

	if err := encoder.Encode(p.StaffType); err != nil {
		return fmt.Errorf("PartStaff.MarshalXML: %w", err)
	}

	for _, el := range p.StaffElements {
		if err := encoder.Encode(el); err != nil {
			return fmt.Errorf("PartStaff.MarshalXML: %w", err)
		}
	}

	if err := encoder.EncodeToken(xml.EndElement{Name: xml.Name{Local: "Staff"}}); err != nil {
		return fmt.Errorf("PartStaff.MarshalXML: %w", err)
	}

	return nil
}

// Implements encoding.xml.Unmarshaler interface
func (p *PartStaff) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	for _, attr := range start.Attr {
		switch attr.Name.Local {
		case "id":
			p.ID = attr.Value
		default:
			err := &UnhandledError{Type: "attr", Name: attr.Name.Local, Offset: decoder.InputOffset()}
			return fmt.Errorf("PartStaff.UnmarshalXML: %w", err)
		}
	}

	for {
		token, err := decoder.Token()
		if err != nil {
			return fmt.Errorf("PartStaff.UnmarshalXML: %w", err)
		}

		switch tok := token.(type) {
		case xml.StartElement:
			switch tok.Name.Local {
			case "StaffType":
				if err = decoder.DecodeElement(&p.StaffType, &tok); err != nil {
					return fmt.Errorf("PartStaff.UnmarshalXML: %w", err)
				}
			case "ID":
				if err = decoder.DecodeElement(&p.ID, &tok); err != nil {
					return fmt.Errorf("PartStaff.UnmarshalXML: %w", err)
				}
			case "bracket":
				el := &Bracket{}
				if err = decoder.DecodeElement(el, &tok); err != nil {
					return fmt.Errorf("PartStaff.UnmarshalXML: %w", err)
				}
				p.StaffElements = append(p.StaffElements, el)
			case "barLineSpan":
				el := BarLineSpan(0)
				if err = decoder.DecodeElement(&el, &tok); err != nil {
					return fmt.Errorf("PartStaff.UnmarshalXML: %w", err)
				}
				p.StaffElements = append(p.StaffElements, el)
			case "defaultClef":
				el := &DefaultClef{}
				if err = decoder.DecodeElement(el, &tok); err != nil {
					return fmt.Errorf("PartStaff.UnmarshalXML: %w", err)
				}
				p.StaffElements = append(p.StaffElements, el)
			case "defaultConcertClef":
				el := &DefaultClef{Type: "Concert"}
				if err = decoder.DecodeElement(el, &tok); err != nil {
					return fmt.Errorf("PartStaff.UnmarshalXML: %w", err)
				}
				p.StaffElements = append(p.StaffElements, el)
			case "defaultTransposingClef":
				el := &DefaultClef{Type: "Transposing"}
				if err = decoder.DecodeElement(el, &tok); err != nil {
					return fmt.Errorf("PartStaff.UnmarshalXML: %w", err)
				}
				p.StaffElements = append(p.StaffElements, el)
			case "small":
				el := StaffSmall(0)
				if err = decoder.DecodeElement(&el, &tok); err != nil {
					return fmt.Errorf("PartStaff.UnmarshalXML: %w", err)
				}
				p.StaffElements = append(p.StaffElements, el)
			case "distOffset":
				el := DistOffset(0)
				if err = decoder.DecodeElement(&el, &tok); err != nil {
					return fmt.Errorf("PartStaff.UnmarshalXML: %w", err)
				}
				p.StaffElements = append(p.StaffElements, el)
			default:
				err := &UnhandledError{Type: "token", Name: tok.Name.Local, Offset: decoder.InputOffset()}
				return fmt.Errorf("PartStaff.UnmarshalXML: %w", err)
			}

		case xml.EndElement:
			return nil
		}
	}
}

type BarLineSpan int

func (b BarLineSpan) MarshalXML(encoder *xml.Encoder, start xml.StartElement) error {
	se := xml.StartElement{
		Name: xml.Name{Local: "barLineSpan"},
	}
	if err := encoder.EncodeToken(se); err != nil {
		return fmt.Errorf("BarLineSpan.MarshalXML: %w", err)
	}

	if err := encoder.EncodeToken(xml.CharData(fmt.Sprintf("%v", b))); err != nil {
		return fmt.Errorf("BarLineSpan.MarshalXML: %w", err)
	}

	if err := encoder.EncodeToken(xml.EndElement{Name: xml.Name{Local: "barLineSpan"}}); err != nil {
		return fmt.Errorf("BarLineSpan.MarshalXML: %w", err)
	}

	return nil
}

type DistOffset float64

func (d DistOffset) MarshalXML(encoder *xml.Encoder, start xml.StartElement) error {
	se := xml.StartElement{
		Name: xml.Name{Local: "distOffset"},
	}
	if err := encoder.EncodeToken(se); err != nil {
		return fmt.Errorf("DistOffset.MarshalXML: %w", err)
	}

	if err := encoder.EncodeToken(xml.CharData(fmt.Sprintf("%v", d))); err != nil {
		return fmt.Errorf("DistOffset.MarshalXML: %w", err)
	}

	if err := encoder.EncodeToken(xml.EndElement{Name: xml.Name{Local: "distOffset"}}); err != nil {
		return fmt.Errorf("DistOffset.MarshalXML: %w", err)
	}

	return nil
}

type StaffSmall int

func (s StaffSmall) MarshalXML(encoder *xml.Encoder, start xml.StartElement) error {
	se := xml.StartElement{
		Name: xml.Name{Local: "small"},
	}
	if err := encoder.EncodeToken(se); err != nil {
		return fmt.Errorf("StaffSmall.MarshalXML: %w", err)
	}

	if err := encoder.EncodeToken(xml.CharData(fmt.Sprintf("%v", s))); err != nil {
		return fmt.Errorf("StaffSmall.MarshalXML: %w", err)
	}

	if err := encoder.EncodeToken(xml.EndElement{Name: xml.Name{Local: "small"}}); err != nil {
		return fmt.Errorf("StaffSmall.MarshalXML: %w", err)
	}

	return nil
}

type DefaultClef struct {
	Type  string // One of: "", "Concert", "Transposing"
	Value string `xml:",chardata"`
}

func (d *DefaultClef) MarshalXML(encoder *xml.Encoder, start xml.StartElement) error {
	local := fmt.Sprintf("default%vClef", d.Type)
	se := xml.StartElement{
		Name: xml.Name{Local: local},
	}
	if err := encoder.EncodeToken(se); err != nil {
		return fmt.Errorf("DefaultClef.MarshalXML: %w", err)
	}

	if err := encoder.EncodeToken(xml.CharData(d.Value)); err != nil {
		return fmt.Errorf("DefaultClef.MarshalXML: %w", err)
	}

	if err := encoder.EncodeToken(xml.EndElement{Name: xml.Name{Local: local}}); err != nil {
		return fmt.Errorf("DefaultClef.MarshalXML: %w", err)
	}

	return nil
}

// Bracket represents the XML data of the same name.
type Bracket struct {
	Type int    `xml:"type,attr"`
	Span int    `xml:"span,attr"`
	Col  string `xml:"col,attr,omitempty"`
}

func (b *Bracket) MarshalXML(encoder *xml.Encoder, start xml.StartElement) error {
	se := xml.StartElement{
		Name: xml.Name{Local: "bracket"},
		Attr: []xml.Attr{
			{
				Name:  xml.Name{Local: "type"},
				Value: fmt.Sprintf("%v", b.Type),
			},
			{
				Name:  xml.Name{Local: "span"},
				Value: fmt.Sprintf("%v", b.Span),
			},
		},
	}
	if b.Col != "" {
		se.Attr = append(se.Attr, xml.Attr{
			Name:  xml.Name{Local: "col"},
			Value: b.Col,
		})
	}

	if err := encoder.EncodeToken(se); err != nil {
		return fmt.Errorf("Bracket.MarshalXML: %w", err)
	}

	if err := encoder.EncodeToken(xml.EndElement{Name: xml.Name{Local: "bracket"}}); err != nil {
		return fmt.Errorf("Bracket.MarshalXML: %w", err)
	}

	return nil
}

// StaffType represents the XML data of the same name.
type StaffType struct {
	Group string `xml:"group,attr"`

	Name   string `xml:"name"`
	Lines  int    `xml:"lines,omitempty"`
	KeySig *int   `xml:"keysig,omitempty"`
}

// ScoreStaff represents the XML data of the same name.
type ScoreStaff struct {
	ID string `xml:"id,attr"`

	VBox    *VBox      `xml:"VBox"`
	Measure []*Measure `xml:"Measure"`
}

// Measure represents the XML data of the same name.
type Measure struct {
	Len    string `xml:"len,attr,omitempty"`
	Number int    `xml:"number,attr,omitempty"`

	Irregular   int      `xml:"irregular,omitempty"`
	Voice       []*Voice `xml:"voice"`
	StartRepeat string   `xml:"startRepeat,omitempty"`
	EndRepeat   string   `xml:"endRepeat,omitempty"`

	// older versions
	KeySig        *KeySig  `xml:"KeySig"`
	TimeSig       *TimeSig `xml:"TimeSig"`
	Tempo         *Tempo   `xml:"Tempo"`
	TimedElements []any
}

// Implements encoding.xml.Marshaler interface
func (m *Measure) MarshalXML(encoder *xml.Encoder, start xml.StartElement) error {
	se := xml.StartElement{
		Name: xml.Name{Local: "Measure"},
	}
	if m.Len != "" {
		se.Attr = append(se.Attr, xml.Attr{
			Name:  xml.Name{Local: "len"},
			Value: m.Len,
		})
	}
	if m.Number != 0 {
		se.Attr = append(se.Attr, xml.Attr{
			Name:  xml.Name{Local: "number"},
			Value: fmt.Sprintf("%v", m.Number),
		})
	}
	if err := encoder.EncodeToken(se); err != nil {
		return fmt.Errorf("Measure.MarshalXML: %w", err)
	}

	if m.Irregular != 0 {
		irregularEl := xml.StartElement{Name: xml.Name{Local: "irregular"}}
		if err := encoder.EncodeElement(m.Irregular, irregularEl); err != nil {
			return fmt.Errorf("Measure.MarshalXML: %w", err)
		}
	}

	if m.Voice != nil {
		if err := encoder.Encode(m.Voice); err != nil {
			return fmt.Errorf("Measure.MarshalXML: %w", err)
		}
	}

	if m.StartRepeat != "" {
		if err := encoder.Encode(m.StartRepeat); err != nil {
			return fmt.Errorf("Measure.MarshalXML: %w", err)
		}
	}

	if m.EndRepeat != "" {
		if err := encoder.Encode(m.EndRepeat); err != nil {
			return fmt.Errorf("Measure.MarshalXML: %w", err)
		}
	}

	if m.KeySig != nil {
		if err := encoder.Encode(m.KeySig); err != nil {
			return fmt.Errorf("Measure.MarshalXML: %w", err)
		}
	}

	if m.TimeSig != nil {
		if err := encoder.Encode(m.TimeSig); err != nil {
			return fmt.Errorf("Measure.MarshalXML: %w", err)
		}
	}

	if m.Tempo != nil {
		if err := encoder.Encode(m.Tempo); err != nil {
			return fmt.Errorf("Measure.MarshalXML: %w", err)
		}
	}

	for _, el := range m.TimedElements {
		if err := encoder.Encode(el); err != nil {
			return fmt.Errorf("Measure.MarshalXML: %w", err)
		}
	}

	if err := encoder.EncodeToken(xml.EndElement{Name: xml.Name{Local: "Measure"}}); err != nil {
		return fmt.Errorf("Measure.MarshalXML: %w", err)
	}

	return nil
}

// Implements encoding.xml.Unmarshaler interface
func (m *Measure) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	for _, attr := range start.Attr {
		switch attr.Name.Local {
		case "len":
			m.Len = attr.Value
		case "number":
			v, err := strconv.Atoi(attr.Value)
			if err != nil {
				return fmt.Errorf("Measure.UnmarshalXML: Atoi: %w", err)
			}
			m.Number = v
		default:
			err := &UnhandledError{Type: "attr", Name: attr.Name.Local, Offset: decoder.InputOffset()}
			return fmt.Errorf("Measure.UnmarshalXML: %w", err)
		}
	}

	for {
		token, err := decoder.Token()
		if err != nil {
			return fmt.Errorf("Measure.UnmarshalXML: %w", err)
		}

		switch tok := token.(type) {
		case xml.StartElement:
			switch tok.Name.Local {
			case "irregular":
				if err = decoder.DecodeElement(&m.Irregular, &tok); err != nil {
					return fmt.Errorf("Measure.UnmarshalXML: %w", err)
				}
			case "voice":
				if err = decoder.DecodeElement(&m.Voice, &tok); err != nil {
					return fmt.Errorf("Measure.UnmarshalXML: %w", err)
				}
			case "KeySig":
				if err = decoder.DecodeElement(&m.KeySig, &tok); err != nil {
					return fmt.Errorf("Measure.UnmarshalXML: %w", err)
				}
			case "TimeSig":
				if err = decoder.DecodeElement(&m.TimeSig, &tok); err != nil {
					return fmt.Errorf("Measure.UnmarshalXML: %w", err)
				}
			case "Tempo":
				if err = decoder.DecodeElement(&m.Tempo, &tok); err != nil {
					return fmt.Errorf("Measure.UnmarshalXML: %w", err)
				}
			case "tick":
				el := Tick(0)
				if err = decoder.DecodeElement(&el, &tok); err != nil {
					return fmt.Errorf("Measure.UnmarshalXML: %w", err)
				}
				m.TimedElements = append(m.TimedElements, el)
			case "Clef":
				el := &Clef{}
				if err = decoder.DecodeElement(el, &tok); err != nil {
					return fmt.Errorf("Measure.UnmarshalXML: %w", err)
				}
				m.TimedElements = append(m.TimedElements, el)
			case "Dynamic":
				el := &Dynamic{}
				if err = decoder.DecodeElement(el, &tok); err != nil {
					return fmt.Errorf("Measure.UnmarshalXML: %w", err)
				}
				m.TimedElements = append(m.TimedElements, el)
			case "endSpanner":
				el := &EndSpanner{}
				if err = decoder.DecodeElement(el, &tok); err != nil {
					return fmt.Errorf("Measure.UnmarshalXML: %w", err)
				}
				m.TimedElements = append(m.TimedElements, el)
			case "HairPin":
				el := &HairPin{}
				if err = decoder.DecodeElement(el, &tok); err != nil {
					return fmt.Errorf("Measure.UnmarshalXML: %w", err)
				}
				m.TimedElements = append(m.TimedElements, el)
			case "LayoutBreak":
				el := &LayoutBreak{}
				if err = decoder.DecodeElement(el, &tok); err != nil {
					return fmt.Errorf("Measure.UnmarshalXML: %w", err)
				}
				m.TimedElements = append(m.TimedElements, el)
			case "StaffText":
				el := &StaffText{}
				if err = decoder.DecodeElement(el, &tok); err != nil {
					return fmt.Errorf("Measure.UnmarshalXML: %w", err)
				}
				m.TimedElements = append(m.TimedElements, el)
			case "Chord":
				el := &Chord{}
				if err = decoder.DecodeElement(el, &tok); err != nil {
					return fmt.Errorf("Measure.UnmarshalXML: %w", err)
				}
				m.TimedElements = append(m.TimedElements, el)
			case "Rest":
				el := &Rest{}
				if err = decoder.DecodeElement(el, &tok); err != nil {
					return fmt.Errorf("Measure.UnmarshalXML: %w", err)
				}
				m.TimedElements = append(m.TimedElements, el)
			case "BarLine":
				el := &BarLine{}
				if err = decoder.DecodeElement(el, &tok); err != nil {
					return fmt.Errorf("Measure.UnmarshalXML: %w", err)
				}
				m.TimedElements = append(m.TimedElements, el)
			default:
				err := &UnhandledError{Type: "token", Name: tok.Name.Local, Offset: decoder.InputOffset()}
				return fmt.Errorf("Measure.UnmarshalXML: %w", err)
			}

		case xml.EndElement:
			return nil
		}
	}
}

type EndSpanner struct {
	ID int `xml:"id,attr"`
}

type HairPin struct {
	ID int `xml:"id,attr"`

	Subtype      string       `xml:"subtype"`
	VeloChange   int          `xml:"veloChange,omitempty"`
	Segment      *Segment     `xml:"Segment"`
	BeginText    *TextElement `xml:"beginText"`
	ContinueText *TextElement `xml:"continueText"`
}

type Segment struct {
	Subtype string   `xml:"subtype"`
	Off2    *TextPos `xml:"off2"`
	Visible int      `xml:"visible"`
}

type Tick int64

type Dynamic struct {
	Subtype  string `xml:"subtype"`
	Velocity int    `xml:"velocity,omitempty"`
}

type LayoutBreak struct {
	Subtype string `xml:"subtype"`
}

type Tempo struct {
	Tempo      float64  `xml:"tempo"`
	FollowText int      `xml:"followText,omitempty"`
	Pos        *TextPos `xml:"pos"`
	Visible    int      `xml:"visible"`
	Text       []byte   `xml:"text"`
}

type StaffText struct {
	Pos   *TextPos `xml:"pos"`
	Style string   `xml:"style,omitempty"`
	Text  []byte   `xml:"text"`
}

// type TempoText struct {
// 	Text []byte `xml:",chardata"`
// }

// Instrument represents the XML data of the same name.
type Instrument struct {
	LongName           string `xml:"longName,omitempty"`
	ShortName          string `xml:"shortName,omitempty"`
	TrackName          string `xml:"trackName"`
	MinPitchP          string `xml:"minPitchP,omitempty"`
	MaxPitchP          string `xml:"maxPitchP,omitempty"`
	MinPitchA          string `xml:"minPitchA,omitempty"`
	MaxPitchA          string `xml:"maxPitchA,omitempty"`
	TransposeDiatonic  string `xml:"transposeDiatonic,omitempty"`
	TransposeChromatic string `xml:"transposeChromatic,omitempty"`
	InstrumentID       string `xml:"instrumentId"`

	UseDrumset int     `xml:"useDrumset,omitempty"`
	Drum       []*Drum `xml:"Drum"`

	Clef *Clef `xml:"clef"`

	StringData *StringData `xml:"StringData"`

	Articulation []*ArticulationElement `xml:"Articulation"`
	Channel      []*Channel             `xml:"Channel"`
}

type StringData struct {
	Frets  int   `xml:"frets"`
	String []int `xml:"string"`
}

type Drum struct {
	Pitch int `xml:"pitch,attr"`

	Head     int    `xml:"head"`
	Line     int    `xml:"line"`
	Voice    int    `xml:"voice"`
	Name     string `xml:"name"`
	Stem     int    `xml:"stem"`
	Shortcut string `xml:"shortcut,omitempty"`
}

// Clef represents the XML data of the same name.
type Clef struct {
	Staff string `xml:"staff,attr,omitempty"`

	ConcertClefType     string `xml:"concertClefType,omitempty"`
	TransposingClefType string `xml:"transposingClefType,omitempty"`

	Text string `xml:",chardata"`
}

// ArticulationElement represents the XML data of the same name.
type ArticulationElement struct {
	Name string `xml:"name,attr,omitempty"`

	Velocity string `xml:"velocity"`
	GateTime string `xml:"gateTime"`
}

// Channel represents the XML data of the same name.
type Channel struct {
	Name string `xml:"name,attr"`

	ChannelElements []any
	// Controller []*Controller `xml:"controller"`
	// Program    Program       `xml:"program"`
	Synti string `xml:"synti"`
	Mute  int    `xml:"mute,omitempty"`
}

func (c *Channel) MarshalXML(encoder *xml.Encoder, start xml.StartElement) error {
	se := xml.StartElement{
		Name: xml.Name{Local: "Channel"},
	}
	if c.Name != "" {
		se.Attr = append(se.Attr, xml.Attr{
			Name:  xml.Name{Local: "name"},
			Value: c.Name,
		})
	}
	if err := encoder.EncodeToken(se); err != nil {
		return fmt.Errorf("Channel.MarshalXML: %w", err)
	}

	for _, el := range c.ChannelElements {
		if err := encoder.Encode(el); err != nil {
			return fmt.Errorf("Channel.MarshalXML: %w", err)
		}
	}

	syntiEl := xml.StartElement{Name: xml.Name{Local: "synti"}}
	if err := encoder.EncodeElement(c.Synti, syntiEl); err != nil {
		return fmt.Errorf("Channel.MarshalXML: %w", err)
	}

	if c.Mute != 0 {
		muteEl := xml.StartElement{Name: xml.Name{Local: "mute"}}
		if err := encoder.EncodeElement(c.Mute, muteEl); err != nil {
			return fmt.Errorf("Channel.MarshalXML: %w", err)
		}
	}

	if err := encoder.EncodeToken(xml.EndElement{Name: xml.Name{Local: "Channel"}}); err != nil {
		return fmt.Errorf("Channel.MarshalXML: %w", err)
	}

	return nil
}

// Implements encoding.xml.Unmarshaler interface
func (c *Channel) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	for _, attr := range start.Attr {
		switch attr.Name.Local {
		case "name":
			c.Name = attr.Value
		default:
			err := &UnhandledError{Type: "attr", Name: attr.Name.Local, Offset: decoder.InputOffset()}
			return fmt.Errorf("Channel.UnmarshalXML: %w", err)
		}
	}

	for {
		token, err := decoder.Token()
		if err != nil {
			return fmt.Errorf("Channel.UnmarshalXML: %w", err)
		}

		switch tok := token.(type) {
		case xml.StartElement:
			switch tok.Name.Local {
			case "mute":
				if err = decoder.DecodeElement(&c.Mute, &tok); err != nil {
					return fmt.Errorf("Channel.UnmarshalXML: %w", err)
				}
			case "synti":
				if err = decoder.DecodeElement(&c.Synti, &tok); err != nil {
					return fmt.Errorf("Channel.UnmarshalXML: %w", err)
				}
			case "controller":
				el := &Controller{}
				if err = decoder.DecodeElement(el, &tok); err != nil {
					return fmt.Errorf("Channel.UnmarshalXML: %w", err)
				}
				c.ChannelElements = append(c.ChannelElements, el)
			case "program":
				var el Program
				if err = decoder.DecodeElement(&el, &tok); err != nil {
					return fmt.Errorf("Channel.UnmarshalXML: %w", err)
				}
				c.ChannelElements = append(c.ChannelElements, el)
			default:
				err := &UnhandledError{Type: "token", Name: tok.Name.Local, Offset: decoder.InputOffset()}
				return fmt.Errorf("Channel.UnmarshalXML: %w", err)
			}

		case xml.EndElement:
			return nil
		}
	}
}

// Program represents the XML data of the same name.
type Program struct {
	Value string `xml:"value,attr"`
}

func (p Program) MarshalXML(encoder *xml.Encoder, start xml.StartElement) error {
	se := xml.StartElement{
		Name: xml.Name{Local: "program"},
		Attr: []xml.Attr{
			{
				Name:  xml.Name{Local: "value"},
				Value: p.Value,
			},
		},
	}
	if err := encoder.EncodeToken(se); err != nil {
		return fmt.Errorf("Program.MarshalXML: %w", err)
	}

	if err := encoder.EncodeToken(xml.EndElement{Name: xml.Name{Local: "program"}}); err != nil {
		return fmt.Errorf("Program.MarshalXML: %w", err)
	}

	return nil
}

type Controller struct {
	Ctrl  int `xml:"ctrl,attr"`
	Value int `xml:"value,attr"`
}

func (c *Controller) MarshalXML(encoder *xml.Encoder, start xml.StartElement) error {
	se := xml.StartElement{
		Name: xml.Name{Local: "controller"},
		Attr: []xml.Attr{
			{
				Name:  xml.Name{Local: "ctrl"},
				Value: fmt.Sprintf("%v", c.Ctrl),
			},
			{
				Name:  xml.Name{Local: "value"},
				Value: fmt.Sprintf("%v", c.Value),
			},
		},
	}
	if err := encoder.EncodeToken(se); err != nil {
		return fmt.Errorf("Bracket.MarshalXML: %w", err)
	}

	if err := encoder.EncodeToken(xml.EndElement{Name: xml.Name{Local: "controller"}}); err != nil {
		return fmt.Errorf("Bracket.MarshalXML: %w", err)
	}

	return nil
}

type Voice struct {
	KeySig        *KeySig  `xml:"KeySig"`
	TimeSig       *TimeSig `xml:"TimeSig"`
	TimedElements []any
	// 	// 	Dynamic *Dynamic `xml:"Dynamic,omitempty"`
	// 	// 	Tempo *Tempo   `xml:"Tempo,omitempty"`
	// 	Chord []*Chord `xml:"Chord"`
	// 	// Tuplet    *TupletClass `xml:"Tuplet,omitempty"`
	// 	EndTuplet *string  `xml:"endTuplet,omitempty"`
	// 	BarLine   *BarLine `xml:"BarLine,omitempty"`
	// 	//	Spanner   *VoiceSpanner `xml:"Spanner"`
	// 	Rest    *Rest    `xml:"Rest"`
	// 	Fermata *BarLine `xml:"Fermata,omitempty"`
	// 	// 	StaffText *StaffText `xml:"StaffText,omitempty"`
}

// Implements encoding.xml.Marshaler interface
func (v *Voice) MarshalXML(encoder *xml.Encoder, start xml.StartElement) error {
	if err := encoder.EncodeToken(xml.StartElement{Name: xml.Name{Local: "voice"}}); err != nil {
		return fmt.Errorf("Voice.MarshalXML: %w", err)
	}

	if v.KeySig != nil {
		if err := encoder.Encode(v.KeySig); err != nil {
			return fmt.Errorf("Voice.MarshalXML: %w", err)
		}
	}

	if v.TimeSig != nil {
		if err := encoder.Encode(v.TimeSig); err != nil {
			return fmt.Errorf("Voice.MarshalXML: %w", err)
		}
	}

	for _, el := range v.TimedElements {
		if err := encoder.Encode(el); err != nil {
			return fmt.Errorf("Voice.MarshalXML: %w", err)
		}
	}

	if err := encoder.EncodeToken(xml.EndElement{Name: xml.Name{Local: "voice"}}); err != nil {
		return fmt.Errorf("Voice.MarshalXML: %w", err)
	}

	return nil
}

// Implements encoding.xml.Unmarshaler interface
func (v *Voice) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	for _, attr := range start.Attr {
		switch attr.Name.Local {
		default:
			err := &UnhandledError{Type: "attr", Name: attr.Name.Local, Offset: decoder.InputOffset()}
			return fmt.Errorf("Voice.UnmarshalXML: %w", err)
		}
	}

	for {
		token, err := decoder.Token()
		if err != nil {
			return fmt.Errorf("Voice.UnmarshalXML: %w", err)
		}

		switch tok := token.(type) {
		case xml.StartElement:
			switch tok.Name.Local {
			case "KeySig":
				if err = decoder.DecodeElement(&v.KeySig, &tok); err != nil {
					return fmt.Errorf("Voice.UnmarshalXML: %w", err)
				}
			case "TimeSig":
				if err = decoder.DecodeElement(&v.TimeSig, &tok); err != nil {
					return fmt.Errorf("Voice.UnmarshalXML: %w", err)
				}
			case "Chord":
				el := &Chord{}
				if err = decoder.DecodeElement(el, &tok); err != nil {
					return fmt.Errorf("Voice.UnmarshalXML: %w", err)
				}
				v.TimedElements = append(v.TimedElements, el)
			case "Rest":
				el := &Rest{}
				if err = decoder.DecodeElement(el, &tok); err != nil {
					return fmt.Errorf("Voice.UnmarshalXML: %w", err)
				}
				v.TimedElements = append(v.TimedElements, el)
			case "BarLine":
				el := &BarLine{}
				if err = decoder.DecodeElement(el, &tok); err != nil {
					return fmt.Errorf("Voice.UnmarshalXML: %w", err)
				}
				v.TimedElements = append(v.TimedElements, el)
			default:
				err := &UnhandledError{Type: "token", Name: tok.Name.Local, Offset: decoder.InputOffset()}
				return fmt.Errorf("Voice.UnmarshalXML: %w", err)
			}

		case xml.EndElement:
			return nil
		}
	}
}

type Rest struct {
	DurationType string `xml:"durationType"`
}

type BarLine struct {
	Subtype string `xml:"subtype"`
}

type KeySig struct {
	Accidental string `xml:"accidental"`
}

type TimeSig struct {
	SigN            string `xml:"sigN"`
	SigD            string `xml:"sigD"`
	ShowCourtesySig int    `xml:"showCourtesySig,omitempty"`
}

type ImageElement struct {
	Pos      *TextPos  `xml:"pos"`
	Path     string    `xml:"path"`
	LinkPath string    `xml:"linkPath"`
	Size     ImageSize `xml:"size"`
}

type ImageSize struct {
	W float64 `xml:"w,attr"`
	H float64 `xml:"h,attr"`
}

type TextElement struct {
	Pos   *TextPos  `xml:"pos"`
	Style StyleEnum `xml:"style,omitempty"`
	Text  []byte    `xml:"text"`
}

type TextPos struct {
	X float64 `xml:"x,attr"`
	Y float64 `xml:"y,attr"`
}

type VBox struct {
	Height string         `xml:"height"`
	Text   []TextElement  `xml:"Text"`
	Image  []ImageElement `xml:"Image"`
}

type StyleEnum string

const (
	Composer StyleEnum = "Composer"
	Subtitle StyleEnum = "Subtitle"
	Title    StyleEnum = "Title"
	Tuplet   StyleEnum = "Tuplet"
)

type Chord struct {
	Dots         int        `xml:"dots,omitempty"`
	DurationType string     `xml:"durationType"`
	Lyrics       []*Lyrics  `xml:"Lyrics"`
	Spanner      []*Spanner `xml:"Spanner"`
	Note         []*Note    `xml:"Note"`
}

type Lyrics struct {
	No       int    `xml:"no,omitempty"`
	Syllabic string `xml:"syllabic,omitempty"`
	TicksF   string `xml:"ticks_f,omitempty"`
	Text     string `xml:"text,omitempty"`
}

type Spanner struct {
	Type string `xml:"type,attr"`

	Slur *Slur     `xml:"Slur"`
	Next *NextPrev `xml:"next"`
	Prev *NextPrev `xml:"prev"`
}

type Slur struct {
	Up string `xml:"up,omitempty"`
}

type NextPrev struct {
	Location *Location `xml:"location"`
}

type Location struct {
	Fractions string `xml:"fractions"`
}

type Note struct {
	Pitch int `xml:"pitch"`
	TPC   int `xml:"tpc"`
}

type Synthesizer struct {
	Master   *SynthVals `xml:"master"`
	Fluid    *SynthVals `xml:"Fluid"`
	Zerberus *SynthVals `xml:"Zerberus"`
	Zita1    *SynthVals `xml:"Zita1"`
}

type SynthVals struct {
	Val []*Val `xml:"val"`
}

type Val struct {
	ID int `xml:"id,attr"`

	Text string `xml:",chardata"`
}

type PageLayout struct {
	PageHeight  float64        `xml:"page-height"`
	PageWidth   float64        `xml:"page-width"`
	PageMargins []*PageMargins `xml:"page-margins"`
}

type PageMargins struct {
	Type string `xml:"type,attr"`

	LeftMargin   float64 `xml:"left-margin"`
	RightMargin  float64 `xml:"right-margin"`
	TopMargin    float64 `xml:"top-margin"`
	BottomMargin float64 `xml:"bottom-margin"`
}

type PageList struct {
	Page []*Page `xml:"Page"`
}

type Page struct {
	System System `xml:"System"`
}

type System struct {
}
