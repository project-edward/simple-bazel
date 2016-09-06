package main

import "net/http"
import "html"
import "fmt"

func main() {
	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})
    panic(http.ListenAndServe(":8080", nil))
}
