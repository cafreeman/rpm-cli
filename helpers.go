package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// Convert a `/` sepearated path to a valid Windows path using os.PathSeparator
func convertToWindowsPath(path string) string {
	return filepath.FromSlash(path)
}

// Catch-all function for error checking
func errCheck(err error) {
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
}

// Determine the rooted path to SVN on the machine executing this script
func svnRoot() string {
	wd, _ := os.Getwd()
	return strings.TrimSuffix(wd, `\3rdParty\R`)
}

func pause() {
	fmt.Print("Press enter to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
