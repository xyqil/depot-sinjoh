package main

import (
    "log"
)

func miiinfo() {
    var a[2] string
    sqlStatement: = `SELECT code, msg FROM miiinfo WHERE id=$1;`
    var code string
    var msg int
        // Replace 3 with an ID from your database or another random
        // value to test the no rows use case.
    row: = db.QueryRow(sqlStatement, 3)
    switch err: = row.Scan( & code, & msg);
    err {
        case sql.ErrNoRows:
            log.Fatal(err)
        case nil:
            a[0] = code
            a[1] = msg
            return a
        default:
            log.Fatal(err)
    }
}

func related() {
    var b[3] string
    sqlStatement: = `SELECT rank, movieid, title FROM related WHERE id=$1;`
    var rank string
    var movieid int
    var title int
        // Replace 3 with an ID from your database or another random
        // value to test the no rows use case.
    row: = db.QueryRow(sqlStatement, 3)
    switch err: = row.Scan( & rank, & movieid, & title);
    err {
        case sql.ErrNoRows:
            log.Fatal(err)
        case nil:
            a[0] = rank
            a[1] = movieid
            a[2] = title
            return a
        default:
            log.Fatal(err)
    }
}

func evaluate() {
    var d[2] string
    sqlStatement: = `SELECT code, msg FROM evaluate WHERE id=$1;`
    var code string
    var msg int
        // Replace 3 with an ID from your database or another random
        // value to test the no rows use case.
    row: = db.QueryRow(sqlStatement, 3)
    switch err: = row.Scan( & code, & msg);
    err {
        case sql.ErrNoRows:
            log.Fatal(err)
        case nil:
            d[0] = code
            d[1] = msg
            return d
        default:
            log.Fatal(err)
    }
}
