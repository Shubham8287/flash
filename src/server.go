package main

import (
    "fmt"
    "html"
    "log"
    "net/http"
    "encoding/json"
    "io/ioutil"
    T "flash/trie"
    "flash/util"
)

const DataPath = "./data/words_dictionary.json";

func main() {
    trie := T.InitTrie();

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
        trie.Insert(intialWordSet[i]);
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
        matches := trie.Find(prefix);
        w.Header().Set("Content-Type", "application/json");
        response := util.Response {
            Prefix: prefix,
            Matches: matches,
        }

        json.NewEncoder(w).Encode(response);
    })

    log.Fatal(http.ListenAndServe(":8080", nil))

}
