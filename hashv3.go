package main

import (
  "crypto/sha256"
  "encoding/json"
  "encoding/hex"
  "net/http"
)

type Profile struct {
     Name    string
}

func main() {
  http.HandleFunc("/messages/", foo)
  http.HandleFunc("/messages", foo)
  http.ListenAndServe(":3000", nil)
}

func foo(rw http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
	var t Profile
        err := decoder.Decode(&t)
        js, err := json.Marshal(t.Name)
        if err != nil {
          http.Error(rw, err.Error(), http.StatusInternalServerError)
          return
        }
        hasher := sha256.New()
        hasher.Write(js)
        sha256_hash := hex.EncodeToString(hasher.Sum(nil))
        rw.Header().Set("Content-Type", "application/json")
        rw.Write([]byte(sha256_hash))
}
