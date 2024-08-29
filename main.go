package main

import (
	cmd "calma-kcheck/cmd"
)

var (
	version = "dev"
)

func main() {
	cmd.Version(version)
	cmd.Helm()
}
