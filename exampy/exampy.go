package main

// trying to fix just xargs falgs
import (
    "C"
)

//export examplefunc
func examplefunc(x, y int) int {
    return x+y
}

func main() {}
