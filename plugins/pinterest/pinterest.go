package pinterest

import (
	"fmt"
	"net/http"

	"github.com/lateralusd/sochecker/checker"
)

type PTCheck struct {
	link string
}

func (pt *PTCheck) Check(client *http.Client, name string) bool {
	url := fmt.Sprintf("https://www.pinterest.com/%s/", name)
	pt.link = url
	resp, err := client.Get(url)
	if err != nil {
		fmt.Println("error", err)
		return false
	}
	defer resp.Body.Close()

	return resp.StatusCode == 200
}

func (pt *PTCheck) Info() string {
	return "Check the username on pinterest"
}

func (pt *PTCheck) Link() string {
	return pt.link
}

func init() {
	checker.Register("pinterest", &PTCheck{})
}
