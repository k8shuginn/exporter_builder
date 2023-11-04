package build

import (
	"testing"

	"github.com/k8shuginn/exporter_builder/cmd/builder/config"
)

func TestGenerateExporter(t *testing.T) {
	cfg := config.NewConfig()
	if err := GenerateExporter(cfg); err != nil {
		t.Errorf("GenerateExporter() error = %v", err)
	}
}
