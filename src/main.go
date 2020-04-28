package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
)

var config Config

//create initialisation functions
func init() {
	config = CreateConfig()
	fmt.Println("Config file has loaded")
	fmt.Printf("UsermanagerHost: %v\n", config.USERMANAGERHost)
	fmt.Printf("UsermanagerPort: %v\n", config.USERMANAGERPort)
	fmt.Printf("TextHost: %v\n", config.TEXTHost)
	fmt.Printf("TextPort: %v\n", config.TEXTPort)
}

//create config functions
func CreateConfig() Config {
	conf := Config{
		USERMANAGERHost: os.Getenv("USERMANAGER_Host"),
		USERMANAGERPort: os.Getenv("USERMANAGER_Port"),
		TEXTHost:       os.Getenv("TEXT_Host"),
		TEXTPort:       os.Getenv("TEXT_PORT"),
	}
	return conf
}
func main() {

	server := Server{
		router: mux.NewRouter(),
	}
	//Set up routes for server
	server.routes()
	handler := removeTrailingSlash(server.router)
	fmt.Print("starting server on port " + config.TEXTPort + "\n")
	log.Fatal(http.ListenAndServe(":"+config.TEXTPort, handler))

}

func removeTrailingSlash(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")
		next.ServeHTTP(w, r)
	})
}
