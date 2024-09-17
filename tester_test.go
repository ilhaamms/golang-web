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

func httpTester(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Selamat datang")
}

func TestHttpTester(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	httpTester(recorder, request)

	response := recorder.Result()
	body, err := io.ReadAll(response.Body)

	assert.Nil(t, err)
	assert.Equal(t, "Selamat datang", string(body))
}

func SayHelloQueryParam(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	address := r.URL.Query().Get("address")
	if name == "" {
		fmt.Fprint(w, "harap masukan nama")
	} else {
		fmt.Fprintf(w, "selamat datang %s di %s", name, address)
	}
}

func TestSayHelloQueryParamSuccess(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080?name=Ilham&address=Bekasi", nil)
	recorder := httptest.NewRecorder()

	SayHelloQueryParam(recorder, request)

	response := recorder.Result()
	body, err := io.ReadAll(response.Body)

	assert.Nil(t, err)
	assert.Equal(t, "selamat datang Ilham di Bekasi", string(body))
}

func TestSayHelloQueryParamFailed(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	SayHelloQueryParam(recorder, request)

	response := recorder.Result()
	body, err := io.ReadAll(response.Body)

	assert.Nil(t, err)
	assert.Equal(t, "harap masukan nama", string(body))
}

func RequestHeaderGaes(w http.ResponseWriter, r *http.Request) {
	r.Header.Add("content-type", "application/json")
	fmt.Fprint(w, "Hallo gaes")
}

func TestRequestHeaderGaes(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	RequestHeaderGaes(recorder, request)

	response := recorder.Result()
	body, err := io.ReadAll(response.Body)

	assert.Nil(t, err)
	assert.Equal(t, "application/json", request.Header.Get("content-type"))
	assert.Equal(t, "Hallo gaes", string(body))
}

func ResponseHeaderGaes(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	fmt.Fprint(w, "selamat datang gaes response")
}

func TestResponseHeaderGaes(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)

	//Recorder digunakan untuk merekam status code, header,
	//dan body dari respons yang dihasilkan oleh handler selama pengujian.
	//seperti "rekaman" virtual yang menyimpan apa yang ditulis oleh handler ke respons,
	//sehingga Anda bisa memeriksanya di dalam unit test.
	recorder := httptest.NewRecorder()

	ResponseHeaderGaes(recorder, request)

	response := recorder.Result()
	body, err := io.ReadAll(response.Body)

	assert.Nil(t, err)
	assert.Equal(t, "application/json", response.Header.Get("content-type"))
	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, "selamat datang gaes response", string(body))
}

func FormPostGaes(w http.ResponseWriter, r *http.Request) {
	//r.Header.Add("content-type", "application/x-www-form-urlencoded")

	name := r.PostFormValue("name")
	address := r.PostFormValue("address")

	w.WriteHeader(201)
	fmt.Fprintf(w, "selamat datang bro %s di %s", name, address)
}

func TestFormPostGaes(t *testing.T) {
	requestBody := strings.NewReader("name=Ilham&address=Bekasi")
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/", requestBody)
	//request.Header.Add("content-type", "application/x-www-form-urlencoded")
	//w.WriteHeader(201)

	recorder := httptest.NewRecorder()

	FormPostGaes(recorder, request)

	response := recorder.Result()
	body, err := io.ReadAll(response.Body)

	assert.Nil(t, err)
	assert.Equal(t, "application/x-www-form-urlencoded", request.Header.Get("content-type"))
	assert.Equal(t, 201, response.StatusCode)
	assert.Equal(t, "selamat datang bro Ilham di Bekasi", string(body))
}

//func TestFormPostGaes(t *testing.T) {
//	requestBody := strings.NewReader("name=Ilham&address=Bekasi")
//	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/", requestBody) // Ganti MethodGet dengan MethodPost
//	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")                // Set header content type
//	recorder := httptest.NewRecorder()
//
//	FormPostGaes(recorder, request)
//
//	response := recorder.Result()
//	body, err := io.ReadAll(response.Body)
//
//	assert.Nil(t, err)
//	//assert.Equal(t, "application/x-www-form-urlencoded", request.Header.Get("content-type"))
//	//assert.Equal(t, 201, response.StatusCode)
//	assert.Equal(t, "selamat datang bro Ilham di Bekasi", string(body))
//}
