package main

import (
	"fmt"
	"os"
	"sort"
)

func main() {
	name := os.Getenv("USER")
	fmt.Println("Hi, ", name)

	env := os.Environ()
	sort.Strings(env)
	for _, value := range env {
		fmt.Println(value)
	}
}
