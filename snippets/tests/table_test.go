package main

import "testing"

func TestAdd(t *testing.T) {
    tables := []struct {            // declare table of data
        x int                   // {x, y, n}       x + y = n
        y int
        n int
    }{
        {10, 10, 20},
        {11, 23, 34},
        {7, 9, 16},
        {3, 4, 7},
    }

    for _, table := range tables {
        total := Add(table.x, table.y)
        CheckInt(total, table.n, t)
    }
}

func CheckStr(ret, expected string, t *testing.T) {
    if ret != expected {
        t.Error("Expected:", expected, "Got:", ret)
    } else {
        t.Log("SUCCESS")
    }
}

func CheckInt(ret, expected int, t *testing.T) {
    if ret != expected {
        t.Error("Expected:", expected, "Got:", ret)
    } else {
        t.Log("SUCCESS")
    }
}
