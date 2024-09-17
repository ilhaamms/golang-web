package restful_api

import (
	"embed"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

func SimpleHTMLFile(w http.ResponseWriter, r *http.Request) {
	// ini sama aja kayak template.New("SIMPLE").Parse(templateText)
	//bedanya kalau yang template.Must udah dihandle sendiri sama golang errornya
	t := template.Must(template.ParseFiles("./templates/simple.gohtml"))
	err := t.ExecuteTemplate(w, "simple.gohtml", "Hallo HTML template") // nama templatenya emang sama kayak nama filenya
	if err != nil {
		return
	}
}

func TestSimpleHTMLFile(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	SimpleHTMLFile(recorder, request)

	response := recorder.Result()
	body, err := io.ReadAll(response.Body)

	assert.Nil(t, err)
	fmt.Println(string(body))
}

func TemplateDirectory(w http.ResponseWriter, r *http.Request) {
	//pake template.ParseGlob untuk supaya manggil semua template dengan akhir .gohtml
	t := template.Must(template.ParseGlob("./templates/*.gohtml"))
	err := t.ExecuteTemplate(w, "simple.gohtml", "Hallo Template Directory")
	if err != nil {
		return
	}
}

func TestTemplateDirectory(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	TemplateDirectory(recorder, request)

	response := recorder.Result()
	body, err := io.ReadAll(response.Body)

	assert.Nil(t, err)
	fmt.Println(string(body))
}

//go:embed templates/*.gohtml
var templates embed.FS

func TemplateEmbed(w http.ResponseWriter, r *http.Request) {
	//pake template.ParseFS untuk File System
	//karna pake golang embed jadinya gausah pake ./templates, jadi langsung nama directorynya aja
	t := template.Must(template.ParseFS(templates, "templates/*.gohtml"))
	err := t.ExecuteTemplate(w, "simple.gohtml", "Hallo Template Embed")
	if err != nil {
		return
	}
}

func TestTemplateEmbed(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	TemplateEmbed(recorder, request)

	response := recorder.Result()
	body, err := io.ReadAll(response.Body)

	assert.Nil(t, err)
	fmt.Println(string(body))
}
