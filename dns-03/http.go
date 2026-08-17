package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func getIPAddress(domain string) (string ,error) {
    url := fmt.Sprintf("https://1.1.1.1/dns-query?name=%s&type=A", domain)
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return "", fmt.Errorf("error creating request: %w\n", err)
    }
    req.Header.Set("accept", "application/dns-json")

    client := http.Client{}
    res, err := client.Do(req)
    if err != nil {
        return "", fmt.Errorf("error sending request: %w\n", err)
    }
    defer res.Body.Close()

    body, err := io.ReadAll(res.Body)
    if err != nil {
        return "", fmt.Errorf("error reading response body: %w\n", err)
    }
    
    // decode the body
    // body is in json-format
    
    dnsresponse := DNSResponse{}

    err = json.Unmarshal(body, &dnsresponse)

    if err != nil {
        return "", fmt.Errorf("error in unpacking json: %w\n", err)
    }
    
    if len(dnsresponse.Answer) == 0 {
        return "", fmt.Errorf("error no response received\n")
    }
    
    return dnsresponse.Answer[0].Data, nil
}
