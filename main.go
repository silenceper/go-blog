package main

import (
    "net/http"
    "./xingyun"
    "log"
    "os"
    "./controller"
    "./model"
)

func init(){
    model.InitDb()
    model.Dbmap.TraceOn("[gorp]", log.New(os.Stdout, "go-blog:", log.Lmicroseconds))
}

func main(){
    cfg := &xingyun.Config{
        StaticHost:"http://192.168.4.122:9000",
        ViewPath:"view",
        Layout:"layout",
    }
    server := xingyun.NewServer(cfg)
    logger := server.Logger()

    //pipe setting
    pipe := server.NewPipe("normal", xingyun.PipeHandlerFunc(func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
        next(w, r)
    }))

    //route mapping
    server.Get("/",pipe.Wrap(controller.IndexHandler))

    err := server.ListenAndServe(":9000")
    logger.Errorf("%s", err)
}


