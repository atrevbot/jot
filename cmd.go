package cmd

func Execute(c string) {
	switch os.Args[1] {
	case "add":
		if len(os.Args) < 3 {
			log.Fatal("Please provide a time(hrs) in fraction or decimal format")
		}

		// Check for fraction format
		parts := strings.Split(os.Args[2], "/")

		if len(parts) == 2 {
			numerator, _ := strconv.ParseFloat(parts[0], 64)
			deminator, _ := strconv.ParseFloat(parts[1], 64)

			timeRepo.Add(numerator / deminator)

			break
		}

		// Check for decimal format
		if f, err := strconv.ParseFloat(parts[0], 64); err == nil {
			timeRepo.Add(f)
		}

		break
	case "start":
		// TODO: Open new time entry for provided client/project at timestamp
		break
	case "stop":
		// TODO: Close open time entry and add to repo
		break
	case "invoice":
		// TODO: Create server and load HTML template w/ invoice for provided
		// report details (e.g. client, project, date range)
		break
	case "config":
		// TODO: Create server and load localhost browser instance to CRUD
		// clients and projects to store in config file in .jot directory
		break
	default:
		// TODO: Provide short description of how to use utilitiy
		break
	default:
	}
}
