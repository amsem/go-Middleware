package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

//hello 
func handle(w http.ResponseWriter, r *http.Request)  {
    log.Println("Processing request")
    w.Write([]byte("OK"))
    log.Println("Finishjed processeing request")
}


func main()  {
    r := mux.NewRouter()
    r.HandleFunc("/", handle)
    loggerRouter := handlers.LoggingHandler(os.Stdout, r)
    log.Fatal(http.ListenAndServe(":8000", loggerRouter))
}
