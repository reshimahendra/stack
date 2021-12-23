package main

import (
	"fmt"
	"reflect"
)

type Item struct {
    id int
    name string
}

func main() {
    a := []Item{
        Item{id: 1, name:"test 1"},
        Item{id: 2, name:"test 2"},
        Item{id: 3, name:"test 3"},
        Item{id: 4, name:"test 4"},
        Item{id: 5, name:"test 5"},
        Item{id: 6, name:"test 6"},
        Item{id: 7, name:"test 7"},
        Item{id: 8, name:"test 8"},
    }

    for v, k := range a{
        fmt.Println(v, k)
    }

    b := reflect.ValueOf(a[1])
    fmt.Println(b)
    values := make([]interface{}, b.NumField())

    fmt.Println(values...)
    for i := 0; i < b.NumField();i++ {
        // values[i] = b.Field(i).Interface()
        fmt.Println(values[i])
    }
}
