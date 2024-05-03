package main

import (
	"embed"

	"github.com/kawana77b/tmp4lig/cmd"
)

var version string = "0.0.0"

//go:embed templates/*
var fs embed.FS

func main() {
	cmd.Version = version
	cmd.Fs = fs
	cmd.Execute()
}
