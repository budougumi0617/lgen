package main

import (
	"flag"
	"log"
	"os"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	log.SetFlags(0)
	err := Run(os.Args, os.Stdout, os.Stderr)
	if err != nil && err != flag.ErrHelp && err != ErrShowVersion {
		log.Println(err)
		os.Exit(2)
	}
}
