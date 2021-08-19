package examples

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type CustomMiddlewares func(http.HandlerFunc) http.HandlerFunc // Middleware always takes handler func and returns main handler func

type SingleMiddleware struct {
}
type MultiMiddleware struct {
}

func (SingleMiddleware) logging(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		f(w, r)
	}
}

func (SingleMiddleware) foo(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "foo")

}

func (SingleMiddleware) bar(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "bar")

}

func (SingleMiddleware) Run() {
	router := mux.NewRouter()
	mWare := SingleMiddleware{}

	router.HandleFunc("/foo", mWare.logging(mWare.foo))
	router.HandleFunc("/bar", mWare.logging(mWare.bar))
	http.ListenAndServe(":80", router)

}

//Multi Middle Wares

func (MultiMiddleware) logging() CustomMiddlewares {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, req *http.Request) {
			// Do middleware things
			start := time.Now()
			defer func() { log.Println(req.URL.Path, time.Since(start)) }()
			f(w, req)
		}
	}
}

func (MultiMiddleware) checkMethod(m string) CustomMiddlewares {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, req *http.Request) {
			if req.Method != m {
				http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				return
			}
			f(w, req)
		}
	}
}
func (MultiMiddleware) ClosureHandlers(f http.HandlerFunc, middleWares ...CustomMiddlewares) http.HandlerFunc {

	for _, m := range middleWares {
		f = m(f) // logging handler(checkMethod handler(foo handler))
	}
	return f
}
func (MultiMiddleware) foo(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello Foo")
}
func (MultiMiddleware) Run() {
	router := mux.NewRouter()
	mWare := MultiMiddleware{}
	router.HandleFunc("/foo", mWare.ClosureHandlers(mWare.foo, mWare.logging(), mWare.checkMethod("GEeT")))
	http.ListenAndServe(":80", router)
}
