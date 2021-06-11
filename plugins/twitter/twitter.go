package twitter

import (
	"fmt"

	"github.com/lateralusd/sochecker/checker"

	tsc "github.com/n0madic/twitter-scraper"
)

var link = ""

type TWCheck struct{}

func (tw *TWCheck) Check(name string) bool {
	link = fmt.Sprintf("https://www.twitter.com/%s/", name)

	scp := tsc.New().SetSearchMode(tsc.SearchUsers)
	_, err := scp.GetProfile(name)
	return err == nil
}

func (tw *TWCheck) Link() string {
	return link
}

func init() {
	checker.Register("twitter", &TWCheck{})
}
