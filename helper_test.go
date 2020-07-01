package dingbot

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSign(t *testing.T) {
	a := assert.New(t)
	var timestamp int64 = 1593582402271
	secret := "this is secret"
	a.Equal("wxe2mfGpZT6HNHwEwSWkOhBkslhKEEbPQjyEiXx1xxw%3D", sign(timestamp, secret))
}

func TestNewRequest(t *testing.T) {
	a := assert.New(t)
	request, err := newRequest("http://test.com", []byte("test"))
	a.Nil(err)
	a.Equal("POST", request.Method)
	a.Equal([]byte("test"), readBody(request.Body))
	a.Equal("test.com", request.URL.Host)
}
