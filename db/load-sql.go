package db

import (
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

func readSQL(folderPath string) (map[string]string, error) {
	sqlfiles := make(map[string]string)
	err := filepath.WalkDir(folderPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			log.Printf("Error accessing path %s: %s", path, err)
		}
		if d.IsDir() {
			return nil
		}
		contents, err := os.ReadFile(path)
		if err != nil {
			log.Printf("Error Reading file %s: %s", path, err)
		}
		sqlfiles[path] = string(contents)
		return nil
	})

	if err != nil {
		log.Printf("Error walking %s: %s", folderPath, err)
	}

	log.Printf("Map: %v", sqlfiles)
	return sqlfiles, nil
}
