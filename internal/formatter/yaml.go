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
	"fmt"
	"io"

	"gopkg.in/yaml.v3"
)

const (
	yamlName = "yaml"
)

// NewYAML acts as the factory for formatter.YAML.
func NewYAML(writer io.Writer) *YAML {
	return &YAML{
		writer: writer,
	}
}

// YAML formats into yaml.
type YAML struct {
	writer io.Writer
}

// Name obtains the name of the formatter.
func (f *YAML) Name() string {
	return yamlName
}

// Format formats data as yaml output.
func (f *YAML) Format(data interface{}) error {
	r, err := yaml.Marshal(data)
	if err != nil {
		return fmt.Errorf("yaml formatter:%w", err)
	}
	fmt.Fprintln(f.writer, string(r))

	return nil
}
