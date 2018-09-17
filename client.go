package plain_http_client

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
)

func NewRequest(ctx context.Context, method, url string, json interface{}) (*http.Request, error) {
	io, err := EncodeBody(json)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, url, io)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req = req.WithContext(ctx)

	return req, nil
}

func DoRequest(ctx context.Context, method, url string, json interface{}, ret interface{}) error {
	req, err := NewRequest(ctx, method, url, json)
	if err != nil {
		return err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	err = DecodeBody(res, ret)
	if err != nil {
		return err
	}

	return nil
}

func DecodeBody(resp *http.Response, out interface{}) error {
	defer resp.Body.Close()
	var r io.Reader = resp.Body
	decoder := json.NewDecoder(r)
	return decoder.Decode(out)
}

func EncodeBody(jsonMap interface{}) (io.Reader, error) {
	b, err := json.Marshal(jsonMap)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(b), nil
}

