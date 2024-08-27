package main

import (
	"github.com/CalmaIndustry/calma-kcheck/cmd"
)

var (
	// version is set during build
	version = "development"
	// commit is set during build
	commit = "n/a"
)

func main() {
	cmd.Execute(version, commit)
}
