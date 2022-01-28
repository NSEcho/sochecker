package instagram

import (
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"

	"github.com/lateralusd/sochecker/checker"
	"gopkg.in/yaml.v2"
)

var link = ""

type IGCheck struct{}

func (ig *IGCheck) Check(client *http.Client, name string) bool {
	igurl := fmt.Sprintf("https://www.instagram.com/%s/", name)
	link = igurl
	igurl += "?__a=1"

	req, err := http.NewRequest("GET", igurl, nil)
	if err != nil {
		panic(err)
	}

	c, err := readCookiesFromFile()
	if err != nil {
		panic(err)
	}

	req.Header.Set("User-Agent", "Instagram 10.3.2 (iPhone7,2; iPhone OS 9_3_3; en_US; en-US; scale=2.00; 750x1334) AppleWebKit/420+")

	jar, _ := cookiejar.New(nil)
	client.Jar = jar

	client.Get("https://www.instagram.com")

	urlParsed, _ := url.Parse("https://www.instagram.com")
	jar.SetCookies(urlParsed, []*http.Cookie{
		{Name: "ds_user_id", Value: c.DsUserID},
		{Name: "sessionid", Value: c.SessionID},
		{Name: "csrftoken", Value: c.CSRFToken},
	})

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	return resp.ContentLength != 2
}

var info = `Check the username on instagram.
	You need to create .cookies file containing following cookies:
		ds_user_id
		sessionid
		csrftoken`

func (ig *IGCheck) Info() string {
	return info
}

func (ig *IGCheck) Link() string {
	return link
}

func init() {
	checker.Register("instagram", &IGCheck{})
}

type cookies struct {
	DsUserID  string `yaml:"ds_user_id"`
	SessionID string `yaml:"sessionid"`
	CSRFToken string `yaml:"csrftoken"`
}

func readCookiesFromFile() (*cookies, error) {
	f, err := os.Open("./.cookies")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	c := &cookies{}
	d := yaml.NewDecoder(f)
	if err := d.Decode(c); err != nil {
		return nil, err
	}
	return c, nil
}
