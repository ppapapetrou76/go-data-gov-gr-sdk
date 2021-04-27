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
	"errors"
	"testing"

	"github.com/ppapapetrou76/go-testing/assert"
)

func TestNewText(t *testing.T) {
	var b bytes.Buffer
	actual := NewText(&b)
	expected := &Text{
		writer: &b,
	}

	assert.That(t, actual).IsEqualTo(expected)
}

func TestText_Name(t *testing.T) {
	assert.That(t, NewText(nil).Name()).IsEqualTo("text")
}

func TestText_Format(t *testing.T) {
	var b bytes.Buffer

	tests := []struct {
		name           string
		formatter      *Text
		data           interface{}
		expectedOutput string
		expectedErr    error
	}{
		{
			name:        "should return an error for un-supported type",
			formatter:   NewText(&b),
			data:        "some string",
			expectedErr: errors.New("not supported type string"),
		},
		{
			name:      "should succeed un-marshaling structure",
			formatter: NewText(&b),
			data: sampleStruct{
				Field1: "some-data",
				Field2: 199,
			},
			expectedOutput: "some-data      199            \n",
		},
		{
			name:      "should succeed un-marshaling slice",
			formatter: NewText(&b),
			data: []sampleStruct{
				{
					Field1: "some-data",
					Field2: 199,
				},
				{
					Field1: "some-more-data",
					Field2: 166,
				},
			},
			expectedOutput: "Field1         Field2         \n" +
				"some-data      199            \n" +
				"some-more-data 166            \n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer b.Reset()
			err := tt.formatter.Format(tt.data)
			assert.ThatError(t, err).IsSameAs(tt.expectedErr)
			assert.ThatString(t, b.String()).IsEqualTo(tt.expectedOutput)
		})
	}
}
