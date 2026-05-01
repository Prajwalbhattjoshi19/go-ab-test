package main

import (
	"goabtest/internal/abtestconf"
	"goabtest/internal/utils"
	"log"
	"net/http"
)

func main() {
	conn := utils.GetDBDriver()
	log.Println("starting server at port ", utils.GetEnv("server_port", ":8080"))
	// /conf url to route to configuration handler
	http.HandleFunc("/conf", func(w http.ResponseWriter, r *http.Request) {
		// call the configuration handler function
		abtestconf.ConfigurationHandler(w, r, conn)
	})
	http.ListenAndServe(utils.GetEnv("server_port", ":8080"), nil)

}
