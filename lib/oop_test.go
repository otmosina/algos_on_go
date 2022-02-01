package lib

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	yaurl     = "http://yandex.ru?key=value&key2=value2"
	googleUrl = "https://google.com:80?a=b&c=d&lala=value"
)

func TestScheme(t *testing.T) {
	u := NewUrl(yaurl)
	require.Equal(t, "http", u.Scheme())
}

func TestHost(t *testing.T) {
	u := NewUrl(yaurl)
	require.Equal(t, "yandex.ru", u.Host())
}

func TestQueryString(t *testing.T) {
	expectedParams := map[string]string{"a": "b", "c": "d", "lala": "value"}
	u := NewUrl(googleUrl)

	require.Equal(t, expectedParams, u.QueryParams())
}

func TestQueryParam(t *testing.T) {
	// expectedParams := map[string]string{"a": "b", "c": "d", "lala": "value"}
	u := NewUrl(yaurl)

	require.Equal(t, "value", u.QueryParam("key"))
}

func TestCompare(t *testing.T) {
	u1 := NewUrl(yaurl)
	u2 := NewUrl(yaurl)

	require.True(t, reflect.DeepEqual(u1, u2))
}

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
