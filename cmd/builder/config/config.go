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

package config

type Config struct {
	Name       string   `mapstructure:"name"`
	Module     string   `mapstructure:"module"`
	Collectors []string `mapstructure:"collectors"`
}

// NewConfig returns a new Config.
func NewConfig() *Config {
	return &Config{
		Name:   "diy_exporter",
		Module: "diy_exporter",
		Collectors: []string{
			"sample1",
			"sample2",
		},
	}
}
