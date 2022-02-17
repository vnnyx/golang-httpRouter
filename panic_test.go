package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPanicHandler(t *testing.T) {
	router := httprouter.New()
	//Panic Handler
	router.PanicHandler = func(writer http.ResponseWriter, request *http.Request, error interface{}) {
		_, _ = fmt.Fprint(writer, "Panic: ", error)
	}
	router.GET("/", func(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
		panic("Ups")
	})

	request := httptest.NewRequest("GET", "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	assert.Equal(t, "Panic: Ups", string(body))
}
