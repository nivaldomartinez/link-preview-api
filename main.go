package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"fmt"
	"io"
	"os"
	"net/http"
	"github.com/badoux/goscraper"
)

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

func getUrl(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Link-Preview API")
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", getUrlData).Methods("POST")
	router.HandleFunc("/", getUrl).Methods("GET")
	http.ListenAndServe(GetPort(), router)
}

func GetPort() string {
	var port = os.Getenv("PORT")
	if port == "" {
		port = "4747"
		fmt.Println("INFO: No PORT environment variable detected, defaulting to " + port)
	}
	return ":" + port
}
