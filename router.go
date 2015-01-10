package main

import (
	"fmt"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Path(route.Pattern).
			Methods(route.Method).
			Name(route.Name).
			Handler(route.HandlerFunc)
		fmt.Printf(
			"Pattern: %-20s\tMethod: %-20s\tName: %s\n",
			route.Pattern,
			route.Method,
			route.Name,
		)
	}
	return router
}
