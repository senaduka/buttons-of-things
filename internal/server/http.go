package server

import "net/http"

func Start(pattern string, handler Handler) {
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    stringHandler(w,r,handler)
  })
  if err := http.ListenAndServe(":8080", nil); err != nil {
    panic(err)
  }
}


