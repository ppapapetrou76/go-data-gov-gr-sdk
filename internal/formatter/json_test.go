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
	"bytes"
	"testing"

	"github.com/ppapapetrou76/go-testing/assert"
)

func TestNewJSON(t *testing.T) {
	var b bytes.Buffer
	actual := NewJSON(&b)
	expected := &JSON{
		writer:      &b,
		prefix:      "",
		indentation: "  ",
	}

	assert.That(t, actual).IsEqualTo(expected)
}

func TestJSON_Name(t *testing.T) {
	assert.That(t, NewJSON(nil).Name()).IsEqualTo("json")
}

type sampleStruct struct {
	Field1 string `json:"field_1"`
	Field2 int    `json:"field_2"`
}

func TestJSON_Format(t *testing.T) {
	var b bytes.Buffer

	tests := []struct {
		name           string
		formatter      *JSON
		data           interface{}
		expectedOutput string
	}{
		{
			name:      "should succeed un-marshaling structure",
			formatter: NewJSON(&b),
			data: sampleStruct{
				Field1: "some-data",
				Field2: 199,
			},
			expectedOutput: `{
  "field_1": "some-data",
  "field_2": 199
}
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.formatter.Format(tt.data)
			assert.ThatError(t, err).IsNil()
			assert.ThatString(t, b.String()).IsEqualTo(tt.expectedOutput)
		})
	}
}
