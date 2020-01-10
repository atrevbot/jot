# jot

Command line tool for tracking time. Inspired by git, and strive for similar command syntax and workflow when adding clients and projects, and clocking time.

This project is in active development and all functionality is to be implemented per below specification.

## Functionality

- Add clients
- Add projects for clients
- Log time
    - Add block of time against project for client
    - Start/stop timer for working on project for client
- Edit previously logged time
- Print report of time for date range
- Categorize time entries
- Generate formatted invoices for categorized time for date range

## Configuration File

.jotrc || .jot/config

## Commands

### Add

Adds block of time, in hours, to client (if defined in config) w/ optional message and project information.

#### Arguments

Time: Hour increments defined as decimal number or fraction w/o spaces (e.g. 0.25, 1/4)

Client: The handle of the client to record time against. Client must exist in config file from previous create command

#### Options

`-m`/`--message`: An optional message to record w/ added time

`-p`/`--project`: An optional project to record time against for specified client. If project does not exist for client, a new one will be created. Time entries can be updated to change project later.

#### Example

```shell
jot add 1.5 <client> -m <message> -p <project>
```

### Start

Starts timer for work against client (if defined in config)

#### Arguments

Client

#### Example

```shell
jot start <client>
```

### Stop

Stops currently running timer and logs against optional project w/ optional message

#### Options

`-m`/`--message`

`-p`/`--project`

#### Example

```shell
jot stop -m <message> -p <project>
```

### Log

Prints log of most recent time entries w/ optional client, project, and date flags. Prints scrollable list of last time entries similar to `git log`

#### Options

`-c`/`--client`: An optional client to filter time log

`-p`/`--project`: An optional project to filter time log

`--from`: An optional flag to only print time after a given datetime (e.g. 30d). Defaults to 30 days before now.

`--to`: An optional flag to only print time until a given datetime (e.g. yesterday). Defaults to now.

#### Example

```shell
jot log -c <client> -p <project> --before <datetime> --after <datetime>
```

### Invoice

Generates invoice for client (if defined in config) for specified date range (defaults to last 30 days) w/ optional project flag

#### Arguments

Client

#### Options

`--from`: An optional flag to only invoice time after a given datetime (e.g. 30d). Defaults to 30 days before now.

`--to`: An optional flag to only invoice time until a given datetime (e.g. yesterday). Defaults to now.

#### Example

```shell
jot invoice <client> --from <datetime> --to <datetime>
```
