package util

import (
	"flag"
	"os"
)

type Flags struct {
	Mode           *string
	ConsumerKey    *string
	ConsumerSecret *string
}

func GetFlags() *Flags {
	flags := &Flags{}
	allflag := flag.NewFlagSet("flags", flag.ExitOnError)
	flags.Mode = allflag.String("mode", "production", "switch mode")
	flags.ConsumerKey = allflag.String("consumer-key", "", "consumer key")
	flags.ConsumerSecret = allflag.String("consumer-secret", "", "consumer secret")
	allflag.Parse(os.Args[1:])
	return flags
}
