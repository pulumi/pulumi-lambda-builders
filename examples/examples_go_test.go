// Copyright 2016-2021, Pulumi Corporation.  All rights reserved.
//go:build go || all
// +build go all

package examples

import (
	"path/filepath"
	"testing"

	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
	"github.com/stretchr/testify/assert"
)

func TestSimpleGo(t *testing.T) {
	test := getGoBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: filepath.Join(getCwd(t), "simple-go"),
			ExtraRuntimeValidation: func(t *testing.T, stack integration.RuntimeValidationStackInfo) {
				out, ok := stack.Outputs["result"].(string)
				assert.True(t, ok)
				assert.Equal(t, "\"Hello World!\"", out)
			},
		})
	integration.ProgramTest(t, &test)
}

func getGoBaseOptions(t *testing.T) integration.ProgramTestOptions {
	base := getBaseOptions(t)
	baseGo := base.With(integration.ProgramTestOptions{
		Verbose: true,
		Dependencies: []string{
			"github.com/pulumi/pulumi-lambda-builders/sdk",
		},
	})
	return baseGo
}
