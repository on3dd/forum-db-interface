package controller

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"reflect"
	"runtime"
)

func (u *UserController) NewRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/messages", logHandlerCall(u.GetMessages))

	return router
}

// logHandlerCall logs any handler call
func logHandlerCall(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := runtime.FuncForPC(reflect.ValueOf(handler).Pointer()).Name()
		log.Printf("Handler function called: %v", name)
		handler(w, r)
	}
}