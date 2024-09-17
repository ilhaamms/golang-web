package restful_api

import (
	"embed"
	"io/fs"
	"net/http"
	"testing"
)

func TestFileServer(t *testing.T) {
	directory := http.Dir("./resources")     // ambil directory yang akan digunakan untuk file server
	fileServer := http.FileServer(directory) // jadikan file server

	mux := http.NewServeMux()
	// karena file server itu baca url nya, lalu mencari file berdasarkan url nya
	//jadi kalaukita buat /static/index.js maka file server akan mencari ke /resources/static/index.js
	//jadinya kita bisa hapus prefix di url pake strip prefix, maka menjadi /resources/index.js
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		return
	}
}

//go:embed resources
var resources embed.FS

func TestFileServerGolangEmbed(t *testing.T) {
	// ini artinya kita masuk ke folder directory resources, jadi sama aja kayak /static/resources/index.html
	directory, _ := fs.Sub(resources, "resources")
	// jadi di server otomatis tinggal langsung /static/index.html aja, karna udah di setting di line 36
	fileServer := http.FileServer(http.FS(directory)) // http.FS ini konversi dari golang embed menjadi file system

	mux := http.NewServeMux()
	// karena file server itu baca url nya, lalu mencari file berdasarkan url nya
	//jadi kalaukita buat /static/index.js maka file server akan mencari ke /resources/static/index.js
	//jadinya kita bisa hapus prefix di url pake strip prefix, maka menjadi /resources/index.js
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		return
	}
}
