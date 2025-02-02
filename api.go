package main

import "github.com/gorilla/mux"

func server() *mux.Router {
	router := mux.NewRouter()

	// Endpoints

	//JSON Client

	// ISO Client
	router.HandleFunc("/payment/channel/iso", sendIso).Methods("POST")
	router.HandleFunc("/payment/channel/file", sendFile).Methods("POST")
	router.HandleFunc("/payment/channel/upload", uploadFile).Methods("POST")

	return router
}
