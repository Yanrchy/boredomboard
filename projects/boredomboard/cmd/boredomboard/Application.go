package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/Yanrchy/boredomboard/projects/boredomboard/internal/server"
)

type Application struct {
	Server http.Server
}

func (app *Application) Initialize() {

	var router *mux.Router = mux.NewRouter()
	var channel *server.CommChannel = server.NewChannel()

	go channel.Serve()

	router.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {

		server.OnClientConnection(channel, writer, request)

	})

	app.Server = http.Server{
		Addr:    ":8000",
		Handler: router,
	}

}

func (app *Application) Serve() {

	log.Printf("server: listening at port '%v'\n", app.Server.Addr)

	log.Fatal(http.ListenAndServe(app.Server.Addr, app.Server.Handler))

}
