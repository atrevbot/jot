package main

import (
	"fmt"
	"log"
	"os"

	"github.com/atrevbot/jot/cmd"
	bolt "go.etcd.io/bbolt"
)

const DB_NAME = "entries.db"

func main() {
	// Make sure command is passed
	if len(os.Args) == 1 {
		log.Fatal("Please provide a command. For a list of commands use the --help flag")
	}

	// Read config file to parse overwritable default values.
	config, err := os.Open(".jot/config")
	if err != nil {
		log.Fatal(err)
	}
	defer config.Close()

	// Open DB connection for executing command
	db, err := bolt.Open(fmt.Sprintf(".jot/%s", DB_NAME), 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	cmd.Execute(os.Args[1], config, db)
}
