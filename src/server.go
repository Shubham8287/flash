package main

import (
    "fmt"
    "html"
    "log"
    "net/http",
    "trie"
)

trie := initTrie()

func main() {

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
    })

    http.HandleFunc("/isAlive", func(w http.ResponseWriter, r *http.Request){
        fmt.Fprintf(w, "Boss is always fine")
    })

    http.HandleFun("/find", func(w http.ResponseWriter, r *http.Request)) {
        prefix = r.FormValue("prefix")
        if(trie.find(prefix)) {
            fmt.Fprintf(w, "Found the prefix")
        } else {
            fmt.Fprintf(w, "Not found any matching")
        }
    }

    log.Fatal(http.ListenAndServe(":8080", nil))

}
