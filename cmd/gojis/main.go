package main

import (
	"log"
)

var (
	// Version is the version of this build, set by the linker.
	Version string
)

func main() {
	_ = Version

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
