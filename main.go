package main

import (
	"fmt"
	"os"

	"github.com/lateralusd/sochecker/checker"

	_ "github.com/lateralusd/sochecker/plugins/facebook"
	_ "github.com/lateralusd/sochecker/plugins/flickr"
	_ "github.com/lateralusd/sochecker/plugins/github"
	_ "github.com/lateralusd/sochecker/plugins/instagram"
	_ "github.com/lateralusd/sochecker/plugins/olx"
	_ "github.com/lateralusd/sochecker/plugins/pinterest"
	_ "github.com/lateralusd/sochecker/plugins/twitter"
)

var message = `You need to provide username

Usage of sochecker:
	sochecker <username>
Example:
	sochecker testUserName`

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, message)
		os.Exit(1)
	}

	checker.RunAll(os.Args[1])
}
