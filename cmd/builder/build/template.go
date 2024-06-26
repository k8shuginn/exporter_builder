// Copyright 2024 k8shuginn exporter_builder
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package build

import (
	_ "embed"
	"text/template"
)

type CollectorValues struct {
	Name string `mapstructure:"name" yaml:"name"`
}

var (
	//go:embed template/go.mod.tmpl
	goModBytes    []byte
	goModTemplate = parseTemplate("go.mod", goModBytes)

	//go:embed template/main.go.tmpl
	mainBytes    []byte
	mainTemplate = parseTemplate("main.go", mainBytes)

	//go:embed template/collector.go.tmpl
	collectorBytes    []byte
	collectorTemplate = parseTemplate("collector.go", collectorBytes)
)

// GenerateMain generates files(main.go, collector.go, go.mod)
func parseTemplate(name string, bytes []byte) *template.Template {
	return template.Must(template.New(name).Parse(string(bytes)))
}
