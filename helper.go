package dingbot

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

func sign(timestamp int64, secret string) string {
	text := fmt.Sprintf("%d\n%s", timestamp, secret)
	encoder := hmac.New(sha256.New, []byte(secret))
	encoder.Write([]byte(text))
	return url.QueryEscape(base64.StdEncoding.EncodeToString(encoder.Sum(nil)))
}

func newRequest(url string, body []byte) (*http.Request, error) {
	req, err := http.NewRequest("POST", url, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	return req, nil
}

func readBody(body io.ReadCloser) []byte {
	result, _ := ioutil.ReadAll(body)
	_ = body.Close()
	return result
}
