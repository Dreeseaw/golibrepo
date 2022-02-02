package main

// trying to fix xargs flags
import (
    "C"
)

//export examplefunc
func examplefunc(x, y int) int {
    return x+y
}

func main() {}
