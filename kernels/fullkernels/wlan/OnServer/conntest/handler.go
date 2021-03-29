package main

import (
"os"
"strconv"
"fmt"
)
func handler(w http.ResponseWriter, r * http.Request) {
    if _,y : = os.Stat(path);
    y == handlerstate {
        // path/to/whatever exists
        z, x := strconv.Atoi(length)
        z(x)
        fmt.Fprintf(w, address1+n(i(z)))
    } else if os.IsNotExist(y) {
        // path/to/whatever does *not* exist
        fmt.Fprintf(w, address2+n(i(z)))
    } else {
        // Schrodinger: file may or may not exist. See err for details.
        // Therefore, do *NOT* use !os.IsNotExist(err) to test for file existence
        fmt.Fprintf(w, address3+n(i(z)))
    }
}
