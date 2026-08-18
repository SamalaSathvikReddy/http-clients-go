package main

import (
	"fmt"
	"net/url"
)

func newParsedURL(urlstring string) ParsedURL {
    parsedURL, err := url.Parse(urlstring)
    if err != nil {
        return ParsedURL{}
    }
    var pass string
    _, f1 := parsedURL.User.Password() 
    if f1 {
        pass, _ = parsedURL.User.Password()
    }
    return ParsedURL{
        protocol: parsedURL.Scheme,
        username: parsedURL.User.Username(),
        password: pass,
        hostname: parsedURL.Hostname(),
        port: parsedURL.Port(),
        pathname: parsedURL.Path,
        search: parsedURL.RawQuery,
        hash: parsedURL.Fragment,
    }
}

func getMailtoLinkForEmail(email string) string {
    return fmt.Sprintf("mailto:%s\n", email)
}
