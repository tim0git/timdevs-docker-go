package test

import (
	"github.com/gruntwork-io/terratest/modules/docker"
	"github.com/stretchr/testify/assert"
	"testing"
)

const tag = "timdevs-docker-go"

func TestBuildsWithoutErrors(t *testing.T) {
	buildOptions := &docker.BuildOptions{
		Tags: []string{tag},
		Architectures: []string{
			"linux/amd64",
			"linux/arm64",
		},
		Load: true,
	}

	err := docker.BuildE(t, "..", buildOptions)

	if err != nil {
		errorString := err.Error()
		t.Logf("Expected no errors but got %s", errorString)
		t.Fail()
	}

	assert.Equal(t, nil, err)
}
