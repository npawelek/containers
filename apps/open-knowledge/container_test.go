package main

import (
	"testing"

	helpers "github.com/npawelek/containers/tests"
)

func Test(t *testing.T) {
	image := helpers.GetTestImage("ghcr.io/npawelek/open-knowledge:rolling")

	helpers.RequireCommandSucceeds(t, image, nil, "git", "--version")
	helpers.RequireCommandSucceeds(t, image, nil, "ok", "--version")

	helpers.RequireHTTPEndpoint(t, image, helpers.HTTPTestConfig{
		Port: "8080",
		Path: "/readyz",
	}, &helpers.ContainerConfig{
		Env: map[string]string{
			"OK_ALLOW_EXTERNAL": "1",
			"OK_EXTERNAL_URL":   "https://ok.example",
		},
	})
}
