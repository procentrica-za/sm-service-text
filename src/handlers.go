package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func (s *Server) handlesendtext() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handle forgot password Has Been Called in the email service")
		textmessage := SendText{}

		// convert received JSON payload into the declared struct with response from user manager
		err := json.NewDecoder(r.Body).Decode(&textmessage)

		//check for errors when converting JSON payload into struct.
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Bad JSON provided to send text message")
			return
		}

		l := strings.NewReplacer(" ", "+")
		Message := l.Replace(textmessage.Message)

		response, err := http.Get("https://platform.clickatell.com/messages/http/send?apiKey=xc8lCrpgTq-tPUYv3e2sPA==&to=" + textmessage.Number + "&content=" + Message)

		//convert struct back to JSON
		js, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}

		//return back to Front-End user
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(js)
	}
}
