package examples

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

var (
	key = []byte("super-secret-key")
	//store = sessions.NewFilesystemStore("./", key)
	store = sessions.NewCookieStore(key)
)

type AuthSession struct {
}

func secret(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")

	// Check if user is authenticated
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	// Print secret message
	fmt.Fprintln(w, "The cake is a lie!")
}

func login(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")

	// Authentication goes here
	// ...

	// Set user as authenticated
	session.Values["authenticated"] = true
	fmt.Println(session.Options.Path)
	session.Save(r, w)
}

func logout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")

	// Revoke users authentication
	session.Values["authenticated"] = false
	session.Save(r, w)
}
func (AuthSession) Run() {
	router := mux.NewRouter()
	router.HandleFunc("/login", login)
	router.HandleFunc("/secret", secret)
	router.HandleFunc("/logout", logout)

	http.ListenAndServe(":80", router)
}
