package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/alwindoss/dux/internal/engine"
)

var dbLoc string

func init() {
	flag.StringVar(&dbLoc, "db-path", "", "-db-path='/home/user/dbase'")
}

func main() {
	flag.Parse()
	if !valid() {
		fmt.Println("db-path flag is missing pass the flag -db-path=<DB_PATH>")
		os.Exit(0)
	}
	err := engine.Run(dbLoc)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Exiting dux...")
}

func valid() bool {
	if dbLoc == "" {
		return false
	}
	return true
}
