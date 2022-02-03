package main

import (
    "C"
)

//export examplefunc
func examplefunc(x, y int) int {
    z := x+y
    return z
}

func main() {}
