package syft

import (
	"testing"

	//"github.com/anchore/go-logger/adapter/logrus"
	"github.com/anchore/stereoscope/pkg/image"
	//"github.com/anchore/syft/syft"
	"github.com/anchore/syft/syft/source"
	//"gotest.tools/assert"
)

func TestAnalyzer_Run(t *testing.T) {
	//cfg := logrus.Config{
	//	EnableConsole: true,
	//	FileLocation:  ".",
	//	Level:         "debug",
	//}
	//
	//logWrapper, err := logrus.New(cfg)
	//assert.NilError(t, err)
	//syft.SetLogger(logWrapper)
	input := source.Input{
		UserInput:   "registry:asia.gcr.io/gke-release-staging/cluster-proportional-autoscaler@sha256:b8f16a765eea1e53aadf59146b3ffede9d2a869d6edf7d1729fb83d5146bbefd",
		Scheme:      source.ImageScheme,
		ImageSource: image.OciRegistrySource,
		Location:    "asia.gcr.io/gke-release-staging/cluster-proportional-autoscaler@sha256:0f232ba18b63363e33f205d0242ef98324fb388434f8598c2fc8e967dca146bc",
		Platform:    "",
	}
	_, _, err := source.New(input, &image.RegistryOptions{
		InsecureSkipTLSVerify: false,
		InsecureUseHTTP:       false,
		Credentials:           []image.RegistryCredentials{},
		Platform:              "",
	}, []string{})

	if err != nil {
		t.Errorf("failed to create source analyze: %v", err)
	}
}
