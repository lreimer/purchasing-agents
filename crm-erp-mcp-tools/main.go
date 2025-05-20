package main

import "github.com/lreimer/purchasing-agents/cmd"

var version string

func main() {
	cmd.SetVersion(version)
	cmd.Execute()
}
