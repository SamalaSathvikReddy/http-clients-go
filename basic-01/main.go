package main

import (
	"fmt"
	"log"
)

const issueURL = "https://api.boot.dev/v1/courses_rest_api/learn-http/issues"

func main() {
	issues, err := getIssueData(issueURL)
	if err != nil {
		log.Fatalf("error getting issue data: %v", err)
	}
	issuesString := string(issues)
	// fmt.Println(issuesString)

	prettyData, err := prettify(issuesString)
	if err != nil {
		log.Fatalf("error prettifying the data %v\n", err)
	}
	fmt.Println(prettyData)
}
