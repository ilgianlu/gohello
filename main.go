package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Hello())
	fmt.Println("ended")
	fmt.Println(quote.Hello())
}
