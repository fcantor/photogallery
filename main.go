package main

import (
	"fmt"
	"net/http"
)

// handlerFunc is a function that's called everytime someone comes to our web app;
// runs any time a web request hits our server
func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Welcome to my awesome site!</h1>")
}

// tells our server to start up and use the handlerFunc
func main() {
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":3000", nil)
}
