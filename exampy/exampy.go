package exampy

import (
    "C"
)

//export examplefunc
func examplefunc(x, y int) int {
    return x+y
}

func main() {}
