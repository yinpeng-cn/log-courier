// This file is for go run only
// +build ignore

/*
 * Copyright 2015 Jason Woods.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"bytes"
	"io/ioutil"
	"log"
)

func main() {
	var err error

	version, err := ioutil.ReadFile("../../version.txt")
	if err != nil {
		log.Fatalf("Failed to read ../../version.txt: %s", err)
	}

	version = bytes.TrimRight(version, "\r\n")

	versionFile, err := ioutil.ReadFile("version.go.tmpl")
	if err != nil {
		log.Fatalf("Failed to read version.go.tmpl: %s", err)
	}

	versionFile = bytes.Replace(versionFile, []byte("<VERSION>"), version, -1)

	err = ioutil.WriteFile("version.go", versionFile, 0644)
	if err != nil {
		log.Fatalf("Failed to write version.go: %s", err)
	}
}
