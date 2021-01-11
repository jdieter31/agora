package main

import (
	"os"

	"github.com/jdieter31/agora/cmd/agorad/cmd"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd()
	if err := cmd.Execute(rootCmd); err != nil {
		os.Exit(1)
	}
}
