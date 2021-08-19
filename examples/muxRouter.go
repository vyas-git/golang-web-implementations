package examples

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func CustomRouter() {
	r := mux.NewRouter()
	r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		title := vars["title"]
		page := vars["page"]
		fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)

	})
	if err := http.ListenAndServe(":80", r); err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("server listening at 80 ")

}
