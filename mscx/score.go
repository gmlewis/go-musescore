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
	Score           *Score `xml:"score"`
}

// Score represents the XML data of the same name.
type Score struct {
	LayerTag        *LayerTag `xml:"LayerTag"`
	CurrentLayer    int       `xml:"currentLayer"`
	Division        int       `xml:"Division"`
	Style           *Style    `xml:"Style"`
	ShowInvisible   int       `xml:"showInvisible"`
	ShowUnprintable int       `xml:"showUnprintable"`
	ShowFrames      int       `xml:"showFrames"`
	ShowMargins     int       `xml:"showMargins"`
	MetaTags        []MetaTag `xml:"metaTag"`
	Part            *Part     `xml:"Part"`
	Staffs          []*Staff  `xml:"Staff"`
}

// LayerTag represents the XML data of the same name.
type LayerTag struct {
	ID  string `xml:"id,attr"`
	Tag string `xml:"tag,attr"`
}

// Style represents the XML data of the same name.
type Style struct {
}

// MetaTag represents the XML data of the same name.
type MetaTag struct {
}

// Part represents the XML data of the same name.
type Part struct {
}

// Staff represents the XML data of the same name.
type Staff struct {
}
