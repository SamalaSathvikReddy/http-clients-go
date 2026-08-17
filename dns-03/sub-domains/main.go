package main

import "log"

func main() {
    issues, err := getIssues(domain)
    if err != nil {
        log.Fatalf("error getting issues with data : %v", err)
    }
    logIssues(issues)
}

