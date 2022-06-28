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

var test06Data = &ScoreZip{
	MuseScore: &MuseScore{
		Version:         "3.01",
		ProgramVersion:  "3.2.3",
		ProgramRevision: "d2d863f",
		Score: &Score{
			LayerTag:        &LayerTag{ID: "0", Tag: "default"},
			Division:        480,
			Style:           &Style{},
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
		},
	},
}