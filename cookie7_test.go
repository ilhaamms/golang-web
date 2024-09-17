package restful_api

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SetCookie(w http.ResponseWriter, r *http.Request) {
	cookie := new(http.Cookie)
	cookie.Name = "role"
	cookie.Value = r.URL.Query().Get("role")
	cookie.Path = "/"

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	http.SetCookie(w, cookie)
	fmt.Fprintln(w, "success create cookie")
}

func GetCookie(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("role")
	if err != nil {
		fmt.Fprintln(w, "no cookie")
	} else {
		fmt.Fprintf(w, "hallo %s", cookie.Value)
	}
}

func TestSetCookie(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080?role=admin", nil)
	recorder := httptest.NewRecorder()

	SetCookie(recorder, request)

	response := recorder.Result()
	cookies := response.Cookies()

	for _, cookie := range cookies {
		assert.Equal(t, "role", cookie.Name)
		assert.Equal(t, "admin", cookie.Value)
	}

	assert.Equal(t, "application/json", response.Header.Get("content-type"))
	assert.Equal(t, http.StatusCreated, response.StatusCode)

}

func TestGetCookieFailed(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	GetCookie(recorder, request)

	response := recorder.Result()
	body, err := io.ReadAll(response.Body)

	assert.Nil(t, err)
	assert.Equal(t, "no cookie\n", string(body))
}

func TestGetCookieSuccess(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080?role=admin", nil)
	// set dulu request cookienya
	cookie := new(http.Cookie)
	cookie.Name = "role"
	cookie.Value = request.URL.Query().Get("role")
	request.AddCookie(cookie) // jadiin cookie

	recorder := httptest.NewRecorder()

	GetCookie(recorder, request)

	response := recorder.Result()
	body, err := io.ReadAll(response.Body)

	assert.Nil(t, err)
	assert.Equal(t, "hallo admin", string(body))
}
