package main

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	if _, err := os.Stat("chunks"); os.IsNotExist(err) {
		convertFmpeg("nira.mp4", "chunks", 10)
	}
	http.Handle("/", handlers())
	http.ListenAndServe(":8080", nil)
}

func handlers() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", fetchIndex).Methods("GET")
	router.HandleFunc("/media/stream/", streamHandler).Methods("GET")
	router.HandleFunc("/media/stream/{segName}", streamHandler).Methods("GET")
	return router
}
