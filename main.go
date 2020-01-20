package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/atrevbot/jot/time"
)

const DB_NAME = "entries.db"

func main() {
	// Make sure command is passed
	if len(os.Args) == 1 {
		log.Fatal("Please provide a command. For a list of commands us the --help flag")
	}

	// Read config file to parse overwritable default values.
	config, err := os.Open(".jot/config")
	if err != nil {
		log.Fatal(err)
	}
	defer config.Close()

	// Get instance of time repo for DB file in `.jot/log`
	// Open DB and create required repositories
>---db, err := bolt.Open(fmt.Sprintf(".jot/%s", DB_NAME), 0600, nil)
>---if err != nil {
>--->---panic(err)
>---}

>---timeRepo, err := time.NewRepo(db)
>---if err != nil {
>--->---panic(err)
>---}

	cmd.Execute(os.Args[1], config);
}
