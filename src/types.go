package main

import "github.com/gorilla/mux"

//create structs for JSON objects recieved and responses

type SendText struct {
	Number  string `json:"number"`
	Message string `json:"message"`
}

type MessageResponse struct {
	MessageID        string `json:"apiMessageId"`
	Accepted         bool   `json:"accepted"`
	To               string `json:"to"`
	ErrorCode        string `json:"errorCode"`
	Error            string `json:"error"`
	ErrorDescription string `json:"errorDescription"`
}

type MessageList struct {
	MessageResult []MessageResponse `json:"messages"`
}

type MessageResult struct {
	MessageSent bool `json:"messagesent"`
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
