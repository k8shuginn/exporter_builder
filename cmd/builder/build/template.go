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

func parseTemplate(name string, bytes []byte) *template.Template {
	return template.Must(template.New(name).Parse(string(bytes)))
}
