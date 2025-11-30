package utils

import (
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

const MigrationInactiveDefaultMarker = "xXxXx"

const migrationDirRegexString = `^v(\d{8})(\d{4})([Xx]*)-([a-z0-9_]+)$`
const migrationDirPathFormat = "v%s%04d%s-%s"
const migrationDateLayout = "20060102"

var migrationDirRegex = regexp.MustCompile(migrationDirRegexString)

const (
	migDirFDateIndex = iota + 1
	migDirFNumberIndex
	migDirFActiveIndex
	migDirFNameIndex
)

type Migration struct {
	Datetime time.Time
	Number   int
	FullName string
	Active   bool
}

func (m *Migration) Date() string {
	return m.Datetime.Format(migrationDateLayout)
}
func (m *Migration) SetDate(date string) error {
	datetime, err := time.Parse(migrationDateLayout, date)
	if err != nil {
		return errors.Join(err, errors.New("failed to parse date"))
	}
	m.Datetime = datetime
	return nil
}

func (m *Migration) Name() string {
	name_snake_case := strings.ToLower(strings.ReplaceAll(m.FullName, " ", "_"))
	name_snake_case = regexp.MustCompile(`[^a-z0-9\s_]`).ReplaceAllString(name_snake_case, "")
	name_snake_case = strings.Join(strings.Fields(name_snake_case), "_")
	if len(name_snake_case) != 0 && name_snake_case[len(name_snake_case)-1] == '_' {
		name_snake_case = name_snake_case[:len(name_snake_case)-1]
	}
	return name_snake_case
}

func (m *Migration) Dir() string {
	return fmt.Sprintf(migrationDirPathFormat, m.Date(), m.Number, m.getInactiveMarker(), m.Name())
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
	return fmt.Sprintf("%s%04d", m.Date(), m.Number)
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
		migration := Migration{
			Number:   number,
			FullName: groups[migDirFNameIndex],
			Active:   active,
		}
		migration.SetDate(groups[migDirFDateIndex])
		migrations = append(migrations, migration)
	}
	// sort migrations by version lexicographically
	sort.Slice(migrations, func(i, j int) bool {
		return migrations[i].Version() < migrations[j].Version() // ascending order
	})
	return migrations, nil
}
