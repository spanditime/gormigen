package utils

import (
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
)

const MigrationInactiveDefaultMarker = "xXxXx"

const migrationDirRegexString = `^v(\d{8})(\d{4})([Xx]*)-([a-z0-9_]+)$`
const migrationDirPathFormat = "v%s%04d%s-%s"

var migrationDirRegex = regexp.MustCompile(migrationDirRegexString)

const (
	migDirFDateIndex = iota + 1
	migDirFNumberIndex
	migDirFActiveIndex
	migDirFNameIndex
)

type Migration struct {
	Date   string
	Number int
	Name   string
	Active bool
}

func (m *Migration) Path() string {
	return fmt.Sprintf(migrationDirPathFormat, m.Date, m.Number, m.getInactiveMarker(), m.Name)
}

func (m *Migration) getInactiveMarker() string {
	if m.Active {
		return ""
	}
	return MigrationInactiveDefaultMarker
}

func (m *Migration) Package() string {
	return "v" + m.Version()
}

func (m *Migration) Version() string {
	return m.Date + fmt.Sprintf("%04d", m.Number)
}

func ParseMigrations(path string) ([]Migration, error) {
	// get all directories in path
	dirs, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}
	migrations := []Migration{}
	for _, dir := range dirs {
		if !dir.IsDir() {
			continue
		}
		// check if directory name is in valid format
		// get groups from regex
		groups := migrationDirRegex.FindStringSubmatch(dir.Name())
		if groups == nil {
			log.Printf("skipping directory %s not in format %s\n", dir.Name(), migrationDirRegexString)
			continue
		}
		number, err := strconv.Atoi(groups[migDirFNumberIndex])
		if err != nil {
			return nil, errors.Join(err, errors.New("failed to convert migration number to int"))
		}
		active := len(groups[migDirFActiveIndex]) > 0
		if !active {
			active = true
			log.Printf("migration %s is marked as not active, which is currently not supported. Assuming it is active.", dir.Name())
		}
		migrations = append(migrations, Migration{
			Date:   groups[migDirFDateIndex],
			Number: number,
			Name:   groups[migDirFNameIndex],
			Active: active,
		})
	}
	// sort migrations by version lexicographically
	sort.Slice(migrations, func(i, j int) bool {
		return migrations[i].Version() < migrations[j].Version() // ascending order
	})
	return migrations, nil
}
