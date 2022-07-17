package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

func main() {
	text := stringutil.Reverse("Hello, OTUS!")
	fmt.Println(text)
}
