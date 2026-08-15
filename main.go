package main

import (
	"fmt"
	"os"

	"github.com/kannkyo/gh-init/cmd"
)

func main() {
	if err := cmd.NewRootCmd().Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
