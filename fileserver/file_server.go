package main

import (
	//"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/goji/httpauth"
	"github.com/gorilla/mux"
)

const (
	PORT       = "8088"
)

//  http://http:8088/message/retort.wav - example file location to mimic Asterisk

func main() {
	router := mux.NewRouter().StrictSlash(true)

	// middleware for improved logs and basic username/password authentication
	router.Use(middleware.Logger)
	router.Use(httpauth.SimpleBasicAuth("a_username", "a_password"))

	fs := http.FileServer(http.Dir("/audio/"))
	router.PathPrefix("/message/").Handler(http.StripPrefix("/message/", fs))

	log.Printf("Listening on port %s\n", PORT)
	err := http.ListenAndServe(":"+PORT, router)

	if err != nil{
		log.Fatal("ListenAndServe Error: ", err)
	}

}





// func main() {

// 	router := mux.NewRouter().StrictSlash(true)
// 	router.Use(middleware.Logger)

// 	router.HandleFunc("/mbmessage/audio/", messageHandler)
// 	router.PathPrefix("/mbmessage/audio/").Handler(http.StripPrefix("/mbmessage/audio/", http.FileServer(http.Dir("./fileserver/audio"))))
// 	//http.Handle("/mbmessage/audio/", httpauth.SimpleBasicAuth("asterisk_username", "asterisk_password")(router))

// 	log.Fatal(http.ListenAndServe(":"+PORT, nil))
// }

// func messageHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "audio/wav")
// 	w.Write([]byte("file served"))
// }
