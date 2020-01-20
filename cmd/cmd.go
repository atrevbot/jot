package cmd

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/atrevbot/jot/server"
	"github.com/atrevbot/jot/store"
	"github.com/joho/godotenv"
	bolt "go.etcd.io/bbolt"
)

const APP_ADDR = "http://localhost:8080"

func Execute(cmd string, config *os.File, db *bolt.DB) {
	switch os.Args[1] {
	case "init":
		// TODO: Initialize current working directory in global `.jot` directory.
		fmt.Println("TODO: Implement 'init' command functionality.")
	case "add":
		if len(os.Args) < 3 {
			log.Fatal("Please provide a non-zero duration")
		}

		// Parse time duration
		d, err := time.ParseDuration(os.Args[2])
		if err != nil {
			log.Fatal(err)
		}

		// TODO: Parse message from command flag -m
		m := "Test Message"

		// Get repository to create time entries
		r, err := store.New(db)
		if err != nil {
			log.Fatal(err)
		}

		// Create new time entry
		err = r.New(d, m)
		if err != nil {
			log.Fatal(err)
		}

		os.Exit(0)
	case "start":
		// TODO: Open new time entry for provided client/project at timestamp.
		fmt.Println("TODO: Implement 'start' command functionality.")
		break
	case "stop":
		// TODO: Close open time entry and add to repo.
		fmt.Println("TODO: Implement 'stop' command functionality.")
		break
	case "invoice":
		r, err := store.New(db)
		if err != nil {
			log.Fatal(err)
		}

		e, err := godotenv.Read()
		if err != nil {
			e = make(map[string]string)
		}

		// Create server and attach routes
		s := server.New(APP_ADDR, r, e)

		// Make sure server is running before opening page
		go func() {
			for {
				time.Sleep(time.Second)

				r, err := http.Get(APP_ADDR)
				if err != nil {
					continue
				}
				r.Body.Close()
				if r.StatusCode != http.StatusOK {
					continue
				}

				// Server is up and running
				break
			}
			open(APP_ADDR)
		}()

		// Start server
		log.Fatal(s.ListenAndServe())

		os.Exit(0)
	case "config":
		// TODO: Create server and load localhost browser instance to CRUD
		// clients and projects to store in config file in .jot directory.
		fmt.Println("TODO: Implement 'config' command functionality.")
		break
	default:
		// TODO: Provide short description of how to use utilitiy.
		fmt.Println("TODO: Implement help page for empty command.")
		break
	}
}

// Opens the specified URL in the default browser of the user.
func open(url string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
	default: // "linux", "freebsd", "openbsd", "netbsd"
		cmd = "xdg-open"
	}

	args = append(args, url)

	return exec.Command(cmd, args...).Start()
}
