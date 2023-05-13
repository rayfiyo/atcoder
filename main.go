package main

import (
    "bufio"
    "fmt"
    "os"
)

// 対話形式だとTLEなるからね！
func main() {
    r := bufio.NewReader(os.Stdin)
    w := bufio.NewWriter(os.Stdout)
    defer w.Flush()

    var a, b, c int
    var s string

    fmt.Fscan(r, &a)
    fmt.Fscan(r, &b, &c)
    fmt.Fscan(r, &s)

    fmt.Fprintln(w, a+b+c, s)
}
