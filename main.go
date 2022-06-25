package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

//go:embed static
var static embed.FS

func main() {
	r := mux.NewRouter()
	r.PathPrefix("/").Handler(http.FileServer(staticFiles()))

	log.Fatal(http.ListenAndServe(":8080", r))
}

func staticFiles() http.FileSystem {
	if os.Getenv("DEVMODE") == "1" {
		return http.Dir("static")
	}

	files, _ := fs.Sub(static, "static")
	return http.FS(files)
}
