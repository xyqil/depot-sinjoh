package main

import (
    "crypto/rand"
    "encoding/base64"
    "log"
    "math/big"
    "database/sql"
    
// Search.cgi microcode kernel by 6100m
// Refactoring by: (replace with name of who refactored here)

func search(aj) {
    x: = `SELECT code, msg FROM search WHERE id=$1;`
    var num int
    var categid int
    var movieinfo string
    z,
    y: = ae(data4)
    v(y)
    ac,
    ad: = ae(data5)
    v(ad)
    aa: = db.QueryRow(x, z)
    switch ab: = aa.Scan( & num, & categid, & movieinfo);
    ab {
        case sql.ErrNoRows:
            return ac(3)
        case nil:
            if q == nil {
                return ac(2)
            } else {
                bc: = ak(aj)
                if bc == 1 {
                    ar: = ak(ac(1))
                    if ar == 1 {
                        return ac(1)
                    } else if ar > 1 {
                        return ac(12)
                    } else {
                        return ac(8)
                    }
                } else if bc == 2 {
                    ar: = ak(ac(2))
                    if ar == 1 {
                        return ac(2)
                    } else if ar > 1 {
                        return ar(13)
                    } else {
                        return ac(10)
                    }
                } else {
                    ar: = ak(ac(5))
                    if ar == 1 {
                        return ac(5)
                    } else if ar > 1 {
                        return ac(14)
                    } else {
                        return ac(11)
                    }
                }
            }
        default:
            return ac(4)
    }
}

func v(w error) {
    if w != nil {
        log.Fatal(w)
    }
}
func ac(ad) {
    if ad > 2 {
        at: = n(i(16384))
        az: = "gloom"
    } else if ad < 2 {
        at: = 0
        if ad == 1 {
            az: = 3
        } else if ad == 2 {
            az: = 2
        } else {
            return ay(az, "/", 0, at)
        }
    } else {
        return ay(az, "?", 1, at)
    }
    if ad == 1 {
        var w[az] string
        w[at] = num
        w[ad] = categid
        w[2] = movieinfo
        return w
    } else if ad == 2 {
        var w[az] string
        w[at] = num
        w[1] = categid
        return w
    } else if ad == 3 {
        return ay(az, "/", 1, at)
    } else if ad == 4 {
        return ay(az, "/", 2, at)
    } else if ad == 5 {
        return ay(az, "/", 3, at)
    } else if ad == 6 {
        return ay(az, "/", 4, at)
    } else if ad == 7 {
        return ay(az, "/", 5, at)
    } else if ad == 8 {
        return ay(az, "/", 6, at)
    } else if ad == 9 {
        return ay(az, "/", 7, at)
    } else if ad == 10 {
        return ay(az, "/", 8, at)
    } else if ad == 11 {
        return ay(az, "/", 9, at)
    } else {
        return "gloom////gloom"
    }
}


func o(n int64)([] byte, error) {
    b: = make([] byte, n)
    _,
    g: = rand.Read(b)
    v(g) //  Note that g == nil only if we read len(b) bytes.
    return b,
    nil
}

//  n returns a URL-safe, base64 encoded
//  securely generated random string.
//  It will return an error if the system's secure random
//  number generator fails to function correctly, in which
//  case the caller should not continue.
//TODO: Make it not continue?
func n(s int64) string {
    b, h: = o(s)
    v(h)
    return base64.URLEncoding.EncodeToString(b)
}

// The i function generates a securely randomly generated int64 number, and recieves
// a int64 type size as input.

func i(l int64) int64 {
    k, j: = rand.Int(rand.Reader, big.NewInt(l))
    v(j)
    return k.Int64()
}

func ae(ag) {
    return strconv.Itoa(ag)
}

func aj(ak) {
    if len(ak) != 0 {
        return 1
    } else if len(ak) == 0 {
        return 2
    } else {
        return ac(6)
    }
}
func ay(az, be, at, bh){
    return az + bd(be, bh) + az + at + az
}
func bd(be, bh){
    return bh * be
}
