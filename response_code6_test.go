package restful_api

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func ResponseCode(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	if name == "" {
		w.Header().Add("content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "silahkan masukan query parameter")
	} else {
		w.Header().Add("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Selamat datang %s", name)
	}
}

func TestResponseCodeInvalid(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	ResponseCode(recorder, request)

	response := recorder.Result()
	body, err := io.ReadAll(response.Body)

	assert.Nil(t, err)
	assert.Equal(t, "application/json", response.Header.Get("content-type"))
	assert.Equal(t, http.StatusBadRequest, response.StatusCode)
	assert.Equal(t, "silahkan masukan query parameter", string(body))
}

func TestResponseCodeSuccess(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080?name=Ilham", nil)
	recorder := httptest.NewRecorder()

	ResponseCode(recorder, request)

	response := recorder.Result()
	body, err := io.ReadAll(response.Body)

	assert.Nil(t, err)
	assert.Equal(t, "application/json", response.Header.Get("content-type"))
	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, "Selamat datang Ilham", string(body))
}
