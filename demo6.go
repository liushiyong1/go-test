package main

import (
	"flag"
	"fmt"
)

func main() {
	name := *flag.String("name", "everyone", "The greeting object.")
	//var name string                                                   // [1]
	//flag.StringVar(&name, "name", "everyone", "The greeting object.") // [2]
	flag.Parse()
	fmt.Printf("Hello %v\n", name)
}
