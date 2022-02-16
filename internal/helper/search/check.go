package search

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Checker func(string, Site) (bool, string)

type Check struct {
	status   Checker
	bodyText Checker
	redirect Checker
}

var check = Check{
	status:   checkStatus,
	bodyText: checkBodyText,
	redirect: checkRedirect,
}

func (results *Results) add(result Result) {
	*results = append(*results, result)
}

const userAgent = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/44.0.2403.157 Safari/537.36"

func checkStatus(username string, site Site) (bool, string) {

	req, _ := http.NewRequest("GET", fmt.Sprintf(site.URL, username), nil)
	req.Header.Set("User-Agent", userAgent)

	client := http.Client{}
	res, err := client.Do(req)

	if err != nil {
		return false, ""
	}

	defer res.Body.Close()

	if condition := res.StatusCode == 200; condition {
		usernameurl := fmt.Sprintf(site.URL, username)
		return true, usernameurl
	} else {
		return false, ""
	}
}

func checkBodyText(username string, site Site) (bool, string) {

	usernameurl := fmt.Sprintf(site.URL, username)

	req, _ := http.NewRequest("GET", usernameurl, nil)
	req.Header.Set("User-Agent", userAgent)

	client := http.Client{}
	res, err := client.Do(req)

	if err != nil {
		return false, ""
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return false, ""
	}

	contains := strings.Contains(string(body), site.BodyText)

	if contains {
		return false, ""
	} else {
		return true, usernameurl
	}
}

func checkRedirect(username string, site Site) (bool, string) {

	usernameurl := fmt.Sprintf(site.URL, username)

	req, _ := http.NewRequest("GET", usernameurl, nil)
	req.Header.Set("User-Agent", userAgent)

	client := http.Client{}
	res, err := client.Do(req)

	if err != nil {
		return false, ""
	}

	// defer res.Body.Close()

	var errorURL string

	if strings.Contains(site.RedirectLink, "%s") {
		errorURL = fmt.Sprintf(site.RedirectLink, username)
	} else {
		errorURL = site.RedirectLink
	}

	contains := strings.Contains(res.Header.Get("Locations"), errorURL)

	if contains {

		return true, usernameurl

	} else {

		return false, ""

	}

}
