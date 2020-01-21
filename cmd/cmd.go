package cmd

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/atrevbot/jot/server"
	"github.com/atrevbot/jot/store"
	bolt "go.etcd.io/bbolt"
)

const APP_ADDR = "http://localhost:8080"

func Execute(cmd string, config *os.File, db *bolt.DB) {
	switch os.Args[1] {
	case "init":
		// TODO: Initialize current working directory in global `.jot` directory.
		fmt.Println("TODO: Implement 'init' command functionality.")
		os.Exit(0)
	case "add":
		if len(os.Args) < 3 {
			log.Fatal("Please provide a non-zero duration")
		}

		// Parse time duration
		d, err := time.ParseDuration(os.Args[2])
		if err != nil {
			log.Fatal(err)
		}

		// Parse message from command flag -m
		var m string
		logFlag := flag.NewFlagSet("log", flag.ExitOnError)
		logFlag.StringVar(&m, "m", "", "Description of time entry")
		logFlag.Parse(os.Args[3:])

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
	case "log":
		// Get repository to create time entries
		r, err := store.New(db)
		if err != nil {
			log.Fatal(err)
		}

		// Get all entries
		es, err := r.All()
		if err != nil {
			log.Fatal(err)
		}

		for _, e := range es {
			fmt.Printf(
				"Spent %s on %s - %s \n",
				e.Duration,
				e.Message,
				e.Created.Format("Mon Jan 2 @3:04PM"),
			)
		}

		os.Exit(0)
	case "start":
		// TODO: Open new time entry for provided client/project at timestamp.
		fmt.Println("TODO: Implement 'start' command functionality.")
		os.Exit(0)
	case "stop":
		// TODO: Close open time entry and add to repo.
		fmt.Println("TODO: Implement 'stop' command functionality.")
		os.Exit(0)
	case "invoice":
		r, err := store.New(db)
		if err != nil {
			log.Fatal(err)
		}

		// Create server and attach routes
		s := server.New(APP_ADDR, r)

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
	default:
		// TODO: Provide short description of how to use utilitiy.
		fmt.Println("TODO: Implement help page for empty command.")
		os.Exit(0)
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
