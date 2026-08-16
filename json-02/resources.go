package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
)

func getResources(url string) ([]map[string]any, error) {
    var resources []map[string]any
    
    res, err := http.Get(url)
    if err != nil {
        return resources, err
    }
    defer res.Body.Close()

    dec := json.NewDecoder(res.Body);

    if err := dec.Decode(&resources); err != nil {
        return nil, err
    }

    return resources, nil
}

func logResources(resources []map[string]any) {
    var formattedString []string
    
    for _, v := range resources {
        for key, val := range v {
            formattedString = append(formattedString, fmt.Sprintf("Key: %s - Value: %v\n", key, val))
        }
    }

    sort.Strings(formattedString)
    
    for _, str := range formattedString {
        fmt.Println(str)
    }
}
