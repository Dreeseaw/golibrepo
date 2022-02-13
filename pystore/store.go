/*
    Uses kelindar/column for single-node, in-mem 
    storage & operations
*/
package main

import (
    "fmt"
    "errors"

    "github.com/kelindar/column"
)

type ColFunc func() column.Column

var CollectionTypeMap = map[string]ColFunc{
    "string": column.ForEnum,
    "float": column.ForFloat64,
    "int": column.ForInt16,
    "bool": column.ForBool,
}

type Store struct {
    CollMap map[string]*column.Collection
}

func NewStore() *Store {
    cm := make(map[string]*column.Collection)
    ret := Store{
        CollMap: cm,
    }
    return &ret
}

func (s *Store) NewCollection(name string, colls map[string]string) error {
    if _, exists := s.CollMap[name]; exists {
        return errors.New("Collection with that name already exists")
    }
    newColl := column.NewCollection()
    newColl.CreateColumn("serial", column.ForKey())
    for colname, coltype := range colls {
        if colfunc, exists := CollectionTypeMap[coltype]; !exists {
            newColl.CreateColumn(colname, colfunc())
        } else {
            return fmt.Errorf("Column type does not exist: %s", coltype)
        }
    }
    s.CollMap[name] = newColl
    return nil
}
