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
)

// ScoreZip represents a MuseScore 3 score in `mscz` (zip'd) format.
type ScoreZip struct {
	// MetaInf
	MuseScore *MuseScore `xml:"museScore"`
	// Thumbnails
}

var (
	xmlEndingsToShorten = []string{
		"></bracket>",
		"></program>",
	}
)

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
	result = bytes.ReplaceAll(result, []byte("&#xA;"), []byte("\n"))
	result = bytes.ReplaceAll(result, []byte("&#39;"), []byte("'"))
	result = bytes.ReplaceAll(result, []byte("  </"), []byte("    </"))
	result = bytes.ReplaceAll(result, []byte("</Slur>"), []byte("\n</Slur>"))
	result = bytes.ReplaceAll(result, []byte("</MuseScore>"), []byte("  </museScore>\n"))
	return result, nil
}

// MuseScore represents MuseScore 3 data in XML.
type MuseScore struct {
	Version         string `xml:"version,attr"`
	ProgramVersion  string `xml:"programVersion"`
	ProgramRevision string `xml:"programRevision"`
	Score           *Score `xml:"Score"`
}

// Score represents the XML data of the same name.
type Score struct {
	LayerTag        *LayerTag     `xml:"LayerTag"`
	CurrentLayer    int           `xml:"currentLayer"`
	Division        int           `xml:"Division"`
	Style           *Style        `xml:"Style"`
	ShowInvisible   int           `xml:"showInvisible"`
	ShowUnprintable int           `xml:"showUnprintable"`
	ShowFrames      int           `xml:"showFrames"`
	ShowMargins     int           `xml:"showMargins"`
	MetaTags        []*MetaTag    `xml:"metaTag"`
	Part            *Part         `xml:"Part"`
	Staffs          []*ScoreStaff `xml:"Staff"`
}

// LayerTag represents the XML data of the same name.
type LayerTag struct {
	ID  string `xml:"id,attr"`
	Tag string `xml:"tag,attr"`
}

// Style represents the XML data of the same name.
type Style struct {
	PageWidth            float64 `xml:"pageWidth"`
	PageHeight           float64 `xml:"pageHeight"`
	PagePrintableWidth   float64 `xml:"pagePrintableWidth"`
	UseStandardNoteNames int     `xml:"useStandardNoteNames,omitempty"`
	Spatium              float64 `xml:"Spatium"`
}

// MetaTag represents the XML data of the same name.
type MetaTag struct {
	Text string `xml:",chardata"`
	Name string `xml:"name,attr"`
}

// Part represents the XML data of the same name.
type Part struct {
	Staff      []*PartStaff `xml:"Staff"`
	TrackName  string       `xml:"trackName"`
	Instrument *Instrument  `xml:"Instrument,omitempty"`
}

// Staff represents the XML data of the same name.
type PartStaff struct {
	StaffType   StaffType `xml:"StaffType"`
	Bracket     *Bracket  `xml:"bracket,omitempty"`
	BarLineSpan int       `xml:"barLineSpan,omitempty"`
	ID          string    `xml:"id,attr"`
	DefaultClef string    `xml:"defaultClef,omitempty"`
}

// Bracket represents the XML data of the same name.
type Bracket struct {
	Type string `xml:"type,attr"`
	Span string `xml:"span,attr"`
	Col  string `xml:"col,attr"`
}

// StaffType represents the XML data of the same name.
type StaffType struct {
	Name  string `xml:"name"`
	Group string `xml:"group,attr"`
}

// ScoreStaff represents the XML data of the same name.
type ScoreStaff struct {
	VBox    *VBox      `xml:"VBox,omitempty"`
	Measure []*Measure `xml:"Measure"`
	ID      string     `xml:"id,attr"`
}

// Measure represents the XML data of the same name.
type Measure struct {
	Irregular   int      `xml:"irregular,omitempty"`
	Voice       []*Voice `xml:"voice"`
	Len         string   `xml:"len,attr,omitempty"`
	StartRepeat string   `xml:"startRepeat,omitempty"`
	EndRepeat   string   `xml:"endRepeat,omitempty"`
}

