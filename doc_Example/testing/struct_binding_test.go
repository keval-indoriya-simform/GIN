package main

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

var jsonlogin Login

func TestStructBindingRoute(t *testing.T) {
	jsonlogin = Login{
		User:     "keval",
		Password: "123",
	}
	w := httptest.NewRecorder()
	jsonstr, _ := json.Marshal(jsonlogin)

	req, _ := http.NewRequest(http.MethodPost, "/struct", bytes.NewBuffer(jsonstr))
	req.Header.Add("Content-Type", "application/json")
	Router().ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, string(jsonstr), w.Body.String())
}
