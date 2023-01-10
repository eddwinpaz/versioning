package utils

import (
	"bufio"
	"fmt"
	"os"
)

// Get_current_version gets the current version number from the file
func Open_file(file_name string) string {
	file, err := os.Open(file_name)
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

// Save_file saves the updated version number to the file
func Save_file(version string) {
	// Save the updated version number to the file
	file, err := os.Create("./version.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	_, err = writer.WriteString(version)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	writer.Flush()
}
