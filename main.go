// Based on this tutorial: https://thenewstack.io/make-a-restful-json-api-go/

package main

import (
    "fmt"
    "html"
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
    })

    // Add a log statement so you know the web server is running
    log.Println("listening on port :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))

}
