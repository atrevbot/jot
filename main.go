package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type timeEntry struct {
	time float64
}

func addTime(time float64) {
	// TODO: Log time to bolt DB in file in `.jot` config directory using
	// timestamp as seed for hash to generate key for time entry
	fmt.Printf("Time: %v\n", time)
}

func main() {
	// Read config file to parse overwritable default values.
	config, err := os.Open(".jot/config")
	if err != nil {
		log.Fatal(err)
	}
	defer config.Close()

	// Make sure command is passed
	if len(os.Args) < 2 {
		log.Fatal("Please provide a command. For a list of commands us the --help flag")
	}

	command := os.Args[1]

	switch command {
	case "add":
		if len(os.Args) < 3 {
			log.Fatal("Please provide an addable type")
		}

		// TODO: Check if default-add-type is defined in config
		scanner := bufio.NewScanner(config)
		for scanner.Scan() { // internally, it advances token based on sperator
			// c := strings.Split(scanner.Text(), "=") // token in unicode-char
		}
		adding := "time"

		if len(os.Args) > 2 {
			adding = os.Args[2]
		}

		switch adding {
		case "time":
			if len(os.Args) < 4 {
				log.Fatal("Please provide a time(hrs) in fraction or decimal format")
			}

			// Check for fraction format
			time := os.Args[3]
			parts := strings.Split(time, "/")

			if len(parts) == 2 {
				numerator, _ := strconv.ParseFloat(parts[0], 64)
				deminator, _ := strconv.ParseFloat(parts[1], 64)

				addTime(numerator / deminator)
			}

			// Check for decimal format
			if f, err := strconv.ParseFloat(time, 64); err == nil {
				addTime(f)
			}

			break
		default:
			fmt.Println("Only addable currently supported is \"time\"")
		}

		break
	default:
		fmt.Println("Only command currently supported is \"add\"")
	}

	// TODO: Provide short description of how to use utilitiy
}
