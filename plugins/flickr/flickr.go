package flickr

import (
	"fmt"
	"net/http"

	"github.com/lateralusd/sochecker/checker"
)

var link = ""

type FLCheck struct{}

func (ol *FLCheck) Check(client *http.Client, name string) bool {
	url := fmt.Sprintf("https://www.flickr.com/people/%s/", name)
	link = url
	resp, err := client.Get(url)
	if err != nil {
		fmt.Println("error", err)
		return false
	}
	defer resp.Body.Close()

	return resp.StatusCode == 200
}

func (fl *FLCheck) Info() string {
	return "Check the username on flickr"
}

func (fl *FLCheck) Link() string {
	return link
}

func init() {
	checker.Register("flickr", &FLCheck{})
}
