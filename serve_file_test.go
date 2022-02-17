package main

import (
	"embed"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"io"
	"io/fs"
	"net/http"
	"net/http/httptest"
	"testing"
)

//go:embed resources
var resources embed.FS

func TestServeFile(t *testing.T) {
	router := httprouter.New()
	//Untuk serve file harus menggunakan catch all dan nama paramnya "filepath"
	directory, _ := fs.Sub(resources, "resources")
	router.ServeFiles("/files/*filepath", http.FS(directory))

	request := httptest.NewRequest("GET", "http://localhost:8080/files/hello.txt", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	assert.Equal(t, "Hello HttpRouter", string(body))
}

func TestServeFileGoodBye(t *testing.T) {
	router := httprouter.New()
	//Untuk serve file harus menggunakan catch all dan nama paramnya "filepath"
	directory, _ := fs.Sub(resources, "resources")
	router.ServeFiles("/files/*filepath", http.FS(directory))

	request := httptest.NewRequest("GET", "http://localhost:8080/files/goodbye.txt", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	assert.Equal(t, "GoodBye HttpRouter", string(body))
}
