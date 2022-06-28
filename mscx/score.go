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
	xmlMoreIndent = []string{
		"  </Articulation>",
		"  </Channel>",
		"  </Chord>",
		"  </Instrument>",
		"  </KeySig>",
		"  </location>",
		"  </Lyrics>",
		"  </Measure>",
		"  </next>",
		"  </Note>",
		"  </Part>",
		"  </Spanner>",
		"  </Staff>",
		"  </StaffType>",
		"  </Style>",
		"  </Text>",
		"  </TimeSig>",
		"  </VBox>",
		"  </voice>",
	}
)

// XML renders the embedded MuseScore to XML format.
func (s *ScoreZip) XML() ([]byte, error) {
	b, err := xml.MarshalIndent(s.MuseScore, "", "  ")
	if err != nil {
		return nil, err
	}

	result := append([]byte(`<?xml version="1.0" encoding="UTF-8"?>`+"\n<m"), b[2:]...)
	for _, ending := range xmlEndingsToShorten {
		result = bytes.ReplaceAll(result, []byte(ending), []byte("/>"))
	}
	result = bytes.ReplaceAll(result, []byte("&#xA;"), []byte("\n"))
	result = bytes.ReplaceAll(result, []byte("&#39;"), []byte("'"))
	for _, indent := range xmlMoreIndent {
		result = bytes.ReplaceAll(result, []byte(indent), []byte("  "+indent))
	}
	// result = bytes.ReplaceAll(result, []byte("  </"), []byte("    </"))
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
	Irregular   int    `xml:"irregular,omitempty"`
	Voice       *Voice `xml:"voice"`
	Len         string `xml:"len,attr,omitempty"`
	StartRepeat string `xml:"startRepeat,omitempty"`
	EndRepeat   string `xml:"endRepeat,omitempty"`
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
	KeySig  *KeySig  `xml:"KeySig,omitempty"`
	TimeSig *TimeSig `xml:"TimeSig,omitempty"`
	// 	Dynamic *Dynamic `xml:"Dynamic,omitempty"`
	// 	Tempo *Tempo   `xml:"Tempo,omitempty"`
	Chord []*Chord `xml:"Chord"`
	// Tuplet    *TupletClass `xml:"Tuplet,omitempty"`
	EndTuplet *string  `xml:"endTuplet,omitempty"`
	BarLine   *BarLine `xml:"BarLine,omitempty"`
	//	Spanner   *VoiceSpanner `xml:"Spanner"`
	Rest    *Rest    `xml:"Rest"`
	Fermata *BarLine `xml:"Fermata,omitempty"`
	// 	StaffText *StaffText `xml:"StaffText,omitempty"`
}

type Rest struct {
}

type BarLine struct {
	Subtype string `xml:"subtype"`
}

// type ChordElement struct {
// 	DurationType BaseNote           `xml:"durationType"`
// 	Note         *StickyNote        `xml:"Note"`
// 	Dots         *string            `xml:"dots,omitempty"`
// 	Articulation *ArticulationUnion `xml:"Articulation"`
// 	Spanner      *ChordSpanner      `xml:"Spanner"`
// 	BeamMode     *string            `xml:"BeamMode,omitempty"`
// }
//
// type PurpleNote struct {
// 	Pitch      string       `xml:"pitch"`
// 	Tpc        string       `xml:"tpc"`
// 	Accidental *Accidental  `xml:"Accidental,omitempty"`
// 	Spanner    *NoteSpanner `xml:"Spanner,omitempty"`
// }
//
// type Accidental struct {
// 	Subtype Subtype `xml:"subtype"`
// 	Role    *string `xml:"role,omitempty"`
// }
//
// type NoteSpanner struct {
// 	Tie  *string     `xml:"Tie,omitempty"`
// 	Next *PurpleNext `xml:"next,omitempty"`
// 	Type PurpleType  `xml:"type,attr"`
// 	Prev *PurpleNext `xml:"prev,omitempty"`
// 	Slur *string     `xml:"Slur,omitempty"`
// }
//
// type PurpleNext struct {
// 	Location PurpleLocation `xml:"location"`
// }
//
// type PurpleLocation struct {
// 	Fractions Fractions `xml:"fractions"`
// }
//
// type FluffyNote struct {
// 	Pitch      string   `xml:"pitch"`
// 	Tpc        string   `xml:"tpc"`
// 	Accidental *BarLine `xml:"Accidental,omitempty"`
// }
//
// type PurpleChord struct {
// 	Dots         *string     `xml:"dots,omitempty"`
// 	DurationType string      `xml:"durationType"`
// 	Note         *IndigoNote `xml:"Note"`
// 	Articulation *BarLine    `xml:"Articulation,omitempty"`
// }
//
// type TentacledNote struct {
// 	Pitch string `xml:"pitch"`
// 	Tpc   string `xml:"tpc"`
// }
//
// type Dynamic struct {
// 	Subtype  string  `xml:"subtype"`
// 	Velocity string  `xml:"velocity"`
// 	Offset   *Offset `xml:"offset,omitempty"`
// }
//
// type Offset struct {
// 	X string `xml:"x,attr"`
// 	Y string `xml:"y,attr"`
// }

