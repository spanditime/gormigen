package utils

import (
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
)

type Migration struct {
	Path    string
	Package string
	Version string
	Name    string
}

func ParseMigrations(path string) ([]Migration, error) {
	// get all directories in path
	dirs, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}
	// filter directories by name format v{{version}}-{{name}}
	migrations := []Migration{}
	for _, dir := range dirs {
		if !dir.IsDir() {
			continue
		}
		// check if directory name is in format v{{version}}-{{name}}
		if !regexp.MustCompile(`^v\d{12}-.+$`).MatchString(dir.Name()) {
			log.Printf("skipping directory %s not in format v{version}-{name}\n", dir.Name())
			continue
		}
		// get version and name from directory name
		version := strings.Split(dir.Name(), "-")[0][1:]
		name := strings.Split(dir.Name(), "-")[1]
		// add migration to list
		migrations = append(migrations, Migration{
			Path:    dir.Name(),
			Package: "v" + version,
			Version: version,
			Name:    name,
		})
	}
	// sort migrations by version lexicographically
	sort.Slice(migrations, func(i, j int) bool {
		return migrations[i].Version < migrations[j].Version // ascending order
	})
	return migrations, nil
}
