package server

import (
	"RuLang2GoLang/translator"
	"fmt"
	"log"
	"net/http"
	"time"
)

var Port = ":5555"

func Init() {

	http.HandleFunc("/", ServeFiles)
	fmt.Println("🚀 Serving at ", "http://127.0.0.1"+Port)
	log.Fatal(http.ListenAndServe(Port, nil))
}

func ServeFiles(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case "GET":

		path := r.URL.Path

		//fmt.Println(path)

		if path == "/" {

			path = "./web/static/index.html"
		} else {

			path = "." + path
		}

		http.ServeFile(w, r, path)

	case "POST":

		r.ParseMultipartForm(0)

		message := r.FormValue("message")

		ruOut := translator.Init(message)

		//fmt.Println("----------------------------------")
		//fmt.Println("Message from Client: ", message)
		// respond to client's request
		fmt.Fprintf(w, time.Now().Format(time.RFC3339)+" Ru Output: ")
		fmt.Fprintf(w, "\n%s\n", ruOut)

	default:
		fmt.Fprintf(w, "Request type other than GET or POSt not supported")

	}

}
