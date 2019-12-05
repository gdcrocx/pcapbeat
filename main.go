package main

import (
	"os"

	"github.com/gdcrocx/pcapbeat/cmd"

	_ "github.com/gdcrocx/pcapbeat/include"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
