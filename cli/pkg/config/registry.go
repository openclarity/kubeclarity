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
	"github.com/spf13/viper"

	shared "github.com/sambetts-cisco/kubeclarity/shared/v2/pkg/config"
)

const (
	RegistrySkipVerifyTlS = "REGISTRY_SKIP_VERIFY_TLS"
	RegistryUseHTTP       = "REGISTRY_USE_HTTP"
)

func setRegistryConfigDefaults() {
	viper.SetDefault(RegistrySkipVerifyTlS, false)
	viper.SetDefault(RegistryUseHTTP, false)
}

func loadRegistryConfig() *shared.Registry {
	setRegistryConfigDefaults()
	var auths shared.Auths
	if err := viper.UnmarshalKey("registry.auths", &auths); err != nil {
		auths = shared.Auths{}
	}
	return &shared.Registry{
		SkipVerifyTLS: viper.GetBool(RegistrySkipVerifyTlS),
		UseHTTP:       viper.GetBool(RegistryUseHTTP),
		Auths:         auths,
	}
}
