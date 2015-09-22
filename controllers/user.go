package controllers

import (
    "log"
    "net/http"
    "encoding/json"

    "github.com/javinc/playgo/goryo/services/user"
    res "github.com/javinc/playgo/goryo/resources/user"
)

func UserHandler(w http.ResponseWriter, r *http.Request) {
    log.Println("controllers UserHandler")

    // exsract field from get params
    o := new(res.Options)
    err := decoder.Decode(o, r.URL.Query())
    if err != nil {
        log.Panic(err)
    }

    log.Println("options", o)

    // exsract field from post data
    err = r.ParseForm()
    if err != nil {
        log.Panic(err)
    }

    p := new(res.Model)
    err = decoder.Decode(p, r.PostForm)
    if err != nil {
        log.Panic(err)
    }

    log.Println("payload", p)

    // lets base on the request type
    // use service
    var b []byte
    switch r.Method {
    case "GET":
        b, err = json.Marshal(user.Find(o))

    case "POST":
        b, err = json.Marshal(user.Create(p))
    }

    // render some value
    if err != nil {
        log.Panic(err)
    }

    w.Write(b)
}
