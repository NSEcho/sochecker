package main

import (
	"fmt"
	"os"

	"github.com/lateralusd/sochecker/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error ocurred: %v\n", err)
		os.Exit(1)
	}
}
