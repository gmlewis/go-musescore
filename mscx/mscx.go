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

//  Package mscx parses MuseScore 3 `*.mscx` or `*.mscz` files into Go structs.
package mscx

import (
	"archive/zip"
	"bytes"
	"encoding/xml"
	"fmt"
	_ "image/png"
	"io"
	"log"
	"os"
	"strings"
)

// CallbackFn is an optional function that will be called with each file entry
// found in the ZIP.
type CallbackFn func(filename string, content []byte)

// NewFromFile reads a `*.mscx` or `*.mscz` file and returns the resulting parsed score.
func NewFromFile(filename string, callback CallbackFn) (*ScoreZip, error) {
	b, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return New(b, callback)
}

// New reads `mscx` or `mscz` data and returns the resulting parsed score.
func New(buf []byte, callback CallbackFn) (*ScoreZip, error) {
	if len(buf) > len(xmlStart) && string(buf[0:len(xmlStart)]) == xmlStart {
		return parseXML(buf)
	}
	return parseZip(buf, callback)
}

const xmlStart = "<?xml "

func parseXML(buf []byte) (*ScoreZip, error) {
	var s *MuseScore
	if err := xml.Unmarshal(buf, &s); err != nil {
		return nil, err
	}
	return &ScoreZip{MuseScore: s}, nil
}

func parseZip(buf []byte, callback CallbackFn) (*ScoreZip, error) {
	r, err := zip.NewReader(bytes.NewReader(buf), int64(len(buf)))
	if err != nil {
		return nil, fmt.Errorf("zip.NewReader: %w", err)
	}

	var result *ScoreZip
	for _, fh := range r.File {
		log.Printf("fh.Name=%v", fh.Name)
		if fh.FileInfo().IsDir() {
			log.Printf("found dir: %v", fh.Name)
			continue
		}

		rc, err := fh.Open()
		if err != nil {
			return nil, fmt.Errorf("zip.fh.Open(%q): %w", fh.Name, err)
		}

		nb, err := io.ReadAll(rc)
		if err != nil {
			return nil, fmt.Errorf("zip.io.ReadAll(%q): %w", fh.Name, err)
		}

		if callback != nil {
			callback(fh.Name, nb)
		}

		if strings.HasSuffix(fh.Name, ".mscx") {
			result, err = parseXML(nb)
			if err != nil {
				return nil, fmt.Errorf("zip.parseXML(%q): %w", fh.Name, err)
			}
		}

		if err := rc.Close(); err != nil {
			return nil, fmt.Errorf("zip.rc.Close() for %q: %v", fh.Name, err)
		}
	}

	return result, nil
}
