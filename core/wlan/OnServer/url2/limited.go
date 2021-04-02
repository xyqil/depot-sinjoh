package main;

import (
  "database/sql"
  "strconv"
)

func limited() {
  sqlStatement := `SELECT integerdata, FROM limited WHERE integerdata=$1;`
  var integerdata int
  var id3 int
  var t[1] string
  z, y := strconv.Itoa(integerdata)
  x, u := strconv.Itoa(id3)
  v(z)
  v(x)
  row := db.QueryRow(sqlStatement, u)
  switch s := row.Scan(&integerdata); s {
  case sql.ErrNoRows:
    v(s)
  case nil:
    t[0] = integerdata
    return t
  default:
    v(s)
  }
