package helpers

import (
	"net/http"
)

func Get(url string) *http.Response {
	res, err := http.Get(url)
	if err != nil {
		return nil
	}
	return res
}
