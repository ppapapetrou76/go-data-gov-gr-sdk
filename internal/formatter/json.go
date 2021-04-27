// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package formatter

import (
	"encoding/json"
	"fmt"
	"io"
)

const (
	jsonName           = "json"
	defaultPrefix      = ""
	defaultIndentation = "  "
)

// NewJSON acts as the factory for formatter.JSON.
func NewJSON(writer io.Writer) *JSON {
	return &JSON{
		writer:      writer,
		prefix:      defaultPrefix,
		indentation: defaultIndentation,
	}
}

// JSON formats into text.
type JSON struct {
	writer              io.Writer
	prefix, indentation string
}

// Name obtains the name of the formatter.
func (f *JSON) Name() string {
	return jsonName
}

// Format formats data as json output.
func (f *JSON) Format(data interface{}) error {
	r, err := json.MarshalIndent(data, f.prefix, f.indentation)
	if err != nil {
		return fmt.Errorf("json formatter:%w", err)
	}
	fmt.Fprintln(f.writer, string(r))

	return nil
}
