package entity

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Version is the struct that holds the version number
type Version struct {
	Major int `json:"major"`
	Minor int `json:"minor"`
	Patch int `json:"patch"`
}

func (v *Version) extract_schema(version string) {

	fields := strings.Split(version, ".")
	v.Major = str2int(fields[0])
	v.Minor = str2int(fields[1])
	v.Patch = str2int(fields[2])

}

// Increment_version increments the version number
func (v *Version) Increment_version(raw_version string, increment string) string {

	v.extract_schema(raw_version)

	if increment == "minor" {
		v.Minor = v.Minor + 1
	} else if increment == "patch" {
		v.Patch = v.Patch + 1
	} else if increment == "major" {
		v.Major = v.Major + 1
	} else {
		fmt.Println("Error: Invalid command line argument")
		os.Exit(1)
	}

	return fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Patch)

}

func str2int(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return i
}
