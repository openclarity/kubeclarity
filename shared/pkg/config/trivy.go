// Copyright © 2022 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
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

import (
	"time"

	"github.com/spf13/viper"
)

const AnalyzerTrivyTimeoutSecondsDefault = 300

const (
	AnalyzerTrivyTimeoutSeconds = "ANALYZER_TRIVY_TIMEOUT_SECONDS"
	AnalyzerTrivyCacheDir       = "ANALYZER_TRIVY_CACHE_DIRECTORY"
	AnalyzerTrivyTempDir        = "ANALYZER_TRIVY_TEMP_DIRECTORY"
)

type AnalyzerTrivyConfig struct {
	Timeout  int    `yaml:"timeout" mapstructure:"timeout"`
	CacheDir string `yaml:"cache_dir" mapstructure:"cache_dir"`
	TempDir  string `yaml:"temp_dir" mapstructure:"temp_dir"`
}

func setAnalyzerTrivyConfigDefaults() {
	viper.SetDefault(AnalyzerTrivyTimeoutSeconds, AnalyzerTrivyTimeoutSecondsDefault)
	viper.SetDefault(AnalyzerTrivyTempDir, "/tmp")
}

func LoadAnalyzerTrivyConfig() AnalyzerTrivyConfig {
	setAnalyzerTrivyConfigDefaults()

	return AnalyzerTrivyConfig{
		Timeout:  viper.GetInt(AnalyzerTrivyTimeoutSeconds),
		CacheDir: viper.GetString(AnalyzerTrivyCacheDir),
		TempDir:  viper.GetString(AnalyzerTrivyTempDir),
	}
}

type AnalyzerTrivyConfigEx struct {
	Timeout  time.Duration
	CacheDir string
	TempDir  string
	Registry *Registry
}

func CreateAnalyzerTrivyConfigEx(analyzer *Analyzer, registry *Registry) AnalyzerTrivyConfigEx {
	return AnalyzerTrivyConfigEx{
		Timeout:  time.Duration(analyzer.TrivyConfig.Timeout) * time.Second,
		CacheDir: analyzer.TrivyConfig.CacheDir,
		TempDir:  analyzer.TrivyConfig.TempDir,
		Registry: registry,
	}
}

const ScannerTrivyTimeoutSecondsDefault = 300

const (
	ScannerTrivyTimeoutSeconds = "SCANNER_TRIVY_TIMEOUT_SECONDS"
	ScannerTrivyCacheDir       = "SCANNER_TRIVY_CACHE_DIRECTORY"
	ScannerTrivyTempDir        = "SCANNER_TRIVY_TEMP_DIRECTORY"
	ScannerTrivyServerAddress  = "SCANNER_TRIVY_SERVER_ADDRESS"
	ScannerTrivyServerToken    = "SCANNER_TRIVY_SERVER_TOKEN" // nolint:gosec
)

type ScannerTrivyConfig struct {
	Timeout     int
	ServerAddr  string
	ServerToken string
	CacheDir    string
	TempDir     string
}

func setScannerTrivyConfigDefaults() {
	viper.SetDefault(ScannerTrivyTimeoutSeconds, ScannerTrivyTimeoutSecondsDefault)
	viper.SetDefault(ScannerTrivyTempDir, "/tmp")
}

func LoadScannerTrivyConfig() ScannerTrivyConfig {
	setScannerTrivyConfigDefaults()

	return ScannerTrivyConfig{
		Timeout:     viper.GetInt(ScannerTrivyTimeoutSeconds),
		CacheDir:    viper.GetString(ScannerTrivyCacheDir),
		TempDir:     viper.GetString(ScannerTrivyTempDir),
		ServerAddr:  viper.GetString(ScannerTrivyServerAddress),
		ServerToken: viper.GetString(ScannerTrivyServerToken),
	}
}

type ScannerTrivyConfigEx struct {
	Timeout     time.Duration
	CacheDir    string
	TempDir     string
	ServerAddr  string
	ServerToken string
	Registry    *Registry
}

func CreateScannerTrivyConfigEx(scanner *Scanner, registry *Registry) ScannerTrivyConfigEx {
	return ScannerTrivyConfigEx{
		Timeout:     time.Duration(scanner.TrivyConfig.Timeout) * time.Second,
		Registry:    registry,
		ServerAddr:  scanner.TrivyConfig.ServerAddr,
		ServerToken: scanner.TrivyConfig.ServerToken,
		TempDir:     scanner.TrivyConfig.TempDir,
		CacheDir:    scanner.TrivyConfig.CacheDir,
	}
}
