package search

import (
	"fmt"
	"sync"
	"time"
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

	now := time.Now()

	var wg sync.WaitGroup = sync.WaitGroup{}

	// check for websites
	for _, site := range sites {

		wg.Add(1)

		if priority && !site.Priority {
			wg.Done() // done with this site
			continue
		}

		go func(site Site) {
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
			wg.Done() // done with this site

		}(site)

	}

	wg.Wait() // wait for all the goroutines to finish

	// loop through results
	for _, result := range results {
		if condition := result.Status; condition {
			fmt.Printf("%s: %s\n", result.Name, result.Link)
		}

	}

	fmt.Println("Time taken:", time.Since(now))
}
