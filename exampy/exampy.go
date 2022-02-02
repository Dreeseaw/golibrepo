package main

// another meaningless comment
import (
    "C"
)

//export examplefunc
func examplefunc(x, y int) int {
    return x+y
}

func main() {}
