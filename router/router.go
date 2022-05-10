package router

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Router() {
	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "hello world")
	})
	http.ListenAndServe(":8080", router)
}
