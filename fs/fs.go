package fs

import (
	"encoding/csv"
	"log"
	"os"
	"path/filepath"
)

// GetThisDir is used to get the directory of the currently running program
func GetThisDir() (string, error) {
	return filepath.Abs(filepath.Dir(os.Args[0]))
}

// Write is used to write data to a file (and creates it if it doesn't exist)
func Write(path, data string) error {
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0)
	defer f.Close()
	if err != nil {
		return err
	}

	_, err = f.Write([]byte(data))
	if err != nil {
		return err
	}

	return nil
}

// WriteTo is used to append data to an existing file
func WriteTo(path, data string) error {
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND, 0644)
	defer f.Close()
	if err != nil {
		return err
	}

	_, err = f.Write([]byte(data))
	if err != nil {
		return err
	}

	return nil
}

// ReadCSV takes a csv file and returns the records
func ReadCSV(path string) ([][]string, error) {
	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		return nil, err
	}

	r := csv.NewReader(f)

	records, err := r.ReadAll()
	if err != nil {
		return nil, err
	}

	return records, nil
}

// FileExists checks if file exists
func FileExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
		log.Fatal(err)
	}
	return true
}

// Mkdir is used to make directories with default permissions of: rwx,rx,r
func Mkdir(path string) error {
	return os.Mkdir(path, 0751)
}
