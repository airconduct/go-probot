package web

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:embed dist/*
var static embed.FS

func Static() fs.FS {
	f, err := fs.Sub(static, "dist")
	if err != nil {
		panic(err)
	}
	return f
}

func StaticHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rawFile, _ := static.ReadFile("dist/index.html")
		w.Write(rawFile)
	})
}

func RegisterHandler(mux *http.ServeMux) {
	mux.Handle("/dashboard/", StaticHandler())
	mux.Handle("/assets/", http.FileServer(http.FS(Static())))
}
