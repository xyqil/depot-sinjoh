package main

import (
  "log"
  "net/http"

  "github.com/gorilla/mux"
)

var (
  Router *mux.Router
)

func init() {
  Router = mux.NewRouter()

  // Register handlers
  Router.HandleFunc("/", EndpointHome)
}

func main() {
  log.Println("Starting gamestats2 server...")
  log.Fatal(http.ListenAndServe(":10004", Router))
}
