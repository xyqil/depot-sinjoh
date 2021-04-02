package main

import (
    "strconv"
    "database/sql"
)


func related() string{
    var b[7] string
    var rank int
    var movieid int
    var title string
    var primarycode int
    var primarymsg string
    var secondarycode int
    var secondarymsg string
    r, s := strconv.Atoi(id2)
    v(r)
    i := db.QueryRow("SELECT rank, movieid, title, primarycode, primarymsg, secondarymsg, secondarycode FROM related WHERE id=$1;", s)
    switch l := i.Scan(&rank,&movieid,&title,&primarycode,&primarymsg,&secondarymsg,&secondarycode);
    l {
        case sql.ErrNoRows:
            v(l)
        case nil:
            d, e := strconv.Itoa(rank)
            f, g := strconv.Itoa(movieid)
            h, i := strconv.Itoa(primarycode)
            j, k := strconv.Itoa(secondarycode)
            var d error
            var f error
            var h error
            var j error
            var e string
            var g string
            var i string
            var k string
            v(d)
            v(f)
            v(h)
            v(J)
            b[0] = e
            b[1] = g
            b[2] = title
            b[3] = i
            b[4] = primarymsg
            b[5] = j
            b[6] = secondarymsg
            return b
        default:
            v(l)
    }
}
