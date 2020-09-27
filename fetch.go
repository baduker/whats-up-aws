package whats_up_aws

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

const (
	baseURL = "https://aws.amazon.com"
	path    = "/api/dirs/items/search?"
	query   = "item.directoryId=whats-new&item.locale=en_US&"
	sort    = "sort_by=item.additionalFields.postDateTime&sort_order=desc&"
)

func buildURL(page int, size int) string {
	v := url.Values{}
	v.Set("page", strconv.Itoa(page))
	v.Set("size", strconv.Itoa(size))

	elements := []string{baseURL, path, query, sort, v.Encode()}
	fullURL, err := url.Parse(strings.Join(elements, ""))

	if err != nil {
		panic(err)
	}

	return fullURL.String()
}

func fetch(url string) []byte {
	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	return body
}

func FetchData() []byte {
	return fetch(buildURL(25, 0))
}
