package main

import (
	jsonparse "encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
)




type Args struct{
    ID string
}

type Book struct{
    ID string
    Name string
    Author string
}

type JSONServer struct{}

func (t *JSONServer) GiveBookDetail(r *http.Request, args *Args, reply *Book) error{
    var books []Book


    absPath, _ := filepath.Abs("books.json")
    raw, readerr := ioutil.ReadFile(absPath)
    if readerr != nil {
        log.Fatal("error", readerr)
        os.Exit(1)
    }
    marshalerr := jsonparse.Unmarshal(raw, &books)
    if marshalerr != nil {
        log.Fatal("error :", marshalerr)
        os.Exit(1)
    }
    for _,book:= range books {
        if book.ID == args.ID {
            *reply = book
            break
        }
        
    }
     return nil

}

func main()  {
    s := rpc.NewServer()
    s.RegisterCodec(json.NewCodec(), "application/json")

    s.RegisterService(new(JSONServer), "")
    r := mux.NewRouter()
    r.Handle("/rpc", s)
    log.Fatal(http.ListenAndServe(":1234", r))
    
}
