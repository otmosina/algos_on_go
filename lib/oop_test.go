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
	u := NewUrl(yaurl)

	require.Equal(t, "value", u.QueryParam("key"))
	require.Equal(t, "", u.QueryParam("keyUnderfined"))
}

func TestCompare(t *testing.T) {
	u1 := NewUrl(yaurl)
	u2 := NewUrl(yaurl)

	require.True(t, reflect.DeepEqual(u1, u2))

	u2 = NewUrl(googleUrl)
	require.False(t, reflect.DeepEqual(u1, u2))
}
