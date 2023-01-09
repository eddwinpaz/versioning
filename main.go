package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Error: You must provide a command line argument (minor, patch, or mayor)")
		os.Exit(1)
	}
	Increment_version(os.Args[1])
}

// get_current_version gets the current version number from the file
func get_current_version() string {
	file, err := os.Open("version.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	version := scanner.Text()

	return version
}

// save_file saves the updated version number to the file
func save_file(fields []string) {
	// Save the updated version number to the file
	file, err := os.Create("./version.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(strings.Join(fields, "."))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	writer.Flush()
}

// str2int is a helper function that converts a string to an integer
func str2int(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return i
}

// increment_version increments the version number based on the command line argument
func Increment_version(increment string) {

	current_version := get_current_version()
	fmt.Println("Incrementing version number..." + increment)
	fmt.Println("Current version: " + current_version)

	// Increment the version number based on the command line argument
	fields := strings.Split(current_version, ".")
	if increment == "minor" {
		fields[1] = fmt.Sprintf("%d", str2int(fields[1])+1)
	} else if increment == "patch" {
		fields[2] = fmt.Sprintf("%d", str2int(fields[2])+1)
	} else if increment == "major" {
		fields[0] = fmt.Sprintf("%d", str2int(fields[0])+1)
	} else {
		fmt.Println("Error: Invalid command line argument")
		os.Exit(1)
	}

	save_file(fields)
}
