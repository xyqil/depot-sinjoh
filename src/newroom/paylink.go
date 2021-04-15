package main

import (
	"database/sql"
	"strconv"
)

func limited() string {
	var integerdata int
	var id4 int
	var t [4]string
	z, y := strconv.Itoa(integerdata)
	x, u := strconv.Itoa(id3)
	v(z)
	v(x)
	n := db.QueryRow("SELECT primarymovieid, primarycategoryid, secondarymovieid, secondarycategoryid, FROM paylink WHERE integerdata=$1;", u)
	switch s := n.Scan(&primarymovieid, &primarycategoryid, &secondarymovieid, &secondarycategoryid); s {
	case sql.ErrNoRows:
		v(s)
	case nil:
		t[0] = primarymovieid
		t[1] = primarycategoryid
		t[2] = secondarymovieid
		t[3] = secondarycategoryid
		return t
	default:
		v(s)
	}
}
