package http2

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/frozzare/go-assert"
)

type JSONPerson struct {
	Name string
}

func TestGetJSON(t *testing.T) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)

	mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "{\"name\":\"Fredrik\"}")
	}))

	actual := &JSONPerson{}
	err := GetJSON(server.URL, &actual)

	assert.Nil(t, err)
	assert.Equal(t, "Fredrik", actual.Name)
}

func TestGetJSONResponseError(t *testing.T) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)

	var actual *JSONPerson
	err := GetJSON(server.URL, &actual)

	assert.NotNil(t, err)
	assert.Equal(t, "", actual.Name)
}

func TestPostJSON(t *testing.T) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)

	mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var v *JSONPerson

		err := json.NewDecoder(r.Body).Decode(&v)
		assert.Nil(t, err)
		assert.Equal(t, "Input", v.Name)

		fmt.Fprintf(w, "{\"name\":\"Fredrik\"}")
	}))

	var actual *JSONPerson
	err := PostJSON(server.URL, &JSONPerson{Name: "Input"}, &actual)

	assert.Nil(t, err)
	assert.Equal(t, "Fredrik", actual.Name)
}

func TestPostJSONResponseError(t *testing.T) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)

	var actual *JSONPerson
	err := PostJSON(server.URL, &JSONPerson{}, &actual)

	assert.NotNil(t, err)
	assert.Equal(t, "", actual.Name)
}
