package main

import (
	"log"
	"os"

	"github.com/atrevbot/jot/cmd"
	"github.com/atrevbot/jot/store"
	bolt "go.etcd.io/bbolt"
)

func main() {
	// Make sure command is passed
	if len(os.Args) == 1 {
		log.Fatal("Please provide a command. For a list of commands use the --help flag")
	}

	// Open DB connection for executing command
	db, err := bolt.Open(".jot/entries.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Get repository for time entries
	repo, err := store.New(db)
	if err != nil {
		log.Fatal(err)
	}

	cmd.Execute(os.Args[1], repo)
}
