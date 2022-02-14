package search

import (
	"fmt"
	"net/http"
)

type Site struct {
	Name         string
	URL          string
	Strategy     string
	BodyText     string
	RedirectLink string
	Priority     bool
}

type Result struct {
	Name   string
	Status bool
	Link   string
}

const userAgent = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/44.0.2403.157 Safari/537.36"

func checkStatus(username string, url string) (bool, string) {

	req, _ := http.NewRequest("GET", fmt.Sprintf("%s%s", url, username), nil)
	req.Header.Set("User-Agent", userAgent)

	client := http.Client{}
	res, err := client.Do(req)

	if err != nil {
		return false, ""
	}

	defer res.Body.Close()

	// fmt.Println(res.StatusCode)

	if condition := res.StatusCode == 200; condition {
		usernameurl := fmt.Sprintf("%s%s", url, username)
		return true, usernameurl
	} else {
		return false, ""
	}
}

var sites = []Site{
	{
		Name:         "github",
		URL:          "https://github.com/",
		Strategy:     "status",
		BodyText:     "", // if the strategy is status then no bodytext is required
		RedirectLink: "", // if the strategy is status then no redirect link is required
		Priority:     true,
	},
}

func Username(username string, priority bool) {

	fmt.Println("\nSearching...")

	var results = []Result{}

	// check for websites
	for _, site := range sites {
		if site.Priority == priority {

			switch expression := site.Strategy; expression {
			case "status":
				found, url := checkStatus(username, site.URL)
				if found {
					results = append(results, Result{Name: site.Name, Status: found, Link: url})
				}

			}
		}
	}

	// loop through results
	for _, result := range results {
		if condition := result.Status; condition {
			fmt.Printf("%s: %s\n", result.Name, result.Link)
		}

	}
}
