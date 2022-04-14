package facebook

import (
	"fmt"
	"net/http"

	"github.com/lateralusd/sochecker/checker"
)

type FBCheck struct {
	link string
}

func (fb *FBCheck) Check(client *http.Client, name string) bool {
	url := fmt.Sprintf("https://www.facebook.com/%s/", name)
	fb.link = url
	resp, err := client.Get(url)
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	return resp.StatusCode == 200
}

func (fb *FBCheck) Info() string {
	return "Check the username on facebook"
}

func (fb *FBCheck) Link() string {
	return fb.link
}

func init() {
	checker.Register("facebook", &FBCheck{})
}
