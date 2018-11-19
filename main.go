package main

import "fmt"
import "net/http"

func main() {

  http.HandleFunc("/", index_handler)
  http.ListenAndServe("127.0.0.1:8080", nil)

}

func index_handler(w http.ResponseWriter, r *http.Request) {
  ans := "Cool"
  fmt.Fprintf(w, ans)
}
