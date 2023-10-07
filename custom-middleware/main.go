package main

import (
	"fmt"
	"log"
	"net/http"
)


func middleware(originalhandler http.Handler) http.Handler {
    return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request)  {
        fmt.Println("Executing middleware before the request phase!")
        //Pass control back to the handler 
        originalhandler.ServeHTTP(w, r)
        fmt.Println("Executing middleware after the request phase!")
    })
}

func handle(w http.ResponseWriter, r *http.Request)  {
    fmt.Println("Executing mainHandler!!")
    w.Write([]byte("OK"))
}

func main()  {
    originalHandler := http.HandlerFunc(handle)
    http.Handle("/", middleware(originalHandler))
    log.Fatal(http.ListenAndServe(":8000", nil))
    
}
