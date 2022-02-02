package main

import (
    //"fmt"
    //"reflect"
    "testing"

    "github.com/stretchr/testify/assert"
)

func TestExampleFunc(t *testing.T) {
    a, b := 11,22
    c := examplefunc(a, b)
    assert.Equal(t, c, 33)
}
