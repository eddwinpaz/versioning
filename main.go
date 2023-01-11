package main

import (
	"fmt"
	"os"

	"github.com/eddwinpaz/versioning/entity"
	"github.com/eddwinpaz/versioning/utils"
)

func main() {

	if len(os.Args[1]) < 3 {
		fmt.Println("Error: You must provide a command line argument (minor, patch, or mayor) and file path.")
		os.Exit(1)
	}

	increment_schema := os.Args[1]
	file_path := os.Args[2]

	f := utils.NewFileManager(file_path)

	fmt.Println("Current version: ", f.Text)

	version := entity.Version{}
	new_version := version.Increment_version(f.Text, increment_schema)

	fmt.Println("New version: ", new_version)
	f.Text = new_version
	f.Save_file()
}
