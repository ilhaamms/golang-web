package restful_api

import (
	"net/http"
	"testing"
)

func UploadFile(w http.ResponseWriter, r *http.Request) {
	err := dataTemplate.ExecuteTemplate(w, "upload.form.gohtml", nil)
	if err != nil {
		return
	}
}

func TestUploadFile(t *testing.T) {
	mux := http.NewServeMux()

	mux.HandleFunc("/form", UploadFile)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		return
	}
}
