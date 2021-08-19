package examples

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type JsonImplentation struct {
}
type User struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Age       string `json:"age"`
}

func (JsonImplentation) Run() {

	router := mux.NewRouter()
	router.HandleFunc("/decode", func(w http.ResponseWriter, req *http.Request) {
		var user User
		json.NewDecoder(req.Body).Decode(&user)
		fmt.Fprintf(w, "%s %s is %s years old", user.Firstname, user.Lastname, user.Age)
	})
	router.HandleFunc("/encode", func(w http.ResponseWriter, req *http.Request) {
		vyas := User{
			Firstname: "vyas",
			Lastname:  "Reddy",
			Age:       "25",
		}
		json.NewEncoder(w).Encode(vyas)
	})
	http.ListenAndServe(":80", router)
}
