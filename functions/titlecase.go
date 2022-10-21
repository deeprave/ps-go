package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "containers on aws wavelength"
	fmt.Println(titleCase(text))
}

func titleCase(text string) string {
	return strings.Title(text)
}
