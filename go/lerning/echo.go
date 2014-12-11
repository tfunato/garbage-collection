package main

import (
    "os"
    "flag"
)

var omitNewline = flag.Bool("n", false, "don't print final newline")

const (
    Space   = " "
    Newline = "\n"
)

func main() {
    flag.Parse() // パラメータリストを調べてflagに設定
    var s string = ""
    for i := 0; i < flag.NArg(); i++ {
        if i > 0 {
            s += Space
        }
        s += flag.Arg(i)
    }
    if !*omitNewline {
        s += Newline
    }
    os.Stdout.WriteString(s)
}
