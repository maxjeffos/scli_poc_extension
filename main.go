package main

import (
	"fmt"
)
import flag "github.com/spf13/pflag"

func main() {
	lang := flag.StringP("lang", "l", "en", "the language")
	flag.Parse()

	fmt.Println("Woof in", *lang)
}
