package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func (s *Server) handlesendtext() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Starting the application...")
		response, err := http.Get("https://platform.clickatell.com/messages/http/send?apiKey=xc8lCrpgTq-tPUYv3e2sPA==&to=27824413098&content=Your+Requested+OTP+is:+1234")
		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		} else {
			data, _ := ioutil.ReadAll(response.Body)
			fmt.Println(string(data))
		}
		fmt.Println("Terminating the application...")
	}
}
