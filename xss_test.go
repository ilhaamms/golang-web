package restful_api

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

//untuk menghindari xss(cross site scripting), maka pas diimport pakenya html/template, bukan yang text/template
//xss biasanya untuk mencuri cookie user yg sedang mengakses website kita, jadi akun orang lain bisa diambil alih

func TemplateAutoEscape(w http.ResponseWriter, r *http.Request) {
	err := dataTemplate.ExecuteTemplate(w, "post.gohtml", map[string]any{
		"Title": "Selamat Datang",
		"Body":  `<script>Alert("Anda di hack")</script>`,
	})
	if err != nil {
		return
	}
}

func TestTemplateAutoEscape(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	TemplateAutoEscape(recorder, request)

	response := recorder.Result()
	body, err := io.ReadAll(response.Body)

	assert.Nil(t, err)
	fmt.Println(string(body))
}

func TestTemplateAutoEscapeServer(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(TemplateAutoEscape),
	}

	err := server.ListenAndServe()
	if err != nil {
		return
	}
}

func TemplateDisabledAutoEscape(w http.ResponseWriter, r *http.Request) {
	err := dataTemplate.ExecuteTemplate(w, "post.gohtml", map[string]any{
		"Title": "Selamat Datang",
		// kalau mau matikan auto escape pake cara ini
		// tapi hati2 jangan dari inputan user, karna bisa kena xss
		"Body": template.HTML(r.URL.Query().Get("body")),
	})
	if err != nil {
		return
	}
}

func TestTemplateDisabledAutoEscape(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	TemplateDisabledAutoEscape(recorder, request)

	response := recorder.Result()
	body, err := io.ReadAll(response.Body)

	assert.Nil(t, err)
	fmt.Println(string(body))
}

func TestTemplateDisabledAutoEscapeServer(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(TemplateDisabledAutoEscape),
	}

	err := server.ListenAndServe()
	if err != nil {
		return
	}
}
