package github

import (
	"fmt"
	"net/http"

	"github.com/lateralusd/sochecker/checker"
)

type GHCheck struct {
	link string
	err  error
}

func (gh *GHCheck) Check(client *http.Client, name string) bool {
	url := fmt.Sprintf("https://github.com/%s/", name)
	gh.link = url
	resp, err := client.Get(url)
	if err != nil {
		gh.err = err
		return false
	}
	defer resp.Body.Close()

	return resp.StatusCode == 200
}

func (gh *GHCheck) Info() string {
	return "Check the username on github"
}

func (gh *GHCheck) Link() string {
	return gh.link
}

func (gh *GHCheck) Error() error {
	return gh.err
}

func init() {
	checker.Register("github", &GHCheck{})
}
