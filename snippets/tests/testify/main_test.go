package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestCalculateVanilla(t *testing.T) {
    if Calculate(2) != 4 {
        t.Error("Expected 2 + 2 to equal 4")
    }
}

func TestCalculateTestify(t *testing.T) {
  assert.Equal(t, Calculate(2), 4)
}

func TestCalculateTables(t *testing.T) {
    assert := assert.New(t)

    var tests = []struct {
        input    int
        expected int
    }{
        {2, 4},
        {-1, 1},
        {0, 2},
        {-5, -3},
        {99999, 100001},
    }

    for _, test := range tests {
        assert.Equal(Calculate(test.input), test.expected)
    }
}
