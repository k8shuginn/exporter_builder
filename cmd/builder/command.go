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

func initConfig() error {
	if err := k.Load(file.Provider(configPath), yaml.Parser()); err != nil {
		return fmt.Errorf("error loading config: %v", err)
	}
	if err := k.Unmarshal("", cfg); err != nil {
		return fmt.Errorf("error unmarshaling config: %v", err)
	}

	return nil
}
