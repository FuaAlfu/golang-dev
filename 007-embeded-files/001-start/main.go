package main

import (
	"embed"
	"net/http"
)

func main() {
	/*
		assest/*
		means bring all files inside that folder
	*/

	//go:embed assest/*
	var assest embed.FS

	fs := http.FileServer(http.FS(assest))
	http.ListenAndServe(":9000", fs)
	print(assest)
}
