package handler

import (
	"encoding/json"
	"net/http"

	"github.com/full-cycle-2.0-hexagonal-architecture/application"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func MakeProductHandlers(
	router *mux.Router,
	middleware *negroni.Negroni,
	service application.IProductService,
) {
	router.Handle("/product/{id}", middleware.With(
		negroni.Wrap(getProduct(service)),
	)).Methods("GET", "OPTIONS")
}

func getProduct(service application.IProductService) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", "application/json")
		vars := mux.Vars(r)
		id := vars["id"]
		product, err := service.Get(id)
		if err != nil {
			rw.WriteHeader(http.StatusNotFound)
			return
		}
		err = json.NewEncoder(rw).Encode(product)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
	})
}
