package restful_api

import (
	_ "embed"
	"fmt"
	"net/http"
	"testing"
)

func ServeFileGolang(w http.ResponseWriter, r *http.Request) {
	if name := r.URL.Query().Get("name"); name != "" {
		//digunakan untuk melayani file statis melalui HTTP
		http.ServeFile(w, r, "./resources/ok.html")
	} else {
		http.ServeFile(w, r, "./resources/not-found.html")
	}
}

func TestServeFileGolang(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(ServeFileGolang),
	}

	err := server.ListenAndServe()
	if err != nil {
		return
	}
}

//go:embed resources/ok.html
var resourcesOk string

//go:embed resources/not-found.html
var resourcesNotFound string

func ServeFileEmbed(w http.ResponseWriter, r *http.Request) {
	if name := r.URL.Query().Get("name"); name != "" {
		//digunakan untuk melayani file statis melalui HTTP
		fmt.Fprint(w, resourcesOk)
	} else {
		fmt.Fprint(w, resourcesNotFound)
	}
}

func TestServeFileGolangEmbed(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(ServeFileEmbed),
	}

	err := server.ListenAndServe()
	if err != nil {
		return
	}
}
