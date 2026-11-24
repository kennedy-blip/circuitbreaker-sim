package main

import "net/http"

func Handler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/api/breakers":
		BreakersHandler(w, r)
	case "/api/run":
		RunHandler(w, r)
	default:
		errOut(w, http.StatusNotFound, "endpoint not found")
	}
}
