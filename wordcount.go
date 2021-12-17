package wordcount

import (
	"fmt"
	"os"
	"strings"
)

func wordcount() {
	s := os.Args[1]
	v := strings.Fields(s)
	fmt.Println(len(v))
}