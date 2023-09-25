package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func fetchIndex(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}
func streamHandler(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	segName, ok := vars["segName"]
	if !ok {
		mediaBase := getMediaBase()
		m3u8Name := "playlist.m3u8"
		serveHlsM3u8(response, request, mediaBase, m3u8Name)
	} else {
		mediaBase := getMediaBase()
		serveHlsTs(response, request, mediaBase, segName)
	}
}
func getMediaBase() string {
	mediaRoot := "chunks"
	return mediaRoot
}

func serveHlsM3u8(w http.ResponseWriter, r *http.Request, mediaBase, m3u8Name string) {
	mediaFile := fmt.Sprintf("%s/%s", mediaBase, m3u8Name)
	http.ServeFile(w, r, mediaFile)
	w.Header().Set("Content-Type", "application/x-mpegURL")
}

func serveHlsTs(w http.ResponseWriter, r *http.Request, mediaBase, segName string) {
	mediaFile := fmt.Sprintf("%s/%s", mediaBase, segName)
	http.ServeFile(w, r, mediaFile)
	w.Header().Set("Content-Type", "video/MP2T")
}
