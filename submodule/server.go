package submodule

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func Serve() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi")
	})

	fmt.Print("Listening on port 8081...")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
