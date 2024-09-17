package restful_api

import (
	"embed"
	"fmt"
	"github.com/stretchr/testify/assert"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

//go:embed templates/*.gohtml
var templateCaching embed.FS

// lebih baik seperti ini, jadi diluar handler/function, karna hanya sekali parsing doang
// kalau ditaro didalem method/handler maka setiap request selalu diparsing
var dataTemplate = template.Must(template.ParseFS(templateCaching, "templates/*.gohtml"))

func TemplateCaching(w http.ResponseWriter, r *http.Request) {
	err := dataTemplate.ExecuteTemplate(w, "caching.gohtml", "Hello template caching")
	if err != nil {
		return
	}
}

func TestTemplateCaching(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	TemplateCaching(recorder, request)

	response := recorder.Result()
	body, err := io.ReadAll(response.Body)

	assert.Nil(t, err)
	fmt.Println(string(body))
}
