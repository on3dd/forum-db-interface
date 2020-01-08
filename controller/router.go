package controller

import (
	"github.com/gorilla/mux"
	"net/http"
)

func (u *UserController) NewRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/messages", headersMiddleware(u.GetMessages)).Methods("GET")
	router.HandleFunc("/messages", headersMiddleware(u.AddMessage)).Methods("POST", "OPTIONS")

	router.HandleFunc("/categories", headersMiddleware(u.GetCategories)).Methods("GET")

	router.HandleFunc("/forum", headersMiddleware(u.GetCategory)).Methods("GET")
	router.HandleFunc("/forum/subcategories", headersMiddleware(u.GetSubcategories)).Methods("GET")
	router.HandleFunc("/forum/messages", headersMiddleware(u.GetMessages)).Methods("GET")
	router.HandleFunc("/forum/messages", headersMiddleware(u.AddMessage)).Methods("POST", "OPTIONS")

	return router
}

// headersMiddleware adds headers to the response
func headersMiddleware(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		handler(w, r)
	}
}