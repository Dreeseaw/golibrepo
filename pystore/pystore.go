/*
    CPython interface to cache
*/
package main

import (
    "C"
)

var StorePtr *Store = nil

//export Initstore
func Initstore() bool {
    if StorePtr != nil {
        return false
    }
    StorePtr = NewStore()
    return true
}

//export Newcoll
func Newcoll(name string, cols map[string]string) bool {
    if StorePtr == nil {
        return false
    }
    if err := StorePtr.NewCollection(name, cols); err != nil {
        return false
    }
    return true
}

//export Set
func Set(coll string) bool {
    return false
}

//export Get
func Get(coll string) bool {
    return false
}

func main() {}
