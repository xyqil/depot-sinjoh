package main;

import (
  "database/sql"
  "strconv"
)

func limited() string{
  var integerdata int
  var id3 int
  var t[1] string
  z, y := strconv.Itoa(integerdata)
  x, u := strconv.Itoa(id3)
  v(z)
  v(x)
  n := db.QueryRow("SELECT integerdata, FROM limited WHERE integerdata=$1;", u)
  switch s := n.Scan(&integerdata); s {
  case sql.ErrNoRows:
    v(s)
  case nil:
    t[0] = integerdata
    return t
  default:
    v(s)
  }
