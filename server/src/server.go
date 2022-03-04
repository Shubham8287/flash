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

func enableCors(w *http.ResponseWriter) {
    (*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
    (*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
    (*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}


func getTrieObject() {
    if (trie) {
        return trie;
    }
    trie := T.InitTrie();
    return trie;
}

func initStuff() {
    localTrieObj := getTrieObject();
    for i := 0; i < len(intialWordSet); i++ {
        localTrieObj.Insert(intialWordSet[i]);
    }

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
}

func homeReq(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path));
}

func pinReq(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Boss is always fine");
}

func findReq(w http.ResponseWriter, r *http.Request) {
    localTrieObj:= getTrieObject();
    setupResponse(&w, r)
        fmt.Println("Hello World!")
        prefix := r.FormValue("prefix")
        matches := localTrieObj.Find(prefix);
        w.Header().Set("Content-Type", "application/json");
        response := util.Response {
            Prefix: prefix,
            Matches: matches,
        }
        json.NewEncoder(w).Encode(response);
}


func handleRequests() {
    fmt.Println("Server Started...");
    http.HandleFunc("/", homeReq);
    http.HandleFunc("/isAlive", pingReq);
    http.HandleFunc("/find", findReq);
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
    // Intialize stuff
    initStuff();
    // Spin up the server and ready to serve requests.
    handleRequests();
}
