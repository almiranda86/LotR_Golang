package data

import (
	"io/ioutil"
	"lotr/core"
	"net/http"
)

func ListAllBooks() string {
	response, err := http.Get(core.URL_ALL_BOOKS())

	if err != nil {
		return "The HTTP request failed with error"
	}
	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "Problem while reading response body"
	}

	return string(data)
}
