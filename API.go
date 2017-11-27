package main

import (
    //"fmt"
    //"html"
    "log"
    "net/http"
    "encoding/json"
    "io/ioutil"
)

type Todo struct {
    Name      string
    Completed bool
}

type Todos []Todo

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
      b,_ := ioutil.ReadFile("output.json");
      rawIn := json.RawMessage(string(b))
      var objmap map[string]*json.RawMessage
      err := json.Unmarshal(rawIn, &objmap)
      if err != nil {
      }

      json.NewEncoder(w).Encode(objmap)
//  json.NewEncoder(w).Encode(b)
    })
    log.Fatal(http.ListenAndServe(":80", nil))

}
