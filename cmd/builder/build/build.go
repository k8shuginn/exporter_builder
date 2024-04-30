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
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/k8shuginn/exporter_builder/cmd/builder/config"
)

func GenerateExporter(cfg *config.Config) error {
	cfg.Collectors = removeDuplicateCollector(cfg.Collectors)
	if err := createDirectory(cfg.Name, cfg.Collectors); err != nil {
		return err
	}
	if err := generateMain(cfg); err != nil {
		return err
	}
	if err := generateCollector(cfg); err != nil {
		return err
	}
	if err := generateGoMod(cfg); err != nil {
		return err
	}

	return nil
}

func removeDuplicateCollector(collectors []string) []string {
	keys := make(map[string]bool)
	var result []string

	for _, collector := range collectors {
		if _, value := keys[collector]; !value {
			keys[collector] = true
			result = append(result, collector)
		}
	}

	return result
}

func createDirectory(name string, collectors []string) error {
	if _, err := os.Stat(name); os.IsNotExist(err) {
		if err := os.Mkdir(name, 0755); err != nil {
			return fmt.Errorf("failed create directory: %s", err)
		}
	}

	for _, collector := range collectors {
		if _, err := os.Stat(fmt.Sprintf("%s/collector/%s", name, collector)); os.IsNotExist(err) {
			if err := os.MkdirAll(fmt.Sprintf("%s/collector/%s", name, collector), 0755); err != nil {
				return fmt.Errorf("failed create collector directory: %s", err)
			}
		}
	}

	return nil
}

func generateMain(cfg *config.Config) error {
	out, err := os.Create(filepath.Clean(filepath.Join(cfg.Name, mainTemplate.Name())))
	if err != nil {
		return fmt.Errorf("failed create main file: %s", err)
	}

	return mainTemplate.Execute(out, cfg)
}

func generateCollector(cfg *config.Config) error {
	for _, collector := range cfg.Collectors {
		out, err := os.Create(filepath.Clean(filepath.Join(cfg.Name, "collector", collector, collectorTemplate.Name())))
		if err != nil {
			return fmt.Errorf("failed create collector file: %s", err)
		}

		if err := collectorTemplate.Execute(out, CollectorValues{Name: collector}); err != nil {
			return fmt.Errorf("failed execute collector template: %s", err)
		}
	}

	return nil
}

func generateGoMod(cfg *config.Config) error {
	out, err := os.Create(filepath.Clean(filepath.Join(cfg.Name, goModTemplate.Name())))
	if err != nil {
		return fmt.Errorf("failed create go.mod file: %s", err)
	}
	if err = goModTemplate.Execute(out, cfg); err != nil {
		return fmt.Errorf("failed execute go.mod template: %s", err)
	}

	tidyCmd := exec.Command("go", "mod", "tidy")
	tidyCmd.Dir = cfg.Name
	if result, err := tidyCmd.CombinedOutput(); err != nil {
		return fmt.Errorf("failed execute go.mod tidy: %s", result)
	}

	vendorCmd := exec.Command("go", "mod", "vendor")
	vendorCmd.Dir = cfg.Name
	if result, err := vendorCmd.CombinedOutput(); err != nil {
		return fmt.Errorf("failed execute go.mod vendor: %s", result)
	}

	return nil
}
