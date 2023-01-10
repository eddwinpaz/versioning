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

	current_version := utils.Open_file(file_path)
	fmt.Println("Current version: ", current_version)

	version := entity.Version{}
	new_version := version.Increment_version(current_version, increment_schema)

	fmt.Println("New version: ", new_version)

	utils.Save_file(new_version)
}
