package lab

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintln(w, "LAB_OK")
	})

	mux.HandleFunc("/simulate", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		input := r.URL.Query().Get("input")

		w.Header().Set("Content-Type", "text/plain")

		if input == "LAB_TEST" {
			fmt.Fprintln(w, "SIMULATED_UNSAFE_PATH: LAB_TEST")
			return
		}

		fmt.Fprintln(w, "INPUT_RECEIVED")
	})

	server := &http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: mux,
	}

	log.Println("Local security lab: http://127.0.0.1:8080")
	log.Fatal(server.ListenAndServe())
}
