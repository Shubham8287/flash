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
    "flash/conf"
)

const DataPath = "./data/words_dictionary.json";

func enableCors(w *http.ResponseWriter) {
    (*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
    (*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
    (*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}


func main() {
    conf.Init();
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
        setupResponse(&w, r)
        fmt.Println("Hello World!")
        prefix := r.FormValue("prefix")
        matches := trie.Find(prefix);
        response := util.Response {
            Prefix: prefix,
            Matches: matches,
        }

        json.NewEncoder(w).Encode(response);
    })

    log.Fatal(http.ListenAndServe(":8080", nil))

}
