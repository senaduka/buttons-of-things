package server

import "net/http"

type Handler func() string

func stringHandler(w http.ResponseWriter, r *http.Request, handler Handler) {
  ui :=  handler()
  w.Write([]byte(ui))
}


