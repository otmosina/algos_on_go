package lib

import (
	"net/url"
	"strings"
)

type Url struct {
	url    string
	Params map[string]string
}

func NewUrl(url string) *Url {
	return &Url{url: url, Params: make(map[string]string)}
}

func (u *Url) Scheme() string {
	urlObj, err := url.Parse(u.url)
	if err != nil {
		panic("Err when url parse")
	}
	return urlObj.Scheme
}

func (u *Url) Host() string {
	urlObj, err := url.Parse(u.url)
	if err != nil {
		panic("Err when url parse")
	}
	return urlObj.Host
}

func (u *Url) QueryParams() map[string]string {
	urlObj, err := url.Parse(u.url)
	if err != nil {
		panic("Err when url parse")
	}
	pairs := strings.Split(urlObj.RawQuery, "&")
	for _, pair := range pairs {
		pairArr := strings.Split(pair, "=")
		u.Params[pairArr[0]] = pairArr[1]
	}
	return u.Params
}

func (u *Url) QueryParam(k string) string {
	if _, ok := u.Params[k]; !ok {
		u.QueryParams()
	}
	return u.Params[k]
}
