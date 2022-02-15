package search

import (
	"fmt"
)

type Result struct {
	Name   string
	Status bool
	Link   string
}

type Results []Result

func Username(username string, priority bool) {

	fmt.Println("\nSearching...")

	var results Results

	// check for websites
	for _, site := range sites {

		if priority && !site.Priority {
			continue
		}

		switch expression := site.Strategy; expression {
		case "status":
			found, url := check.status(username, site)
			if found {
				results.add(Result{Name: site.Name, Status: found, Link: url})
			}
		case "bodytext":
			found, url := check.bodyText(username, site)
			if found {
				results.add(Result{Name: site.Name, Status: found, Link: url})
			}
		case "redirect":
			found, url := check.redirect(username, site)
			if found {
				results.add(Result{Name: site.Name, Status: found, Link: url})
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
