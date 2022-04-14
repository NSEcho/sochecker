package twitter

import (
	"fmt"
	"net/http"

	"github.com/lateralusd/sochecker/checker"

	tsc "github.com/n0madic/twitter-scraper"
)

type TWCheck struct {
	link string
	err  error
}

func (tw *TWCheck) Check(client *http.Client, name string) bool {
	tw.link = fmt.Sprintf("https://www.twitter.com/%s/", name)

	scp := tsc.New().SetSearchMode(tsc.SearchUsers)
	_, err := scp.GetProfile(name)
	return err == nil
}

func (tw *TWCheck) Info() string {
	return "Check the username on twitter"
}

func (tw *TWCheck) Link() string {
	return tw.link
}

func (tw *TWCheck) Error() error {
	return tw.err
}

func init() {
	checker.Register("twitter", &TWCheck{})
}
