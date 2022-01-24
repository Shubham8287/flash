package main

import (
    "fmt"
    "html"
    "log"
    "net/http"
    "encoding/json"
    "io/ioutil"
)

const DataPath = "./data/words_dictionary.json";

func main() {
    trie := initTrie();

    fmt.Println("reading data...")  
    fileData, err := ioutil.ReadFile(DataPath);
    if err != nil {
        fmt.Print(err)
    }
    
    var intialWordSet = []string{}
    fmt.Println("Unmarshalling data...")
    err = json.Unmarshal(fileData, &intialWordSet)
    if err != nil {
        fmt.Println("error:", err)
    }

    for i := 0; i < len(intialWordSet); i++ {
        trie.insert(intialWordSet[i]);
    }

    fmt.Println("Server Started...")

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
    })

    http.HandleFunc("/isAlive", func(w http.ResponseWriter, r *http.Request){
        fmt.Fprintf(w, "Boss is always fine")
    })

    http.HandleFunc("/find", func(w http.ResponseWriter, r *http.Request) {
        prefix := r.FormValue("prefix")
        if(trie.find(prefix)) {
            fmt.Fprintf(w, "Found the prefix")
        } else {
            fmt.Fprintf(w, "Not found any matching")
        }
    })

    log.Fatal(http.ListenAndServe(":8080", nil))

}
