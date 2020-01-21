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
)

const APP_ADDR = "http://localhost:8080"

func Execute(cmd string, repo store.Repo) {
	switch os.Args[1] {
	case "add":
		if len(os.Args) < 3 {
			log.Fatal("Please provide a non-zero duration")
		}

		// Parse time duration
		d, err := time.ParseDuration(os.Args[2])
		if err != nil {
			log.Fatal(err)
		}

		dir, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}

		// Parse message and project from command flags
		var m string
		var p string
		logFlag := flag.NewFlagSet("log", flag.ExitOnError)
		logFlag.StringVar(&m, "m", "No Message", "Description of time entry")
		logFlag.StringVar(&p, "p", dir, "Description of time entry")
		logFlag.Parse(os.Args[3:])

		// Create new time entry
		err = repo.New(d, m, p)
		if err != nil {
			log.Fatal(err)
		}

		os.Exit(0)
	case "log":
		// Get all entries
		es, err := repo.All()
		if err != nil {
			log.Fatal(err)
		}

		for _, e := range es {
			fmt.Printf(
				"Spent %s doing '%s' for '%s' on %s \n",
				e.Duration,
				e.Message,
				e.Project,
				e.Created.Format("Mon Jan 2 @3:04PM"),
			)
		}

		os.Exit(0)
	case "invoice":
		// Create server and attach routes
		s := server.New(APP_ADDR, repo)

		// Make sure server is running before opening page
		go ping(APP_ADDR, time.Second, func() { open(APP_ADDR) })

		// Start server
		log.Fatal(s.ListenAndServe())

		os.Exit(0)
	default:
		// TODO: Provide short description of how to use utilitiy.
		fmt.Println("Supports subcommands of add, log, and invoice.")
		os.Exit(0)
	}
}

// Ping server at addres until it is ready and accepting requests
func ping(addr string, interval time.Duration, cb func()) {
	for {
		time.Sleep(interval)

		r, err := http.Get(addr)
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

	cb()
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
