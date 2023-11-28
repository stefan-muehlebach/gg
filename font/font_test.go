package font

import (
    "fmt"
    "testing"
)

func TestGoFont(t *testing.T) {
    fmt.Printf("> GoRegular: %v\n", GoRegular)
}

func TestTTFFont(t *testing.T) {
    fmt.Printf("> LucidaBright: %v\n", LucidaBright)
}

