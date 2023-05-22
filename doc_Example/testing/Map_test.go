package main

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestMapRoute(t *testing.T) {
	w := httptest.NewRecorder()
	payload := strings.NewReader("names%5Bfirst%5D=keval&names%5Bsecond%5D=meet&ids%5Ba%5D=1234&ids%5Bb%5D=hello")
	req, _ := http.NewRequest(http.MethodPost, "/map", payload)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	Router().ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, `IDS :  map[a:1234 b:hello]
NAMES : map[first:keval second:meet]`, w.Body.String())
}
