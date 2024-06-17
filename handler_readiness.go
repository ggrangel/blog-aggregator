package main

import "net/http"

func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	respondWithJson(w, 200, map[string]string{"status": "ok"})
}
