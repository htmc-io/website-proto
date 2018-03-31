package main

import (
	"fmt"
	"markdown"
	"net/http"
	"strings"
	"unsafe"
	"util/json"
	"util/log"

	"github.com/julienschmidt/httprouter"
)

func handleMarkdownIntro(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	path := strings.Replace(r.RequestURI, "/intro/", "/content/", 1)

	url := "./static" + path
	intro, e := markdown.ParseIntro(url)
	if e != nil {
		log.E(e)
		writeObj(w, nil)
		return
	}

	writeObj(w, intro)
}

func handleMarkdown(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	path := strings.Replace(r.RequestURI, "/markdown/", "/content/", 1)

	url := "./static" + path
	html, e := markdown.ParseContent(url)
	if e != nil {
		log.E(e)
		writeObj(w, nil)
		return
	}

	writeHTML(w, html)
}

func writeObj(w http.ResponseWriter, obj interface{}) {
	bytes := json.ToJSON(obj)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type, token")
	w.Header().Set("content-type", "application/json")
	fmt.Fprintf(w, *(*string)(unsafe.Pointer(&bytes)))
}

func writeHTML(w http.ResponseWriter, html []byte) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type, token")
	w.Header().Set("content-type", "application/html")
	fmt.Fprintf(w, *(*string)(unsafe.Pointer(&html)))
}
