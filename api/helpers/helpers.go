package helpers

import (
	"os"
	"strings"
)

func EnforceHttp (url string) string {
	if url[:4] != "http" {
		return "http://" + url
	}

	return url 
}

func RemoveDomainError(url string) bool {
	if url == os.Getenv("Domain") {
		return false
	}	

	newURL := strings.Replace(url, "http://", "", 1)
	newURL = strings.Replace(newURL, "https://", "", 1)
	newURL = strings.Replace(newURL, "www.", "", 1)
	newURL = strings.Split(newURL, "/")[0]

	if newURL == os.Getenv("Domain") {
		return false
	}

	return true
}
