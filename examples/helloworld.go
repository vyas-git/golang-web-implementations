package examples

import (
	"fmt"
	"net/http"
)

func Helloworld() {
	http.HandleFunc("/", handleServer)
	http.ListenAndServe(":80", nil)
}
func handleServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world %s ", r.URL.Path)
}
