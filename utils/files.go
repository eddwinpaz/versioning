package utils

import (
	"bufio"
	"fmt"
	"os"
)

type FileManager struct {
	FileName string
	Text     string
}

func NewFileManager(filePath string) *FileManager {

	f := &FileManager{}
	f.FileName = filePath
	f.Open_file()
	return f
}

// Get_current_version gets the current version number from the file
func (f *FileManager) Open_file() {
	file, err := os.Open(f.FileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	version := scanner.Text()
	f.Text = version
}

// Save_file saves the updated version number to the file
func (f *FileManager) Save_file() {
	// Save the updated version number to the file
	file, err := os.Create(f.FileName)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	_, err = writer.WriteString(f.Text)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	writer.Flush()
}
