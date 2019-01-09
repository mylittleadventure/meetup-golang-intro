package server

import (
	"encoding/json"
	"net/http"

	"bitbucket.org/mylittleadventure/example/scraper"
)

func Serve(h []scraper.Hotel) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(h)
	})

	http.ListenAndServe(":80", nil)
}
