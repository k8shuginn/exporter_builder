package main

import "testing"

func TestInitConfig(t *testing.T) {
	configPath = "./test/config.yaml"
	if err := initConfig(); err != nil {
		t.Errorf("initConfig() error = %v", err)
	}

	t.Log(cfg)
}
