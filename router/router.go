package router

import (
	"fmt"
	"net/http"

	"../routes"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

//Router exp
func Router() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	//routes
	router.HandleFunc("/", Index)
	router.HandleFunc("/chars", routes.GetAllChars).Methods("GET")
	router.HandleFunc("/chars/{id}", routes.GetChar).Methods("GET")
	router.HandleFunc("/chars", routes.CreateNewChar).Methods("POST")
	router.HandleFunc("/chars/{id}", routes.DeleteChar).Methods("DELETE")
	router.HandleFunc("/chars/{id}", routes.EditChar).Methods("PUT")

	return router
}
