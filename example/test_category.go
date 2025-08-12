package main

import (
	"fmt"
	"log"

	lawapi "go.ngs.io/jplaw-api-v2"
)

func main() {
	// Create API client
	client := lawapi.NewClient()

	fmt.Println("Testing CategoryCd with meaningful enum names")
	fmt.Println("=============================================")

	// Test 1: Search laws by category using meaningful names
	fmt.Println("\nTest 1: Searching for Constitution category laws")
	params := &lawapi.GetLawsParams{
		CategoryCd: &[]lawapi.CategoryCd{
			lawapi.CategoryCdConstitution, // 001 - 憲法
		},
		Limit: lawapi.Int32Ptr(5),
	}

	result, err := client.GetLaws(params)
	if err != nil {
		log.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Found %d laws in Constitution category\n", result.Count)
		for i, law := range result.Laws {
			if law.RevisionInfo != nil && i < 3 {
				fmt.Printf("  - %s\n", law.RevisionInfo.LawTitle)
			}
		}
	}

	// Test 2: Search with multiple categories
	fmt.Println("\nTest 2: Searching for Criminal and Civil category laws")
	params2 := &lawapi.GetLawsParams{
		CategoryCd: &[]lawapi.CategoryCd{
			lawapi.CategoryCdCriminal, // 002 - 刑事
			lawapi.CategoryCdCivil,    // 046 - 民事
		},
		Limit: lawapi.Int32Ptr(5),
	}

	result2, err := client.GetLaws(params2)
	if err != nil {
		log.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Found %d laws in Criminal and Civil categories\n", result2.Count)
		for i, law := range result2.Laws {
			if law.RevisionInfo != nil && i < 3 {
				fmt.Printf("  - %s\n", law.RevisionInfo.LawTitle)
			}
		}
	}

	// Test 3: Using specialized categories
	fmt.Println("\nTest 3: Searching for Telecommunications category laws")
	params3 := &lawapi.GetLawsParams{
		CategoryCd: &[]lawapi.CategoryCd{
			lawapi.CategoryCdTelecommunications, // 015 - 電気通信
		},
		Limit: lawapi.Int32Ptr(5),
	}

	result3, err := client.GetLaws(params3)
	if err != nil {
		log.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Found %d laws in Telecommunications category\n", result3.Count)
		for i, law := range result3.Laws {
			if law.RevisionInfo != nil && i < 3 {
				fmt.Printf("  - %s (Category: %s)\n", 
					law.RevisionInfo.LawTitle, 
					law.RevisionInfo.Category)
			}
		}
	}

	// Demonstrate the mapping
	fmt.Println("\nCategory Code Mapping Examples:")
	fmt.Printf("  CategoryCdConstitution = %q (憲法)\n", lawapi.CategoryCdConstitution)
	fmt.Printf("  CategoryCdCriminal = %q (刑事)\n", lawapi.CategoryCdCriminal)
	fmt.Printf("  CategoryCdEducation = %q (教育)\n", lawapi.CategoryCdEducation)
	fmt.Printf("  CategoryCdDefense = %q (防衛)\n", lawapi.CategoryCdDefense)
	fmt.Printf("  CategoryCdForeignAffairs = %q (外事)\n", lawapi.CategoryCdForeignAffairs)

	fmt.Println("\nAll tests completed!")
}