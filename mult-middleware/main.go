package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)
type city struct{
    Name string
    Area uint64
}


func filterContentType(handler http.Handler) http.Handler{
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Println("Currently in the check content middleware !!!!")
        if r.Header.Get("Content-type") != "application/json" {
            w.WriteHeader(http.StatusUnsupportedMediaType)
            w.Write([]byte("415 - Unsupported Media Type. Please send JSON"))
            return
        }
        handler.ServeHTTP(w,r)
    })
}

func setServerTimeCookie(handler http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        handler.ServeHTTP(w,r)
        cookie := http.Cookie{Name: "Server-Time(UTC)",
        Value: strconv.FormatInt(time.Now().Unix(), 10)}
        http.SetCookie(w, &cookie)
        log.Println("Currently in the set server middleware")
    })
}
func handle(w http.ResponseWriter, r *http.Request)  {
    if r.Method == "POST"{
        var tempCity city
        decoder := json.NewDecoder(r.Body)
        err := decoder.Decode(&tempCity)
        if err != nil{
            panic(err)
        }
        defer r.Body.Close()
        fmt.Printf("Got %s city with area %d sq miles \n", tempCity.Name, tempCity.Area)
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("201 - Created"))
    } else {
        w.WriteHeader(http.StatusMethodNotAllowed)
        w.Write([]byte("405 - Method Not Allowed"))        
    }
}

func main()  {
    originalHnadler := http.HandlerFunc(handle)
    http.Handle("/city", filterContentType(setServerTimeCookie(originalHnadler)))
    fmt.Println("Server on port 8000")
    log.Fatal(http.ListenAndServe(":8000", nil))
    
}
