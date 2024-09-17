package restful_api

import (
	"embed"
	_ "embed"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/fs"
	"os"
	"testing"
)

// embed ini wajib di global variabel, jadi gabisa dalem method atau yang lainnya
// dan tipe datanya wajib string, []byte, dan embed.FS
//
//go:embed version.txt
var version string

func TestEmbed(t *testing.T) {
	assert.Equal(t, "1.0.0-SNAPSHOT", version)
}

// byte ini adalah binary seperti video, gambar, dll
//
//go:embed logo.jpg
var logo []byte

func TestEmbedByte(t *testing.T) {
	err := os.WriteFile("logo_new.jpg", logo, fs.ModePerm)
	if err != nil {
		return
	}
}

// untuk embed yang file nya didalem directory, maka langsung sebut directiry awalnya aja
//
//go:embed files/a.txt
//go:embed files/b.txt
//go:embed files/c.txt
var files embed.FS

func TestEmbedMultiple(t *testing.T) {
	byte1, err := files.ReadFile("a.txt")
	if err != nil {
		return
	}

	byte2, err := files.ReadFile("a.txt")
	if err != nil {
		return
	}

	byte3, err := files.ReadFile("a.txt")
	if err != nil {
		return
	}

	assert.Equal(t, "AAA", string(byte1))
	assert.Equal(t, "BBB", string(byte2))
	assert.Equal(t, "CCC", string(byte3))
}

// artinya ini akan ambil semua file dengan extension .txt yang ada di folder files
// embed file ini ketika di build/compile ya otomatis masuk kedalam binary golangnya
// jadi pas diubah filenya yang di server maka harus compile ulang
// karena hasilnya akan beda
//
//go:embed files/*.txt
var path embed.FS

func TestPathMatcher(t *testing.T) {
	directory, _ := path.ReadDir("files") // baca files directory yang diembed
	for _, entry := range directory {
		if !entry.IsDir() { // cek kalau bukan directory
			fmt.Println(entry.Name()) // ambil nama filenya
			byte, _ := path.ReadFile("files/" + entry.Name())
			fmt.Println(string(byte))
		}
	}
}
