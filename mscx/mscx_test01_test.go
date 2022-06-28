// -*- compile-command: "go test -v ./..."; -*-

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

var test01Data = &ScoreZip{
	MuseScore: &MuseScore{
		Version:         "3.01",
		ProgramVersion:  "3.2.3",
		ProgramRevision: "d2d863f",
		Score: &Score{
			LayerTag: &LayerTag{ID: "0", Tag: "default"},
			Division: 480,
			Style: &Style{
				PageWidth:          8.27,
				PageHeight:         11.69,
				PagePrintableWidth: 7.4826,
				Spatium:            1.76389,
			},
			ShowInvisible:   1,
			ShowUnprintable: 1,
			ShowFrames:      1,
			MetaTags: []*MetaTag{
				{Name: "arranger"},
				{Name: "composer", Text: "Charles Wesley"},
				{Name: "copyright", Text: "1964, 1966 Board of Publication of The Methodist Church, Inc."},
				{Name: "creationDate", Text: "2022-06-19"},
				{Name: "lyricist"},
				{Name: "movementNumber"},
				{Name: "movementTitle"},
				{Name: "platform", Text: "Linux"},
				{Name: "poet"},
				{Name: "source"},
				{Name: "translator"},
				{Name: "workNumber"},
				{Name: "workTitle", Text: "O For a Thousand Tongues to Sing"},
			},
			Part: &Part{
				Staff: []*Staff{
					{
						StaffType:   StaffType{Name: "stdNormal", Group: "pitched"},
						Bracket:     &Bracket{Type: "1", Span: "2", Col: "0"},
						BarLineSpan: 1,
						ID:          "1",
					},
					{
						StaffType:   StaffType{Name: "stdNormal", Group: "pitched"},
						ID:          "2",
						DefaultClef: "F",
					},
				},
				TrackName: "Piano",
				Instrument: &Instrument{
					LongName:     "Piano",
					ShortName:    "Pno.",
					TrackName:    "Piano",
					MinPitchP:    "21",
					MaxPitchP:    "108",
					MinPitchA:    "21",
					MaxPitchA:    "108",
					InstrumentID: "keyboard.piano",
					Clef:         Clef{Text: "F"},
					Articulation: []*ArticulationElement{
						{Velocity: "100", GateTime: "95"},
						{Velocity: "100", GateTime: "33", Name: "staccatissimo"},
						{Velocity: "100", GateTime: "50", Name: "staccato"},
						{Velocity: "100", GateTime: "67", Name: "portato"},
						{Velocity: "100", GateTime: "100", Name: "tenuto"},
						{Velocity: "120", GateTime: "67", Name: "marcato"},
						{Velocity: "150", GateTime: "100", Name: "sforzato"},
						{Velocity: "150", GateTime: "50", Name: "sforzatoStaccato"},
						{Velocity: "120", GateTime: "50", Name: "marcatoStaccato"},
						{Velocity: "120", GateTime: "100", Name: "marcatoTenuto"},
					},
					Channel: Channel{Program: Program{Value: "0"}, Synti: "Fluid"},
				},
			},
			Staffs: []*Staff{{ID: "1"}, {ID: "2"}},
		},
	},
}
