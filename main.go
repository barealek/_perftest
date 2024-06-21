package main

import (
	"net/http"

	"github.com/bytedance/sonic"
	"github.com/charmbracelet/log"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	type response struct {
		Message string `json:"message"`
		Emoji   string `json:"emoji"`
	}

	mux.HandleFunc("/json", func(w http.ResponseWriter, r *http.Request) {
		var resp response
		resp.Message = "Hello, World!"
		resp.Emoji = "🌍"
		w.Header().Set("Content-Type", "application/json")

		sonic.ConfigFastest.NewEncoder(w).Encode(resp)
	})

	log.Info("Running")
	http.ListenAndServe(":3000", mux)
}
