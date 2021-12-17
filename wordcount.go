package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	s := os.Args[1]
	v := strings.Fields(s)
	fmt.Println(len(v))
}