package olx

import (
	"fmt"
	"net/http"
	"time"

	"github.com/lateralusd/sochecker/checker"
)

var link = ""

type OLXCheck struct{}

func (ol *OLXCheck) Check(name string) bool {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	url := fmt.Sprintf("https://api.olx.ba/profil/%s/", name)
	link = url
	resp, err := client.Get(url)
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	return resp.StatusCode == 200
}

func (ol *OLXCheck) Link() string {
	return link
}

func init() {
	checker.Register("olx", &OLXCheck{})
}