type KeySig struct {
	Accidental string `xml:"accidental"`
}

// type RESTElement struct {
// 	DurationType BaseNote  `xml:"durationType"`
// 	BeamMode     *BeamMode `xml:"BeamMode,omitempty"`
// }
//
// type PurpleSpanner struct {
// 	HairPin *BarLine    `xml:"HairPin,omitempty"`
// 	Next    *FluffyNext `xml:"next,omitempty"`
// 	Type    FluffyType  `xml:"type,attr"`
// 	Prev    *FluffyNext `xml:"prev,omitempty"`
// 	Volta   *VoltaClass `xml:"Volta,omitempty"`
// 	Ottava  *BarLine    `xml:"Ottava,omitempty"`
// }
//
// type FluffyNext struct {
// 	Location FluffyLocation `xml:"location"`
// }
//
// type FluffyLocation struct {
// 	Fractions *string `xml:"fractions,omitempty"`
// 	Measures  *string `xml:"measures,omitempty"`
// }
//
// type VoltaClass struct {
// 	EndHookType string `xml:"endHookType"`
// 	BeginText   string `xml:"beginText"`
// 	Endings     string `xml:"endings"`
// }
//
// type StaffText struct {
// 	Text string `xml:"text"`
// }
//
// type Tempo struct {
// 	Tempo      string    `xml:"tempo"`
// 	FollowText string    `xml:"followText"`
// 	Text       TextClass `xml:"text"`
// }
//
// type TextClass struct {
// 	B    []BElement `xml:"b"`
// 	Font Font       `xml:"font"`
// 	Text string     `xml:",chardata"`
// }
//
// type BClass struct {
// 	Font Font   `xml:"font"`
// 	Text string `xml:",chardata"`
// }
//
// type Font struct {
// 	Face string `xml:"face,attr"`
// }

type TimeSig struct {
	SigN string `xml:"sigN"`
	SigD string `xml:"sigD"`
}

// type TupletClass struct {
// 	Offset      *Offset     `xml:"offset,omitempty"`
// 	NormalNotes string      `xml:"normalNotes"`
// 	ActualNotes string      `xml:"actualNotes"`
// 	BaseNote    BaseNote    `xml:"baseNote"`
// 	Number      TextElement `xml:"Number"`
// }

type TextElement struct {
	Style StyleEnum `xml:"style"`
	Text  string    `xml:"text"`
}

type VBox struct {
	Height string        `xml:"height"`
	Text   []TextElement `xml:"Text"`
}

// type BaseNote string
//
// const (
// 	Eighth  BaseNote = "eighth"
// 	Quarter BaseNote = "quarter"
// 	Half    BaseNote = "half"
// 	The16Th BaseNote = "16th"
// )
//
// type Subtype string
//
// const (
// 	AccidentalFlat    Subtype = "accidentalFlat"
// 	AccidentalNatural Subtype = "accidentalNatural"
// 	AccidentalSharp   Subtype = "accidentalSharp"
// )
//
// type Fractions string
//
// const (
// 	Fractions112 Fractions = "-1/12"
// 	Fractions38  Fractions = "-3/8"
// 	The112       Fractions = "1/12"
// 	The38        Fractions = "3/8"
// )
//
// type PurpleType string
//
// const (
// 	Slur PurpleType = "Slur"
// 	Tie  PurpleType = "Tie"
// )
//
// type BeamMode string
//
// const (
// 	Begin32 BeamMode = "begin32"
// 	Mid     BeamMode = "mid"
// )
//
// type FluffyType string
//
// const (
// 	HairPin FluffyType = "HairPin"
// 	Ottava  FluffyType = "Ottava"
// 	Volta   FluffyType = "Volta"
// )

type StyleEnum string

const (
	Composer StyleEnum = "Composer"
	Subtitle StyleEnum = "Subtitle"
	Title    StyleEnum = "Title"
	Tuplet   StyleEnum = "Tuplet"
)

type Chord struct {
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
	Slur *Slur  `xml:"Slur"`
	Next *Next  `xml:"next"`
}

type Slur struct {
	Text string `xml:",chardata"`
}

type Next struct {
	Location *Location `xml:"location"`
}

type Location struct {
	Fractions string `xml:"fractions"`
}

type Note struct {
	Pitch int `xml:"pitch"`
	TPC   int `xml:"tpc"`
}

// type ArticulationUnion struct {
// 	BarLine      *BarLine
// 	BarLineArray []BarLine
// }
//
// type StickyNote struct {
// 	FluffyNote      *FluffyNote
// 	PurpleNoteArray []PurpleNote
// }
//
// type ChordSpanner struct {
// 	NoteSpanner      *NoteSpanner
// 	NoteSpannerArray []NoteSpanner
// }
//
// type IndigoNote struct {
// 	FluffyNoteArray []FluffyNote
// 	TentacledNote   *TentacledNote
// }
//
// type RESTUnion struct {
// 	RESTElement      *RESTElement
// 	RESTElementArray []RESTElement
// }
//
// type VoiceSpanner struct {
// 	PurpleSpanner      *PurpleSpanner
// 	PurpleSpannerArray []PurpleSpanner
// }
//
// type BElement struct {
// 	BClass *BClass
// 	String *string
// }