// Instrument represents the XML data of the same name.
type Instrument struct {
	LongName     string                 `xml:"longName"`
	ShortName    string                 `xml:"shortName"`
	TrackName    string                 `xml:"trackName"`
	MinPitchP    string                 `xml:"minPitchP"`
	MaxPitchP    string                 `xml:"maxPitchP"`
	MinPitchA    string                 `xml:"minPitchA"`
	MaxPitchA    string                 `xml:"maxPitchA"`
	InstrumentID string                 `xml:"instrumentId"`
	Clef         Clef                   `xml:"clef"`
	Articulation []*ArticulationElement `xml:"Articulation"`
	Channel      Channel                `xml:"Channel"`
}

// Clef represents the XML data of the same name.
type Clef struct {
	Staff string `xml:"staff,attr"`
	Text  string `xml:",chardata"`
}

// ArticulationElement represents the XML data of the same name.
type ArticulationElement struct {
	Velocity string `xml:"velocity"`
	GateTime string `xml:"gateTime"`
	Name     string `xml:"name,attr,omitempty"`
}

// Channel represents the XML data of the same name.
type Channel struct {
	Program Program `xml:"program"`
	Synti   string  `xml:"synti"`
}

// Program represents the XML data of the same name.
type Program struct {
	Value string `xml:"value,attr"`
}

type Voice struct {
	KeySig        *KeySig  `xml:"KeySig,omitempty"`
	TimeSig       *TimeSig `xml:"TimeSig,omitempty"`
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
		return err
	}

	for _, attr := range start.Attr {
		switch attr.Name.Local {
		default:
			return fmt.Errorf("voice: unhandled attr: %v", attr.Name.Local)
		}
	}

	if v.KeySig != nil {
		if err := encoder.Encode(v.KeySig); err != nil {
			return err
		}
	}

	if v.TimeSig != nil {
		if err := encoder.Encode(v.TimeSig); err != nil {
			return err
		}
	}

	for _, el := range v.TimedElements {
		if err := encoder.Encode(el); err != nil {
			return err
		}
	}

	if err := encoder.EncodeToken(xml.EndElement{Name: xml.Name{Local: "voice"}}); err != nil {
		return err
	}

	return nil
}

// Implements encoding.xml.Unmarshaler interface
func (v *Voice) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	for _, attr := range start.Attr {
		switch attr.Name.Local {
		default:
			return fmt.Errorf("voice: unhandled attr: %v", attr.Name.Local)
		}
	}

	for {
		token, err := decoder.Token()
		if err != nil {
			return err
		}

		switch tok := token.(type) {
		case xml.StartElement:
			switch tok.Name.Local {
			case "KeySig":
				if err = decoder.DecodeElement(&v.KeySig, &tok); err != nil {
					return err
				}
			case "TimeSig":
				if err = decoder.DecodeElement(&v.TimeSig, &tok); err != nil {
					return err
				}
			case "Chord":
				el := &Chord{}
				if err = decoder.DecodeElement(el, &tok); err != nil {
					return err
				}
				v.TimedElements = append(v.TimedElements, el)
			case "Rest":
				el := &Rest{}
				if err = decoder.DecodeElement(el, &tok); err != nil {
					return err
				}
				v.TimedElements = append(v.TimedElements, el)
			case "BarLine":
				el := &BarLine{}
				if err = decoder.DecodeElement(el, &tok); err != nil {
					return err
				}
				v.TimedElements = append(v.TimedElements, el)
			default:
				return fmt.Errorf("voice: unhandled token: %v", tok.Name.Local)
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
	SigN string `xml:"sigN"`
	SigD string `xml:"sigD"`
}

type TextElement struct {
	Style StyleEnum `xml:"style"`
	Text  string    `xml:"text"`
}

type VBox struct {
	Height string        `xml:"height"`
	Text   []TextElement `xml:"Text"`
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
	Type string    `xml:"type,attr"`
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
