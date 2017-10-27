// jobsearch parses HTML pages to obtain job information
package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"strings"
)

// get all positions in Dropbox, and their locations
func getDropbox() {
	doc, err := goquery.NewDocument("https://www.dropbox.com/jobs/all-jobs")
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	doc.Find(".open-positions__listing").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		job := s.Find("h5").Text()
		location := s.Find("p").Text()
		// TODO: Make this RegEx to be a bit better.
		if strings.Contains(job, "Engineer") { // TODO: Allow search from cmd line
			fmt.Printf("%s - %s\n", job, location)
		}
	})
}

func main() {
	getDropbox()
}
