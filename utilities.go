package main

import (
	"io/fs"
	"os"
	"path/filepath"
)

func getNumFiles(path string) int {
	count := 0

	// walk path and count number of files
	err := filepath.Walk(path, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// ignore directories
		if info.IsDir() {
			return nil
		}

		count++

		return nil
	})

	if err != nil {
		panic(err)
	}

	return count
}

func writeCSVToFile(path string, text []string) {
	// create file and handle error
	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// write header to file
	_, err = f.WriteString("Bitrate (Kb/s),Path\n")
	if err != nil {
		panic(err)
	}

	// write each line to file
	for _, i := range text {
		_, err = f.WriteString(i + "\n")
		if err != nil {
			panic(err)
		}
	}
}
