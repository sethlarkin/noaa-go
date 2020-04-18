package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	resp, _ := http.Get("https://api.weather.gov/gridpoints/TOP/31,80/forecast")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
		fmt.Fprintf(w, "resp, %q", resp)
		// fmt.Println(resp, err)
	})

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi")
	})

	log.Fatal(http.ListenAndServe(":8081", nil))

}
