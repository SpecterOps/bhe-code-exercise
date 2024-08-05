// +build integration

package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	m.Run()
}

func TestPrimeHandler(t *testing.T) {

	tests := []struct {
		input    string
		expected int64
	}{
		{
			input: "0", expected: 2,
		},
		{
			input: "19", expected: 71,
		},
	}

	for _, test := range tests {
		testData := []byte(`{
			"nth_prime":` + test.input + `}`)

		svr := httptest.NewServer(http.HandlerFunc(primeNumHandler))
		defer svr.Close()

		resp, err := http.Post(svr.URL+"/prime", "application/json", bytes.NewBuffer(testData))
		if err != nil {
			log.Fatal(err)
		}

		var output map[string]int64
		decoder := json.NewDecoder(resp.Body)
		err = decoder.Decode(&output)
		if err != nil {
			log.Fatal(err)
		}

		assert.Equal(t, test.expected, output["Prime Number"])
	}
}

func TestPrimeHandlerBadInput(t *testing.T) {
	testData := []byte(`{
			"nth_prime": "101"}`)

	svr := httptest.NewServer(http.HandlerFunc(primeNumHandler))
	defer svr.Close()

	resp, err := http.Post(svr.URL+"/prime", "application/json", bytes.NewBuffer(testData))
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, 400, resp.StatusCode)
}

func TestPrimeHandlerNotFound(t *testing.T) {
	testData := []byte(`{
			"nth_prime": 101}`)

	svr := httptest.NewServer(http.HandlerFunc(primeNumHandler))
	defer svr.Close()

	resp, err := http.Post(svr.URL+"/foo", "application/json", bytes.NewBuffer(testData))
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, 404, resp.StatusCode)
}

func TestPrimeHandlerMethod(t *testing.T) {
	svr := httptest.NewServer(http.HandlerFunc(primeNumHandler))
	defer svr.Close()

	resp, err := http.Get(svr.URL+"/foo")
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, 405, resp.StatusCode)
}