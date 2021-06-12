package olx

import (
	"fmt"
	"net/http"

	"github.com/lateralusd/sochecker/checker"
)

var link = ""

type OLXCheck struct{}

func (ol *OLXCheck) Check(client *http.Client, name string) bool {
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
