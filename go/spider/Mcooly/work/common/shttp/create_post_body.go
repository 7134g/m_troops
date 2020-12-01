package shttp

import (
	"encoding/json"
	"io"
	"net/url"
	"strings"
)

func CreatePostFormReader(data map[string]string) io.Reader {
	form := url.Values{}
	for k, v := range data {
		form.Add(k, v)
	}
	return strings.NewReader(form.Encode())
}

func CreatePostRawReader(data interface{}) (io.Reader, error) {
	body, err := json.Marshal(data)
	return strings.NewReader(string(body)), err
}
