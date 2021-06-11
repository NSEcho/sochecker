package pinterest

import (
	"fmt"
	"net/http"
	"time"

	"github.com/lateralusd/sochecker/checker"
)

var link = ""

type PTCheck struct{}

func (ol *PTCheck) Check(name string) bool {
	client := &http.Client{
		Timeout: 3 * time.Second,
	}

	url := fmt.Sprintf("https://www.pinterest.com/%s/", name)
	link = url
	resp, err := client.Get(url)
	if err != nil {
		fmt.Println("error", err)
		return false
	}
	defer resp.Body.Close()

	return resp.StatusCode == 200
}

func (pt *PTCheck) Link() string {
	return link
}

func init() {
	checker.Register("pinterest", &PTCheck{})
}
