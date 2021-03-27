package main

import (
    "log"
)

func miiinfo() {
    var a[2] string
    e: = `SELECT code, msg FROM miiinfo WHERE id=$1;`
    var code string
    var msg int
    j: = db.QueryRow(e, id1)
    switch k: = j.Scan( & code, & msg);
    k {
        case sql.ErrNoRows:
            log.Fatal(k)
        case nil:
            a[0] = code
            a[1] = msg
            return a
        default:
            log.Fatal(k)
    }
}

func related() {
    var b[3] string
    f: = `SELECT rank, movieid, title FROM related WHERE id=$1;`
    var rank string
    var movieid int
    var title int
    i: = db.QueryRow(f, id2)
    switch l: = i.Scan( & rank, & movieid, & title);
    l {
        case sql.ErrNoRows:
            log.Fatal(l)
        case nil:
            a[0] = rank
            a[1] = movieid
            a[2] = title
            return b
        default:
            log.Fatal(l)
    }
}

func evaluate() {
    var d[2] string
    g: = `SELECT code, msg FROM evaluate WHERE id=$1;`
    var code string
    var msg int
    h: = db.QueryRow(g, id3)
    switch n: = h.Scan( & code, & msg);
    n {
        case sql.ErrNoRows:
            log.Fatal(n)
        case nil:
            d[0] = code
            d[1] = msg
            return d
        default:
            log.Fatal(n)
    }
}
