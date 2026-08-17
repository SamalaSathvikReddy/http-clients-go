package main

import (
	"fmt"
	"net/url"
)

func getDomainNameFromURL(rawURL string) (string, error) {
    parsedURL, err := url.Parse(rawURL)
    if err != nil {
        return "", fmt.Errorf("error finding the url: %w\n", err)
    }
    return parsedURL.Hostname(), nil
}
