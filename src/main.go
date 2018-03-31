package main

import (
	"net/http"
	"os"
	"util/log"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()

	router.GET("/intro/*f", handleMarkdownIntro)
	router.GET("/markdown/*f", handleMarkdown)
	router.NotFound = http.FileServer(http.Dir("static"))

	port := getPortFromCmd()
	log.H("Server is started on port: ", port)
	log.E(http.ListenAndServe(":"+port, router))
}

func getPortFromCmd() string {
	if len(os.Args) == 2 {
		return os.Args[1]
	}

	return "80"
}
