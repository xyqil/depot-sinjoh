package main

import (
  "net/http"
)

func EndpointHome(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("200\n"))
}
