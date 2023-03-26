// Copyright 2018 Google LLC All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package rest

import (
	"strings"
	"testing"

	name2 "github.com/google/go-containerregistry/pkg/name"
)

type stringConst string

const validDigest = "sha256:deadb33fdeadb33fdeadb33fdeadb33fdeadb33fdeadb33fdeadb33fdeadb33f"

var goodStrictValidationDigestNames = []string{
	"gcr.io/g-convoy/hello-world@" + validDigest,
	"gcr.io/google.com/project-id/hello-world@" + validDigest,
	"us.gcr.io/project-id/sub-repo@" + validDigest,
	"example.text/foo/bar@" + validDigest,
}

var goodStrictValidationTagDigestNames = []string{
	"example.text/foo/bar:latest@" + validDigest,
	"example.text:8443/foo/bar:latest@" + validDigest,
	"example.text/foo/bar:v1.0.0-alpine@" + validDigest,
}

var goodWeakValidationDigestNames = []string{
	"namespace/pathcomponent/image@" + validDigest,
	"library/ubuntu@" + validDigest,
	"registry:docker.io/istio/examples-bookinfo-reviews-v1@sha256:b8f16a765eea1e53aadf59146b3ffede9d2a869d6edf7d1729fb83d5146bbefd",
}

var goodWeakValidationTagDigestNames = []string{
	"nginx:latest@" + validDigest,
	"library/nginx:latest@" + validDigest,
}

var badDigestNames = []string{
	"gcr.io/project-id/unknown-alg@unknown:abc123",
	"gcr.io/project-id/wrong-length@sha256:d34db33fd34db33f",
	"gcr.io/project-id/missing-digest@",
	// https://github.com/google/go-containerregistry/issues/1394
	"repo@sha256:" + strings.Repeat(":", 64),
	"repo@sha256:" + strings.Repeat("sh", 32),
	"repo@sha256:" + validDigest + "@" + validDigest,
}

var goodStrictValidationTagNames = []string{
	"gcr.io/g-convoy/hello-world:latest",
	"gcr.io/google.com/g-convoy/hello-world:latest",
	"gcr.io/project-id/with-nums:v2",
	"us.gcr.io/project-id/image:with.period.in.tag",
	"gcr.io/project-id/image:w1th-alpha_num3ric.PLUScaps",
	"domain.with.port:9001/image:latest",
}

var goodWeakValidationTagNames = []string{
	"namespace/pathcomponent/image",
	"library/ubuntu",
	"gcr.io/project-id/implicit-latest",
	"www.example.test:12345/repo/path",
}

var badTagNames = []string{
	"gcr.io/project-id/bad_chars:c@n'tuse",
	"gcr.io/project-id/wrong-length:white space",
	"gcr.io/project-id/too-many-chars:thisisthetagthatneverendsitgoesonandonmyfriendsomepeoplestartedtaggingitnotknowingwhatitwasandtheyllcontinuetaggingitforeverjustbecausethisisthetagthatneverends",
}

var (
	testDefaultRegistry = "registry.upbound.io"
	testDefaultTag      = "stable"
	inputDefaultNames   = []string{
		"crossplane/provider-gcp",
		"crossplane/provider-gcp:v0.14.0",
		"ubuntu",
		"gcr.io/crossplane/provider-gcp:latest",
	}
	outputDefaultNames = []string{
		"registry.upbound.io/crossplane/provider-gcp:stable",
		"registry.upbound.io/crossplane/provider-gcp:v0.14.0",
		"registry.upbound.io/ubuntu:stable",
		"gcr.io/crossplane/provider-gcp:latest",
	}
)

func TestParseReferenceDefaulting(t *testing.T) {
	for i, name := range inputDefaultNames {
		ref, err := name2.ParseReference(name, name2.WithDefaultRegistry(testDefaultRegistry), name2.WithDefaultTag(testDefaultTag))
		if err != nil {
			t.Errorf("ParseReference(%q); %v", name, err)
		}
		if ref.Name() != outputDefaultNames[i] {
			t.Errorf("ParseReference(%q); got %v, want %v", name, ref.String(), outputDefaultNames[i])
		}
	}
}

func TestParseReference(t *testing.T) {
	for _, name := range goodWeakValidationDigestNames {
		ref, err := name2.ParseReference(name, name2.WeakValidation)
		if err != nil {
			t.Errorf("ParseReference(%q); %v", name, err)
		}
		dig, err := name2.NewDigest(name, name2.WeakValidation)
		if err != nil {
			t.Errorf("NewDigest(%q); %v", name, err)
		}
		if ref != dig {
			t.Errorf("ParseReference(%q) != NewDigest(%q); got %v, want %v", name, name, ref, dig)
		}
	}

	for _, name := range goodStrictValidationDigestNames {
		ref, err := name2.ParseReference(name, name2.StrictValidation)
		if err != nil {
			t.Errorf("ParseReference(%q); %v", name, err)
		}
		dig, err := name2.NewDigest(name, name2.StrictValidation)
		if err != nil {
			t.Errorf("NewDigest(%q); %v", name, err)
		}
		if ref != dig {
			t.Errorf("ParseReference(%q) != NewDigest(%q); got %v, want %v", name, name, ref, dig)
		}
	}

	for _, name := range badDigestNames {
		if _, err := name2.ParseReference(name, name2.WeakValidation); err == nil {
			t.Errorf("ParseReference(%q); expected error, got none", name)
		}
	}

	for _, name := range goodWeakValidationTagNames {
		ref, err := name2.ParseReference(name, name2.WeakValidation)
		if err != nil {
			t.Errorf("ParseReference(%q); %v", name, err)
		}
		tag, err := name2.NewTag(name, name2.WeakValidation)
		if err != nil {
			t.Errorf("NewTag(%q); %v", name, err)
		}
		if ref != tag {
			t.Errorf("ParseReference(%q) != NewTag(%q); got %v, want %v", name, name, ref, tag)
		}
	}

	for _, name := range goodStrictValidationTagNames {
		ref, err := name2.ParseReference(name, name2.StrictValidation)
		if err != nil {
			t.Errorf("ParseReference(%q); %v", name, err)
		}
		tag, err := name2.NewTag(name, name2.StrictValidation)
		if err != nil {
			t.Errorf("NewTag(%q); %v", name, err)
		}
		if ref != tag {
			t.Errorf("ParseReference(%q) != NewTag(%q); got %v, want %v", name, name, ref, tag)
		}
	}

	for _, name := range badTagNames {
		if _, err := name2.ParseReference(name, name2.WeakValidation); err == nil {
			t.Errorf("ParseReference(%q); expected error, got none", name)
		}
	}
}

// Test that MustParseReference can accept a const string or string value.
const str = "valid/string"

var _ = name2.MustParseReference(str)
var _ = name2.MustParseReference("valid/string")
var _ = name2.MustParseReference("valid/prefix/" + str)
