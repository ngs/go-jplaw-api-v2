package main

import (
	"fmt"
	"log"

	lawapi "go.ngs.io/jplaw-api-v2"
)

func main() {
	// Create API client
	client := lawapi.NewClient()

	fmt.Printf("Japan Law API v2 Client Library Test\n")
	fmt.Printf("Base URL: %s\n", lawapi.DefaultBaseURL)

	// Example: Get laws list
	params := &lawapi.GetLawsParams{
		LawTitle: lawapi.StringPtr("電波法"),
		// Limit:    lawapi.Int32Ptr(10),
	}

	fmt.Printf("Fetching laws list...\n")
	result, err := client.GetLaws(params)
	if err != nil {
		log.Printf("Error: %v\n", err)
		return
	} else {
		fmt.Printf("Success: %+v\n", result)
	}

	for _, law := range result.Laws {
		if law.LawInfo != nil {
			fmt.Printf("Law ID: %s\n", law.LawInfo.LawId)
		}
		if law.RevisionInfo != nil {
			fmt.Printf("Law Title: %s\n", law.RevisionInfo.LawTitle)
		}
	}
}
