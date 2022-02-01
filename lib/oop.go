package lib

import (
	"net/url"
	"strings"
)

type Url struct {
	url string
}

func NewUrl(url string) *Url {
	return &Url{url: url}
}

func (u *Url) Scheme() string {
	urlObj, err := url.Parse(u.url)
	if err != nil {
		panic("Err when url parse")
	}
	return urlObj.Scheme
	// return ""
}

func (u *Url) Host() string {
	urlObj, err := url.Parse(u.url)
	if err != nil {
		panic("Err when url parse")
	}
	return urlObj.Host
	// return ""
}

func (u *Url) QueryString() map[string]string {
	result := make(map[string]string)
	// var query string
	urlObj, err := url.Parse(u.url)
	if err != nil {
		panic("Err when url parse")
	}
	pairs := strings.Split(urlObj.RawQuery, "&")
	for _, pair := range pairs {
		pairArr := strings.Split(pair, "=")
		result[pairArr[0]] = pairArr[1]
	}
	return result
}

//scheme
//host
//query_params

// require 'uri'
// require 'forwardable'

// class Url
//   # extend SingleForwardable
//   extend Forwardable
//   include Comparable

//   attr_reader :url

//   def initialize(url)
//     @url = URI(url)
//   end

//   def_delegators :@url, :scheme, :host, :to_s
//   # def_single_delegator "URI", :decode_www_form
//   def query_params
//     # require 'pry'; binding.pry
//     URI.decode_www_form(url.query)
//        .each_with_object({}) { |(k, v), hash| hash.store(k.to_sym, v) }
//   end

//   def query_param(name, value = nil)
//     params = query_params
//     return value unless params.key?(name)

//     params[name]
//   end

//   def <=>(other)
//     url.to_s <=> other.to_s
//   end
// end

// #END

// # frozen_string_literal: true

// require_relative 'test_helper'
// require_relative '../lib/url'

// class UrlTest < Minitest::Test
//   def setup
//     @yandex_url = 'http://yandex.ru?key=value&key2=value2'
//     @google_url = 'https://google.com:80?a=b&c=d&lala=value'
//   end

//   def test_yandex
//     url = Url.new @yandex_url

//     assert { url.scheme == 'http' }
//     assert { url.host == 'yandex.ru' }

//     params = { key: 'value', key2: 'value2' }

//     assert { url.query_params == params }
//     assert { url.query_param(:key) == 'value' }
//     assert { url.query_param(:key2, 'lala') == 'value2' }
//     assert { url.query_param(:new, 'ehu') == 'ehu' }
//   end

//   def test_yandex_equals?
//     url = Url.new @yandex_url

//     assert { Url.new(@yandex_url) == url }
//     assert { Url.new(@google_url) != url }
//   end

//   def test_google
//     url = Url.new @google_url

//     assert { url.scheme == 'https' }
//     assert { url.host == 'google.com' }

//     params = { a: 'b', c: 'd', lala: 'value' }

//     assert { params == url.query_params }
//     assert { url.query_param(:key).nil? }
//   end

//   def test_google_equals?
//     url = Url.new @google_url
//     assert { url == Url.new(@google_url) }
//     assert { url != Url.new('https://google.com') }
//     assert { url != Url.new(@google_url.sub('80', '443')) }
//   end
// end
