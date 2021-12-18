package handler

import (
	"encoding/json"
	"net/http"

	"github.com/full-cycle-2.0-hexagonal-architecture/adapters/dto"
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

	router.Handle("/product", middleware.With(
		negroni.Wrap(createProduct(service)),
	)).Methods("POST", "OPTIONS")

	router.Handle("/product/{id}/enable", middleware.With(
		negroni.Wrap(enableProduct(service)),
	)).Methods("PATCH", "OPTIONS")

	router.Handle("/product/{id}/disable", middleware.With(
		negroni.Wrap(disableProduct(service)),
	)).Methods("PATCH", "OPTIONS")
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

func createProduct(service application.IProductService) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", "application/json")
		var productDto dto.Product
		err := json.NewDecoder(r.Body).Decode(&productDto)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			rw.Write(jsonError(err.Error()))
			return
		}

		product, err := service.Create(productDto.Name, productDto.Price)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			rw.Write(jsonError(err.Error()))
			return
		}

		err = json.NewEncoder(rw).Encode(product)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			rw.Write(jsonError(err.Error()))
			return
		}
	})
}

func enableProduct(service application.IProductService) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", "application/json")
		vars := mux.Vars(r)
		id := vars["id"]
		product, err := service.Get(id)
		if err != nil {
			rw.WriteHeader(http.StatusNotFound)
			return
		}
		result, err := service.Enable(product)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			rw.Write(jsonError(err.Error()))
			return
		}
		err = json.NewEncoder(rw).Encode(result)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
	})
}

func disableProduct(service application.IProductService) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", "application/json")
		vars := mux.Vars(r)
		id := vars["id"]
		product, err := service.Get(id)
		if err != nil {
			rw.WriteHeader(http.StatusNotFound)
			return
		}
		result, err := service.Disable(product)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			rw.Write(jsonError(err.Error()))
			return
		}
		err = json.NewEncoder(rw).Encode(result)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
	})
}
