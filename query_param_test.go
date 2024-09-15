package restful_api

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func SayHello(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name") // ambil query param
	if name == "" {
		fmt.Fprint(w, "Hello")
	} else {
		fmt.Fprint(w, "Hello "+name)
	}
}

func SayHelloMultiple(w http.ResponseWriter, r *http.Request) {
	firstName := r.URL.Query().Get("firstname") // ambil query param
	middleName := r.URL.Query().Get("middlename")
	lastName := r.URL.Query().Get("lastname")

	fmt.Fprintf(w, "Selamat datang %s %s %s", firstName, middleName, lastName)
}

func TestQueryParam(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080?name=Ilham", nil)
	response := httptest.NewRecorder()

	SayHello(response, request)

	result := response.Result()
	body, _ := io.ReadAll(result.Body)

	assert.Equal(t, 200, response.Code)
	assert.Equal(t, "Hello Ilham", string(body))
}

func TestQueryParamMultiple(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080?firstname=Ilham&middlename=Muhammad&lastname=Sidiq", nil)
	response := httptest.NewRecorder()

	SayHelloMultiple(response, request)

	result := response.Result()
	body, _ := io.ReadAll(result.Body)

	assert.Equal(t, 200, response.Code)
	assert.Equal(t, "Selamat datang Ilham Muhammad Sidiq", string(body))
}

func multipleParameterQuery(writer http.ResponseWriter, request *http.Request) {
	name := request.URL.Query()["name"]
	result := strings.Join(name, " ")
	fmt.Fprintf(writer, "Selamat datang %s", result)
}

func TestMultipleParameterQuery(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/values?name=ilham&name=muhammad&name=sidiq", nil)
	response := httptest.NewRecorder()

	multipleParameterQuery(response, request)

	result := response.Result()
	body, _ := io.ReadAll(result.Body)
	name := string(body)

	assert.Equal(t, 200, response.Code)
	assert.Equal(t, "Selamat datang ilham muhammad sidiq", name)
}
