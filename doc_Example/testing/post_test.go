package main

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestPostRoute(t *testing.T) {
	w := httptest.NewRecorder()
	payload := strings.NewReader("name=keval&message=hiii")
	req, _ := http.NewRequest(http.MethodPost, "/post?id=3&page=3", payload)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	Router().ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, `{"id":"3","message":"hiii","name":"keval","page":"3"}`, w.Body.String())
}
