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

package main

import (
	"fmt"
	"log"

	"github.com/k8shuginn/exporter_builder/cmd/builder/build"
	"github.com/k8shuginn/exporter_builder/cmd/builder/config"
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/spf13/cobra"
)

const (
	ExampleMessage = "builer --config config.yaml"
)

var (
	configPath string
	cfg        = config.NewConfig()
	k          = koanf.New(".")
)

// Command create a new builder command
func Command() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:     "builder",
		Short:   "builder is a command line tool to generate exporter",
		Example: ExampleMessage,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := initConfig(); err != nil {
				return err
			}
			if err := build.GenerateExporter(cfg); err != nil {
				return err
			}

			log.Println("build exporter success")
			return nil
		},
	}
	cmd.Flags().StringVar(&configPath, "config", "./config.yaml", "config file path")
	err := cmd.Flags().MarkDeprecated("config", "please use --config-file instead")

	return cmd, err
}

// initConfig load config from file
func initConfig() error {
	if err := k.Load(file.Provider(configPath), yaml.Parser()); err != nil {
		return fmt.Errorf("error loading config: %v", err)
	}
	if err := k.Unmarshal("", cfg); err != nil {
		return fmt.Errorf("error unmarshaling config: %v", err)
	}

	return nil
}
