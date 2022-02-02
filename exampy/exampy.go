package main

// trying to fix xargs & grep flags
import (
    "C"
)

//export examplefunc
func examplefunc(x, y int) int {
    return x+y
}

func main() {}
