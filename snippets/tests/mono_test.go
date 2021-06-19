package main

import "testing"

func TestGetVal(t *testing.T) {
    Check(Greet("John"), "Hello, John", t)
    Check(Greet(""), "Hello, world", t)
}

func Check(ret, expected string, t *testing.T) {
    if ret != expected {
        t.Error("Expected:", expected, "Got:", ret)
    } else {
        t.Log("SUCCESS")
    }
}
