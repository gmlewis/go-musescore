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
				Staff: []*PartStaff{
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
					Clef:         Clef{Staff: "2", Text: "F"},
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
			Staffs: []*ScoreStaff{
				{
					ID: "1",
					VBox: &VBox{
						Height: "10",
						Text: []TextElement{
							{Style: "Title", Text: "O For a Thousand Tongues to Sing"},
							{Style: "Composer", Text: "Charles Wesley"},
						},
					},
					Measure: []*Measure{
						{
							Irregular: 1,
							Voice: &Voice{
								TimeSig: &TimeSig{
									SigN: "3",
									SigD: "2",
								},
								Chord: []*Chord{
									{
										DurationType: "half",
										Lyrics: []*Lyrics{
											{Text: "1.O"},
											{No: 1, Text: "2.My"},
											{No: 2, Text: "3.Je", Syllabic: "begin"},
											{No: 3, Text: "4.He"},
											{No: 4, Text: "5.He"},
											{No: 5, Text: "6.Hear"},
										},
										Note: []*Note{
											{Pitch: 64, TPC: 18},
										},
									},
								},
								KeySig: &KeySig{Accidental: "3"},
							},
							Len: "1/2",
						},
						{
							Voice: &Voice{
								Chord: []*Chord{
									{
										DurationType: "quarter",
										Lyrics: []*Lyrics{
											{Text: "for"},
											{No: 1, Text: "gra", Syllabic: "begin"},
											{No: 2, Text: "sus!", Syllabic: "end", TicksF: "0/2"},
											{No: 3, Text: "breaks"},
											{No: 4, Text: "speaks,"},
											{No: 5, Text: "him,"},
										},
										Note: []*Note{
											{Pitch: 64, TPC: 18},
											{Pitch: 69, TPC: 17},
										},
									},
									{
										DurationType: "quarter",
										Lyrics: []*Lyrics{
											{Text: "a"},
											{No: 1, Syllabic: "end", Text: "cious", TicksF: "0/4"},
											{No: 2, Text: "the"},
											{No: 3, Text: "the"},
											{No: 4, Text: "and"},
											{No: 5, Text: "ye"},
										},
										Note: []*Note{
											{Pitch: 61, TPC: 21},
											{Pitch: 69, TPC: 17},
										},
									},
									{
										DurationType: "quarter",
										Lyrics: []*Lyrics{
											{Syllabic: "begin", Text: "thou"},
											{No: 1, Syllabic: "begin", Text: "Mas"},
											{No: 2, Text: "name"},
											{No: 3, Text: "power"},
											{No: 4, Syllabic: "begin", Text: "listen"},
											{No: 5, Text: "deaf;"},
										},
										Note: []*Note{
											{Pitch: 64, TPC: 18},
											{Pitch: 71, TPC: 19},
										},
									},
									{
										DurationType: "quarter",
										Lyrics: []*Lyrics{
											{Syllabic: "end", Text: "sand", TicksF: "0/4"},
											{No: 1, Syllabic: "end", Text: "ter", TicksF: "0/4"},
											{No: 2, Text: "that"},
											{No: 3, Text: "of"},
											{No: 4, Syllabic: "end", Text: "ing", TicksF: "0/4"},
											{No: 5, Text: "his"},
										},
										Note: []*Note{
											{Pitch: 64, TPC: 18},
											{Pitch: 71, TPC: 19},
										},
									},
								},
							},
						},
						{
							Voice: &Voice{
								Chord: []*Chord{
									{
										DurationType: "quarter",
										Lyrics: []*Lyrics{
											{Text: "tongues"},
											{No: 1, Text: "and"},
											{No: 2, Text: "charms"},
											{No: 3, Syllabic: "begin", Text: "can"},
											{No: 4, Text: "to"},
											{No: 5, Text: "paise,"},
										},
										Note: []*Note{
											{Pitch: 64, TPC: 18},
											{Pitch: 73, TPC: 21},
										},
									},
									{
										DurationType: "quarter",
										Lyrics: []*Lyrics{
											{Text: "to"},
											{No: 1, Text: "my"},
											{No: 2, Text: "our"},
											{No: 3, Syllabic: "end", Text: "celed", TicksF: "0/4"},
											{No: 4, Text: "his"},
											{No: 5, Text: "ye"},
										},
										Note: []*Note{
											{Pitch: 62, TPC: 16},
											{Pitch: 71, TPC: 19},
										},
									},
									{
										DurationType: "half",
										Lyrics: []*Lyrics{
											{Text: "sing"},
											{No: 1, Text: "God,"},
											{No: 2, Text: "fear,"},
											{No: 3, Text: "sin,"},
											{No: 4, Text: "voice,"},
											{No: 5, Text: "dumb,"},
										},
										Note: []*Note{
											{Pitch: 61, TPC: 21},
											{Pitch: 69, TPC: 17},
										},
									},
									{
										DurationType: "quarter",
										Lyrics: []*Lyrics{
											{Text: "My"},
											{No: 1, Syllabic: "begin", Text: "As"},
											{No: 2, Text: "That"},
											{No: 3, Text: "He"},
											{No: 4, Text: "New"},
											{No: 5, Text: "Your"},
										},
										Note: []*Note{
											{Pitch: 64, TPC: 18},
											{Pitch: 71, TPC: 19},
										},
									},
								},
							},
						},
						{
							Voice: &Voice{
								Chord: []*Chord{
									{
										DurationType: "quarter",
										Lyrics: []*Lyrics{
											{Text: "great"}, {No: 1, Syllabic: "end", Text: "sist", TicksF: "0/2"},
											{No: 2, Text: "bids"}, {No: 3, Text: "sets"}, {No: 4, Text: "life"},
											{No: 5, Syllabic: "begin", Text: "loos"},
										},
										Note: []*Note{
											{Pitch: 64, TPC: 18},
											{Pitch: 73, TPC: 21},
										},
									},
									{
										DurationType: "quarter",
										Lyrics: []*Lyrics{
											{Syllabic: "begin", Text: "Re"},
											{No: 1, Text: "me"}, {No: 2, Text: "our"}, {No: 3, Text: "the"},
											{No: 4, Text: "the"}, {No: 5, Syllabic: "end", Text: "ened", TicksF: "0/4"},
										},
										Note: []*Note{
											{Pitch: 64, TPC: 18},
											{Pitch: 73, TPC: 21},
										},
									},
									{
										DurationType: "quarter",
										Lyrics: []*Lyrics{
											{Syllabic: "middle", Text: "deem"}, {No: 1, Text: "to"},
											{No: 2, Syllabic: "begin", Text: "sor"},
											{No: 3, Syllabic: "begin", Text: "pris"},
											{No: 4, Text: "dead"},
											{No: 5, Text: "tongues"},
										},
										Note: []*Note{
											{Pitch: 64, TPC: 18},
											{Pitch: 74, TPC: 16},
										},
									},
									{
										DurationType: "quarter",
										Lyrics: []*Lyrics{
											{Syllabic: "end", Text: "er's", TicksF: "0/4"},
											{No: 1, Syllabic: "begin", Text: "pro"},
											{No: 2, Syllabic: "end", Text: "rows", TicksF: "0/4"},
											{No: 3, Syllabic: "end", Text: "oner", TicksF: "0/4"},
											{No: 4, Syllabic: "begin", Text: "re"},
											{No: 5, Syllabic: "begin", Text: "em"},
										},
										Note: []*Note{
											{Pitch: 64, TPC: 18},
											{Pitch: 73, TPC: 21},
										},
									},
								},
							},
						},
						{
							Voice: &Voice{
								Chord: []*Chord{
									{
										DurationType: "half",
										Lyrics: []*Lyrics{
											{Text: "praise,"}, {No: 1, Syllabic: "end", Text: "claim,", TicksF: "0/2"},
											{No: 2, Text: "cease,"}, {No: 3, Text: "free;"},
											{No: 4, Syllabic: "end", Text: "ceive;", TicksF: "0/2"},
											{No: 5, Syllabic: "end", Text: "ploy;", TicksF: "0/2"},
										},
										Note: []*Note{
											{Pitch: 64, TPC: 18},
											{Pitch: 71, TPC: 19},
										},
									},
									{
										DurationType: "half",
										Lyrics: []*Lyrics{
											{Text: "The"}, {No: 1, Text: "To"},
											{No: 2, Text: "'Tis"}, {No: 3, Text: "His"}, {No: 4, Text: "The"},
											{No: 5, Text: "Ye"},
										},
										Note: []*Note{
											{Pitch: 68, TPC: 22},
											{Pitch: 76, TPC: 18},
										},
									},
								},
							},
						},
						{
							Voice: &Voice{
								Chord: []*Chord{
									{
										DurationType: "half",
										Lyrics: []*Lyrics{
											{Syllabic: "begin", Text: "glo"}, {No: 1, Text: "spread"},
											{No: 2, Syllabic: "begin", Text: "mu"}, {No: 3, Text: "blood"},
											{No: 4, Syllabic: "begin", Text: "mourn"}, {No: 5, Text: "blind,"},
										},
										Note: []*Note{
											{Pitch: 69, TPC: 17},
											{Pitch: 76, TPC: 18},
										},
									},
									{
										DurationType: "half",
										Lyrics: []*Lyrics{
											{Syllabic: "end", Text: "ries", TicksF: "0/4"}, {No: 1, Text: "thro'"},
											{No: 2, Syllabic: "end", Text: "sic", TicksF: "0/4"}, {No: 3, Text: "can"},
											{No: 4, Syllabic: "end", Text: "ful,", TicksF: "0/4"},
											{No: 5, Syllabic: "begin", Text: "be"},
										},
										Note: []*Note{
											{Pitch: 64, TPC: 18},
											{Pitch: 73, TPC: 21},
										},
									},
									{
										DurationType: "half",
										Lyrics: []*Lyrics{
											{Text: "of"}, {No: 1, Text: "all"},
											{No: 2, Text: "in"}, {No: 3, Text: "make"},
											{No: 4, Syllabic: "begin", Text: "bro"},
											{No: 5, Syllabic: "end", Text: "hold", TicksF: "0/4"},
										},
										Note: []*Note{
											{Pitch: 64, TPC: 18},
											{Pitch: 73, TPC: 21},
										},
									},
									{
										DurationType: "half",
										Lyrics: []*Lyrics{
											{Text: "my"},
											{No: 1, Text: "the"},
											{No: 2, Text: "the"},
											{No: 3, Text: "the"},
											{No: 4, Syllabic: "end", Text: "ken", TicksF: "0/4"},
											{No: 5, Text: "your"},
										},
										Note: []*Note{
											{Pitch: 64, TPC: 18},
											{Pitch: 69, TPC: 17},
										},
									},
								},
							},
						},
						{
							Voice: &Voice{
								Chord: []*Chord{
									{
										DurationType: "quarter",
										Lyrics: []*Lyrics{
											{Text: "God"},
											{No: 1, Text: "earth"},
											{No: 2, Syllabic: "begin", Text: "sin"},
											{No: 3, Syllabic: "begin", Text: "foul"},
											{No: 4, Text: "hearts"},
											{No: 5, Syllabic: "begin", Text: "Sav"},
										},
										Note: []*Note{
											{Pitch: 66, TPC: 20},
											{Pitch: 69, TPC: 17},
										},
									},
									{
										DurationType: "quarter",
										Lyrics: []*Lyrics{
											{Text: "and"},
											{No: 1, Syllabic: "begin", Text: "a"},
											{No: 2, Syllabic: "end", Text: "ners'", TicksF: "0/4"},
											{No: 3, Syllabic: "end", Text: "est", TicksF: "0/4"},
											{No: 4, Syllabic: "begin", Text: "re"},
											{No: 5, Syllabic: "end", Text: "ior", TicksF: "0/4"},
										},
										Note: []*Note{
											{Pitch: 62, TPC: 16},
											{Pitch: 66, TPC: 20},
										},
									},
									{
										DurationType: "quarter",
										Lyrics: []*Lyrics{
											{Text: "King,"},
											{No: 1, Syllabic: "end", Text: "broad", TicksF: "0/4"},
											{No: 2, Text: "ears,"},
											{No: 3, Text: "clean;"},
											{No: 4, Syllabic: "end", Text: "joice;", TicksF: "0/4"},
											{No: 5, Text: "come;"},
										},
										Note: []*Note{
											{Pitch: 62, TPC: 16},
											{Pitch: 66, TPC: 20},
										},
									},
									{
										DurationType: "quarter",
										Lyrics: []*Lyrics{
											{Text: "The"},
											{No: 1, Text: "The"},
											{No: 2, Text: "'Tis"},
											{No: 3, Text: "His"},
											{No: 4, Text: "The"},
											{No: 5, Text: "And"},
										},
										Spanner: []*Spanner{
											{
												Type: "Slur",
												Slur: &Slur{},
												Next: &Next{Location: &Location{Fractions: "1/4"}},
											},
											{
												Type: "Slur",
												Slur: &Slur{},
												Next: &Next{Location: &Location{Fractions: "1/4"}},
											},
										},
										Note: []*Note{
											{Pitch: 66, TPC: 20},
											{Pitch: 69, TPC: 17},
											{Pitch: 62, TPC: 16},
											{Pitch: 66, TPC: 20},
										},
									},
								},
							},
						},
						{
							Voice: &Voice{
								Chord: []*Chord{
									{
										DurationType: "half",
										Lyrics: []*Lyrics{
											{Syllabic: "begin", Text: "tri"},
											{No: 1, Syllabic: "middle", Text: "hon"},
											{No: 2, Text: "life,"},
											{No: 3, Text: "blood"},
											{No: 4, Syllabic: "begin", Text: "hum"},
											{No: 5, Text: "leap,"},
										},
										Note: []*Note{
											{Pitch: 61, TPC: 21},
											{Pitch: 64, TPC: 18},
										},
									},
									{
										DurationType: "half",
										Lyrics: []*Lyrics{
											{Syllabic: "end", Text: "umphs", TicksF: "0/4"},
											{No: 1, Syllabic: "end", Text: "ors", TicksF: "0/4"},
											{No: 2, Text: "and"},
											{No: 3, Syllabic: "begin", Text: "a"},
											{No: 4, Syllabic: "end", Text: "ble", TicksF: "0/4"},
											{No: 5, Text: "ye"},
										},
										Note: []*Note{
											{Pitch: 61, TPC: 21},
											{Pitch: 69, TPC: 17},
										},
									},
									{
										DurationType: "half",
										Lyrics: []*Lyrics{
											{Text: "of"},
											{No: 1, Text: "of"},
											{No: 2, Text: "health,"},
											{No: 3, Syllabic: "end", Text: "vailed", TicksF: "0/4"},
											{No: 4, Text: "poor,"},
											{No: 5, Text: "lame,"},
										},
										Note: []*Note{
											{Pitch: 61, TPC: 21},
											{Pitch: 69, TPC: 17},
										},
									},
									{
										DurationType: "half",
										Lyrics: []*Lyrics{
											{Text: "his"},
											{No: 1, Text: "thy"},
											{No: 2, Text: "and"},
											{No: 3, Text: "for"},
											{No: 4, Syllabic: "begin", Text: "be"},
											{No: 5, Text: "for"},
										},
										Note: []*Note{
											{Pitch: 62, TPC: 16},
											{Pitch: 71, TPC: 19},
										},
									},
								},
							},
						},
						{
							Voice: &Voice{
								Chord: []*Chord{
									{
										DurationType: "whole",
										Lyrics: []*Lyrics{
											{Text: "grace!"},
											{No: 1, Text: "name."},
											{No: 2, Text: "peace."},
											{No: 3, Text: "me."},
											{No: 4, Syllabic: "end", Text: "lieve."},
											{No: 5, Text: "joy."},
										},
										Note: []*Note{{Pitch: 61, TPC: 21},
											{Pitch: 69, TPC: 17}},
									},
								},
								BarLine: &BarLine{Subtype: "double"},
								Rest:    &Rest{},
							},
						},
						{
							Voice: &Voice{
								Chord: []*Chord{
									{
										DurationType: "whole",
										Lyrics:       []*Lyrics{{No: 5, Syllabic: "begin", Text: "A"}},
										Note: []*Note{{Pitch: 62, TPC: 16},
											{Pitch: 69, TPC: 17}},
									},
								},
							},
						},
						{
							Voice: &Voice{
								Chord: []*Chord{
									{
										DurationType: "whole",
										Lyrics:       []*Lyrics{{No: 5, Syllabic: "end", Text: "men."}},
										Note: []*Note{{Pitch: 61, TPC: 21},
											{Pitch: 69, TPC: 17}},
									},
								},
							},
						},
					},
				},
				{
					ID: "2",
					Measure: []*Measure{
						{
							Voice: &Voice{
								TimeSig: &TimeSig{
									SigN: "3",
									SigD: "2",
								},
								Chord: []*Chord{
									{
										DurationType: "half",
										Note:         []*Note{{Pitch: 52, TPC: 18}},
									},
								},
								KeySig: &KeySig{Accidental: "3"},
							},
							Len: "1/2",
						},
						{
							Voice: &Voice{
								Chord: []*Chord{
									{
										DurationType: "half",
										Note: []*Note{
											{Pitch: 49, TPC: 21},
											{Pitch: 52, TPC: 18},
											{Pitch: 45, TPC: 17},
											{Pitch: 57, TPC: 17},
											{Pitch: 52, TPC: 18},
											{Pitch: 56, TPC: 22},
											{Pitch: 52, TPC: 18},
											{Pitch: 56, TPC: 22},
										},
									},
								},
							},
						},
						{
							Voice: &Voice{
								Chord: []*Chord{
									{
										DurationType: "half",
										Note: []*Note{
											{Pitch: 57, TPC: 17},
											{Pitch: 52, TPC: 18},
											{Pitch: 56, TPC: 22},
											{Pitch: 54, TPC: 20},
											{Pitch: 57, TPC: 17},
											{Pitch: 52, TPC: 18},
											{Pitch: 56, TPC: 22},
										},
									},
								},
							},
						},
						{
							Voice: &Voice{
								Chord: []*Chord{
									{
										DurationType: "half",
										Note: []*Note{
											{Pitch: 57, TPC: 17},
											{Pitch: 57, TPC: 17},
											{Pitch: 56, TPC: 22},
											{Pitch: 59, TPC: 19},
											{Pitch: 57, TPC: 17},
										},
									},
								},
							},
						},
						{
							Voice: &Voice{
								Chord: []*Chord{
									{
										DurationType: "half",
										Note: []*Note{
											{Pitch: 52, TPC: 18},
											{Pitch: 56, TPC: 22},
											{Pitch: 52, TPC: 18},
											{Pitch: 59, TPC: 19},
										},
									},
								},
							},
						},
						{
							Voice: &Voice{
								Chord: []*Chord{
									{
										DurationType: "half",
										Note: []*Note{
											{Pitch: 57, TPC: 17},
											{Pitch: 61, TPC: 21},
											{Pitch: 57, TPC: 17},
											{Pitch: 57, TPC: 17},
											{Pitch: 49, TPC: 21},
											{Pitch: 57, TPC: 17},
										},
									},
								},
							},
						},
						{
							Voice: &Voice{
								Chord: []*Chord{
									{
										DurationType: "half",
										Note: []*Note{
											{Pitch: 50, TPC: 16},
											{Pitch: 57, TPC: 17},
											{Pitch: 50, TPC: 16},
											{Pitch: 57, TPC: 17},
											{Pitch: 50, TPC: 16},
											{Pitch: 57, TPC: 17},
											{Pitch: 50, TPC: 16},
											{Pitch: 57, TPC: 17},
										},
									},
								},
							},
						},
						{
							Voice: &Voice{
								Chord: []*Chord{
									{
										DurationType: "half",
										Note: []*Note{
											{Pitch: 52, TPC: 18},
											{Pitch: 57, TPC: 17},
											{Pitch: 52, TPC: 18},
											{Pitch: 57, TPC: 17},
											{Pitch: 52, TPC: 18},
											{Pitch: 57, TPC: 17},
											{Pitch: 52, TPC: 18},
											{Pitch: 56, TPC: 22},
										},
									},
								},
							},
						},
						{
							Voice: &Voice{
								Chord: []*Chord{
									{
										DurationType: "whole",
										Note: []*Note{{Pitch: 45, TPC: 17},
											{Pitch: 57, TPC: 17}},
									},
								},
								BarLine: &BarLine{Subtype: "double"},
								Rest:    &Rest{},
							},
						},
						{
							Voice: &Voice{
								Chord: []*Chord{
									{
										DurationType: "whole",
										Note: []*Note{{Pitch: 50, TPC: 16},
											{Pitch: 54, TPC: 20}},
									},
								},
							},
						},
						{
							Voice: &Voice{
								Chord: []*Chord{
									{
										DurationType: "whole",
										Note: []*Note{{Pitch: 45, TPC: 17},
											{Pitch: 52, TPC: 18}},
									},
								},
							},
						},
					},
				},
			},
		},
	},
}
