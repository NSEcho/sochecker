package reddit

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/lateralusd/sochecker/checker"
)

type RedditCheck struct {
	link string
}

func (rc *RedditCheck) Check(client *http.Client, name string) bool {
	checkurl := fmt.Sprintf("https://www.reddit.com/user/%s/", name)
	rc.link = checkurl

	req, err := http.NewRequest("GET", checkurl, nil)
	if err != nil {
		return false
	}

	req.Header.Set("User-Agent", "Instagram 10.3.2 (iPhone7,2; iPhone OS 9_3_3; en_US; en-US; scale=2.00; 750x1334) AppleWebKit/420+")

	resp, err := client.Do(req)
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)
	return !strings.Contains(string(b), "nobody on Reddit")
}

func (rc *RedditCheck) Info() string {
	return "Check the username on reddit"
}

func (rc *RedditCheck) Link() string {
	return rc.link
}

func init() {
	checker.Register("reddit", &RedditCheck{})
}
