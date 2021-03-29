package main

import (
"encoding/base64"
)
//  n returns a URL-safe, base64 encoded
//  securely generated random string.
//  It will return an error if the system's secure random
//  number generator fails to function correctly, in which
//  case the caller should not continue.
//TODO: Make it not continue?
func n(s int64) string {
    b, h: = o(s)
    if h != nil {
        z(h)
    }
    return base64.URLEncoding.EncodeToString(b)
}
