package main

import (
    "log"
    "strconv"
    "database/sql"
)

func miiinfo() {
    var a[2] string
    e: = `SELECT code, msg FROM miiinfo WHERE id=$1;`
    var code int
    var msg string
    p, o: = strconv.Atoi(id1)
    check(o)
    j: = db.QueryRow(e, p)
    switch k: = j.Scan( & code, & msg);
    k {
        case sql.ErrNoRows:
            log.Fatal(k)
        case nil:
            a[0] = code
            a[1] = msg
            return a
        default:
            v(k)
    }
}

func related() {
    var b[3] string
    f: = `SELECT rank, movieid, title FROM related WHERE id=$1;`
    var rank int
    var movieid int
    var title string
    r, s: = strconv.Atoi(id2)
    check(r)
    i: = db.QueryRow(f, s)
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
            v(l)
    }
}

func evaluate() {
    var d[2] string
    g: = `SELECT code, msg FROM evaluate WHERE id=$1;`
    var code int
    var msg string
    t, u: = strconv.Atoi(id3)
    check(u)
    h: = db.QueryRow(g, t)
    switch n: = h.Scan( & code, & msg);
    n {
        case sql.ErrNoRows:
            log.Fatal(n)
        case nil:
            d[0] = code
            d[1] = msg
            return d
        default:
            v(n)
    }
}

func v(e error) {
    if w != nil {
        log.Fatal(w)
    }
}
