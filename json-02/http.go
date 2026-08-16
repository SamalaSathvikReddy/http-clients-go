package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func getIssues(url string) ([]Issue, error) {
    res, err := http.Get(url)
    if err != nil {
        return nil, fmt.Errorf("error creating request %w\n", err)
    }
    defer res.Body.Close()
    var issues []Issue
    
    decoder := json.NewDecoder(res.Body)
    if err := decoder.Decode(&issues); err != nil {
        return nil, fmt.Errorf("error in json - format %w\n", err)
    }
    return issues, nil
}

func usingUnmarshal(url string) ([] Issue, error) {
    res, err := http.Get(url)
    if err != nil {
        return nil, fmt.Errorf("error creating request %w\n", err)
    }
    defer res.Body.Close()
    
    data, err := io.ReadAll(res.Body)
    if err != nil {
        return nil, fmt.Errorf("error in loading data %w\n", err)
    }
   
    var issues []Issue
    
    err = json.Unmarshal(data, &issues)
    if err != nil {
        return nil, fmt.Errorf("error in json - format %w\n", err)
    }

    return issues, nil
}
