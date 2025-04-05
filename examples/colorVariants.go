package main

import (
	"fmt"

	"github.com/stefan-muehlebach/gg/colors"
)

func main() {
	c1 := colors.Map["DarkCyan"]
	fmt.Printf("%v\n", c1.Dark(0.3))
	fmt.Printf("%v\n", c1.Bright(0.3))
	fmt.Printf("%v\n", c1.Bright(0.2))
}
