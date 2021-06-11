package facebook

import (
	"fmt"
	"net/http"

	"github.com/lateralusd/sochecker/checker"
)

var link = ""

type FBCheck struct{}

func (fb *FBCheck) Check(name string) bool {
	url := fmt.Sprintf("https://www.facebook.com/%s/", name)
	link = url
	resp, err := http.Get(url)
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	return resp.StatusCode == 200
}

func (fb *FBCheck) Link() string {
	return link
}

func init() {
	checker.Register("facebook", &FBCheck{})
}
