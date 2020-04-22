package main

import "github.com/gorilla/mux"

//create structs for JSON objects recieved and responses

type SendText struct {
	Number  string `json:"number"`
	Message  string `json:"message"`
}

type EmailResult struct {
	Message string `json:"message"`
}

//touter service struct
type Server struct {
	router *mux.Router
}

//config struct
type Config struct {
	USERMANAGERHost string
	USERMANAGERPort string
	TEXTHost        string
	TEXTPort        string
}
