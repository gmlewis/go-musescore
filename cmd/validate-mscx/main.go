// -*- compile-command: "go run main.go"; -*-

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

// validate-mscx reads and parses a `*.mscz` or `*.mscx` file, then
// compares its generated XML with the original XML and prints the
// differences. This is used to validate the XML parser.
// If there are differences, it will print them and terminate.
package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/gmlewis/go-musescore/mscx"
	"github.com/google/go-cmp/cmp"
)

func main() {
	flag.Parse()

	for _, arg := range flag.Args() {
		if err := processFile(arg); err != nil {
			log.Fatal(err)
		}
	}
}

func processFile(filename string) error {
	log.Printf("Processing file: %v", filename)
	var xml []byte
	cb := func(fn string, buf []byte) {
		log.Printf("%v (%v bytes)", fn, len(buf))
		if strings.HasSuffix(fn, ".mscx") {
			xml = buf
		}
	}
	sz, err := mscx.NewFromFile(filename, cb)
	if err != nil {
		return err
	}
	strippedXML := strip(string(xml))

	// Compare XML in to XML out
	gotXML, err := sz.XML()
	if err != nil {
		return err
	}

	strippedGot := strip(string(gotXML))
	if diff := cmp.Diff(strippedXML, strippedGot); diff != "" {
		log.Println("gotXML:\n", string(gotXML))
		return fmt.Errorf("NewFromFile(%q) got XML differs (-want +got):\n%s", filename, diff)
	}

	return nil
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
