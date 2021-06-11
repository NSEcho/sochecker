package main

import (
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

func main() {
	checker.RunAll(os.Args[1])
}
