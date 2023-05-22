package main

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

func TestUploadRoute(t *testing.T) {
	w := httptest.NewRecorder()
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("name", "keval")
	_ = writer.WriteField("email", "kevalindoriya5@gmail.com")
	file, _ := os.Open("/home/keval/Desktop/practice/GIN/doc_Example/logs_in_file/gin.log")
	defer file.Close()
	part3, _ := writer.CreateFormFile("files", filepath.Base("/home/keval/Desktop/practice/GIN/doc_Example/logs_in_file/gin.log"))
	io.Copy(part3, file)
	writer.Close()
	req, _ := http.NewRequest(http.MethodPost, "/upload", payload)

	req.Header.Set("Content-Type", writer.FormDataContentType())
	Router().ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, `Uploaded successfully 1 files with fields name=keval and email=kevalindoriya5@gmail.com.`, w.Body.String())
}
