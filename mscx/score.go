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

// ScoreZip represents a MuseScore 3 score in `mscz` (zip'd) format.
type ScoreZip struct {
	// MetaInf
	MuseScore *MuseScore `xml:"museScore"`
	// Thumbnails
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
	LayerTag        *LayerTag  `xml:"LayerTag"`
	CurrentLayer    int        `xml:"currentLayer"`
	Division        int        `xml:"Division"`
	Style           *Style     `xml:"Style"`
	ShowInvisible   int        `xml:"showInvisible"`
	ShowUnprintable int        `xml:"showUnprintable"`
	ShowFrames      int        `xml:"showFrames"`
	ShowMargins     int        `xml:"showMargins"`
	MetaTags        []*MetaTag `xml:"metaTag"`
	Part            *Part      `xml:"Part"`
	Staffs          []*Staff   `xml:"Staff"`
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
	UseStandardNoteNames int     `xml:"useStandardNoteNames"`
	Spatium              float64 `xml:"Spatium"`
}

// MetaTag represents the XML data of the same name.
type MetaTag struct {
	Text string `xml:",chardata"`
	Name string `xml:"name,attr"`
}

// Part represents the XML data of the same name.
type Part struct {
	Staff      []*Staff    `xml:"Staff"`
	TrackName  string      `xml:"trackName"`
	Instrument *Instrument `xml:"Instrument,omitempty"`
}

// Staff represents the XML data of the same name.
type Staff struct {
	StaffType   StaffType `xml:"StaffType"`
	Bracket     *Bracket  `xml:"bracket,omitempty"`
	BarLineSpan int       `xml:"barLineSpan"`
	ID          string    `xml:"id,attr"`
	DefaultClef string    `xml:"defaultClef,omitempty"`
}

// StaffType represents the XML data of the same name.
type StaffType struct {
	Name  string `xml:"name"`
	Group string `xml:"group,attr"`
}

// Bracket represents the XML data of the same name.
type Bracket struct {
	Type string `xml:"type,attr"`
	Span string `xml:"span,attr"`
	Col  string `xml:"col,attr"`
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
	Staff string `xml:"staff"`
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
