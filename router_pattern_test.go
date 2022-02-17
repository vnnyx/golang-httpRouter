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

func TestPatternNamedParameter(t *testing.T) {
	router := httprouter.New()
	//Named Params: Spesifik
	router.GET("/products/:id/items/:itemId", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		//Harus sama dengat path nya
		id := params.ByName("id")
		itemId := params.ByName("itemId")
		text := "Product " + id + " item " + itemId
		_, _ = fmt.Fprint(writer, text)
	})

	request := httptest.NewRequest("GET", "http://localhost:8080/products/1/items/1", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	assert.Equal(t, "Product 1 item 1", string(body))
}

func TestPatternCatchAllParameter(t *testing.T) {
	router := httprouter.New()
	//Cath All Params: Mengambil semuanya
	router.GET("/images/*image", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		//Harus sama dengat path nya
		image := params.ByName("image")
		text := "Image: " + image
		_, _ = fmt.Fprint(writer, text)
	})

	request := httptest.NewRequest("GET", "http://localhost:8080/images/gambar/patrick.png", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	assert.Equal(t, "Image: /gambar/patrick.png", string(body))
}
