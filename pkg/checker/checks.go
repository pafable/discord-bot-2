package check

import (
	"log"
	"net/http"
)

func UrlCheck(url string) (int, error) {
	resp, err := http.Get("https://" + url)

	if err != nil {
		return 1, err
	}

	log.Printf("%s : %d", url, resp.StatusCode)
	return resp.StatusCode, nil
}
