package main

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestURLRoute(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/keval/987fbc97-4bed-5078-9f07-9141ba07c9f3", nil)
	Router().ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, `{"name":"keval","uuid":"987fbc97-4bed-5078-9f07-9141ba07c9f3"}`, w.Body.String())
}
