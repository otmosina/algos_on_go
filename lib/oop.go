package lib

import (
	"net/url"
	"strings"
)

type Url struct {
	url    *url.URL
	Params map[string]string
}

func NewUrl(u string) *Url {
	urlObj, err := url.Parse(u)
	if err != nil {
		panic("Err when url parse")
	}
	return &Url{url: urlObj, Params: make(map[string]string)}
}

func (u *Url) Scheme() string {
	return u.url.Scheme
}

func (u *Url) Host() string {
	return u.url.Host
}

func (u *Url) QueryParams() map[string]string {

	pairs := strings.Split(u.url.RawQuery, "&")
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
