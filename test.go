package main

import (
	"flag"
	"fmt"
	"os"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	flags := flag.NewFlagSet("test", flag.ExitOnError)
	c := flags.String("test", "pam", "test flag")
	flags.Parse(os.Args[1:])
	fmt.Println(*c)
}
