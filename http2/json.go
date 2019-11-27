package http2

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// GetJSON will bind JSON response from a url or return a error.
func (s *Client) GetJSON(url string, v interface{}) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	res, err := s.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	return json.NewDecoder(res.Body).Decode(v)
}

// GetJSON will bind JSON response from a url or return a error.
func GetJSON(url string, v interface{}) error {
	return NewClient(nil).GetJSON(url, v)
}

// PostJSON will post JSON from interface value to url and read JSON response from a url or return a error.
func (s *Client) PostJSON(url string, v interface{}, o interface{}) error {
	buf, err := json.Marshal(v)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(buf))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	res, err := s.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	return json.NewDecoder(res.Body).Decode(o)
}

// PostJSON will post JSON from interface value to url and read JSON response from a url or return a error.
func PostJSON(url string, v interface{}, o interface{}) error {
	return NewClient(nil).PostJSON(url, v, o)
}
