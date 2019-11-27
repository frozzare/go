package http2

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/frozzare/go-assert"
)

type XMLPerson struct {
	XMLName xml.Name `xml:"name"`
	Name    string   `xml:",chardata"`
}

func TestGetXML(t *testing.T) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)

	mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<name>Fredrik</name>")
	}))

	actual := &XMLPerson{}
	err := GetXML(server.URL, &actual)

	assert.Nil(t, err)
	assert.Equal(t, "Fredrik", actual.Name)
}

func TestGetXMLResponseError(t *testing.T) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	actual := &XMLPerson{}
	err := GetXML(server.URL, &actual)

	assert.NotNil(t, err)
	assert.Equal(t, "", actual.Name)
}

func TestPostXML(t *testing.T) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)

	mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var v *XMLPerson

		err := xml.NewDecoder(r.Body).Decode(&v)
		assert.Nil(t, err)
		assert.Equal(t, "Input", v.Name)

		fmt.Fprintf(w, "<name>Fredrik</name>")
	}))

	var actual *XMLPerson
	err := PostXML(server.URL, &XMLPerson{Name: "Input"}, &actual)

	assert.Nil(t, err)
	assert.Equal(t, "Fredrik", actual.Name)
}

func TestPostXMLResponseError(t *testing.T) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)

	var actual *XMLPerson
	err := PostJSON(server.URL, &XMLPerson{}, &actual)

	assert.NotNil(t, err)
	assert.Equal(t, "", actual.Name)
}
