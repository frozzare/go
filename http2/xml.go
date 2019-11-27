package http2

import (
	"bytes"
	"encoding/xml"
	"net/http"
)

// GetXML will bind XML response from a url or return a error.
func (s *Client) GetXML(url string, target interface{}) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "text/xml")

	res, err := s.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	return xml.NewDecoder(res.Body).Decode(target)
}

// GetXML will bind XML response from a url or return a error.
func GetXML(url string, target interface{}) error {
	return NewClient(nil).GetXML(url, target)
}

// PostXML will post XML from interface value to url and read XML response from a url or return a error.
func (s *Client) PostXML(url string, v interface{}, o interface{}) error {
	buf, err := xml.Marshal(v)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(buf))
	req.Header.Set("Accept", "text/xml")
	req.Header.Set("Content-Type", "text/xml")

	res, err := s.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	return xml.NewDecoder(res.Body).Decode(o)
}

// PostXML will post XML from interface value to url and read XML response from a url or return a error.
func PostXML(url string, v interface{}, o interface{}) error {
	return NewClient(nil).PostXML(url, v, o)
}
