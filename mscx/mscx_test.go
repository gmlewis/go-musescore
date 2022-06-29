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

import (
	_ "embed"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

//go:embed testfiles/001-O_For_a_Thousand_Tongues_to_Sing.mscz
var test01 []byte

//go:embed testfiles/003-Come,_Thou_Almighty_King.mscz
var test02 []byte

//go:embed testfiles/017-How_Great_Thou_Art.mscz
var test03 []byte

//go:embed testfiles/020-A_Mighty_Fortress_Is_Our_God.mscz
var test04 []byte

//go:embed testfiles/027-Immortal,_Invisible,_God_Only_Wise.mscz
var test05 []byte

//go:embed testfiles/Ben_Hur_Chariot_Race_March.mscz
var test06 []byte

//go:embed testfiles/The_Ice_Palace.mscz
var test07 []byte

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		in   []byte
		want *ScoreZip
	}{
		{
			name: "test01",
			in:   test01,
			want: test01Data,
		},
		{
			name: "test02",
			in:   test02,
			want: test02Data,
		},
		{
			name: "test03",
			in:   test03,
			want: test03Data,
		},
		{
			name: "test04",
			in:   test04,
			want: test04Data,
		},
		{
			name: "test05",
			in:   test05,
			want: test05Data,
		},
		{
			name: "test06",
			in:   test06,
			want: test06Data,
		},
		{
			name: "test07",
			in:   test07,
			want: test07Data,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var xml []byte
			cb := func(fn string, buf []byte) {
				if strings.HasSuffix(fn, ".mscx") {
					xml = buf
				}
			}
			got, err := New(test01, cb)
			if err != nil {
				t.Fatal(err)
			}
			strippedXML := strip(string(xml))

			// Compare Go structs to Go structs
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("New(%q) Go structs to Go struct differs (-want +got):\n%s", tt.name, diff)
			}

			// Compare XML in to Go "tt.want" structs (as XML)
			wantXML, err := tt.want.XML()
			if err != nil {
				t.Fatalf("tt.want.XML: %v", err)
			}

			strippedWant := strip(string(wantXML))
			if diff := cmp.Diff(strippedXML, strippedWant); diff != "" {
				t.Log("strippedXML:\n", string(strippedXML))
				t.Log("wantXML:\n", string(wantXML))
				t.Errorf("New(%q) tt.want XML differs (-want('orig') +got('tt.want')):\n%s", tt.name, diff)
			}

			// Compare XML in to XML out
			gotXML, err := got.XML()
			if err != nil {
				t.Fatalf("got.XML: %v", err)
			}

			strippedGot := strip(string(gotXML))
			if diff := cmp.Diff(strippedXML, strippedGot); diff != "" {
				t.Log("gotXML:\n", string(gotXML))
				t.Fatalf("New(%q) got XML differs (-want +got):\n%s", tt.name, diff)
			}
		})
	}
}

func strip(s string) string {
	lines := strings.Split(s, "\n")
	result := make([]string, 0, len(lines))
	for _, line := range lines {
		v := strings.TrimSpace(line)
		if v == "" {
			continue
		}
		result = append(result, v)
	}
	return strings.Join(result, "\n")
}
