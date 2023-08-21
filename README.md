# Log Writer and Tester

This repository contains a Go code that writes log entries into a JSON file, a Go test code that tests the main code, and a GitHub Actions workflow that automates testing, building, and updating the log file.

## Main Code

The `go_log.go` file contains a Go program that generates log entries and writes them into a JSON file named `logtimes.json`. Each log entry consists of a timestamp and a log number. The log number is determined based on the existing log entries in the file.

## Test Code

The `go_log_test.go` file contains tests for the main code using Go's built-in testing framework. These tests ensure that the main code behaves as expected and accurately produces log entries.

## GitHub Actions Workflow

The `.github/workflows/build-and-test.yml` file sets up a GitHub Actions workflow that triggers on push events to the main branch. It performs the following steps:
1. Run the Go tests using the Go `test` command.
2. If the tests succeed, build the main code using the Go `build` command.
3. Run the built executable.
4. Copy the generated log file to a new file named `logtimes_backup.json`.
5. Delete the original log file `logtimes.json`.
