package main

import (
	"os"

	"github.com/IshanGhelaniCrest/countbeat/cmd"

	_ "github.com/IshanGhelaniCrest/countbeat/include"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
