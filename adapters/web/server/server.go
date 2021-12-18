package server

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/full-cycle-2.0-hexagonal-architecture/adapters/web/handler"
	"github.com/full-cycle-2.0-hexagonal-architecture/application"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

type WebServer struct {
	Service application.IProductService
}

func MakeNewWebServer() *WebServer {
	return &WebServer{}
}

func (w WebServer) Serve() {
	router := mux.NewRouter()
	middleware := negroni.New(
		negroni.NewLogger(),
	)

	handler.MakeProductHandlers(router, middleware, w.Service)
	http.Handle("/", router)

	server := http.Server{
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		Addr:              ":8080",
		Handler:           http.DefaultServeMux,
		ErrorLog:          log.New(os.Stderr, "log: ", log.Lshortfile),
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
