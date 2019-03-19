package main

import (
	"log"
	"net/http"
	"time"
)

func timeHandler(format string) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		tm := time.Now().Format(format)
		writer.Write([]byte("The time is: "+tm))
	}
}


func main() {
	mux := http.NewServeMux()
    //th := &timeHandler{format: time.RFC1123}
    //mux.Handle("/time",th)
    mux.HandleFunc("/time",timeHandler(time.RFC1123))
    log.Println("Listening...")
    http.ListenAndServe(":5000",mux)
}
