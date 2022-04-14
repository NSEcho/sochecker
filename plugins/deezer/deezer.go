package deezer

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"

	"github.com/lateralusd/sochecker/checker"
)

type DeezerCheck struct {
	link string
	err  error
}

func (dc *DeezerCheck) Check(client *http.Client, name string) bool {
	checkurl := fmt.Sprintf("https://www.deezer.com/search/%s", name)
	dc.link = checkurl

	req, err := http.NewRequest("GET", checkurl, nil)
	if err != nil {
		dc.err = err
		return false
	}

	req.Header.Set("User-Agent", "Deezer 10.3.2 (iPhone7,2; iPhone OS 9_3_3; en_US; en-US; scale=2.00; 750x1334) AppleWebKit/420+")

	resp, err := client.Do(req)
	if err != nil {
		dc.err = err
		return false
	}
	defer resp.Body.Close()

	re := regexp.MustCompile(fmt.Sprintf(`"BLOG_NAME":"%s"`, name))

	b, _ := ioutil.ReadAll(resp.Body)
	matches := re.FindString(string(b))

	return len(matches) > 1
}

func (dc *DeezerCheck) Info() string {
	return "Check the username on deezer"
}

func (dc *DeezerCheck) Link() string {
	return dc.link
}

func (dc *DeezerCheck) Error() error {
	return dc.err
}

func init() {
	checker.Register("deezer", &DeezerCheck{})
}
