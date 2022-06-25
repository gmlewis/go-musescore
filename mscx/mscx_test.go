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
		want *Score
	}{
		{
			name: "test01",
			in:   test01,
		},
		{
			name: "test02",
			in:   test02,
		},
		{
			name: "test03",
			in:   test03,
		},
		{
			name: "test04",
			in:   test04,
		},
		{
			name: "test05",
			in:   test05,
		},
		{
			name: "test06",
			in:   test06,
		},
		{
			name: "test07",
			in:   test07,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(test01)
			if err != nil {
				t.Fatal(err)
			}

			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("New(%q) differs (-want +got):\n%s", tt.name, diff)
			}
		})
	}
}
