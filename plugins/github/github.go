package github

import (
	"fmt"
	"net/http"
	"time"

	"github.com/lateralusd/sochecker/checker"
)

var link = ""

type GHCheck struct{}

func (gh *GHCheck) Check(name string) bool {
	client := &http.Client{
		Timeout: 3 * time.Second,
	}

	url := fmt.Sprintf("https://github.com/%s/", name)
	link = url
	resp, err := client.Get(url)
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	return resp.StatusCode == 200
}

func (gh *GHCheck) Link() string {
	return link
}

func init() {
	checker.Register("github", &GHCheck{})
}
