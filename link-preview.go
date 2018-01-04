package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"github.com/badoux/goscraper"
)

type PostData struct {
    Url string
}

func getUrlData(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	url, ok := r.Form["url"]
	if !ok {
		http.Error(w, "The url is required", http.StatusBadRequest)
		return
	}
	s, err := goscraper.Scrape(url[0], 5)
	if err != nil {
		http.Error(w, "can't generate preview", http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(s.Preview)

}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", getUrlData).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", router))
}
