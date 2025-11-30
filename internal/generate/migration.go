package generate

import (
	"errors"
	"fmt"
	"os"
	"text/template"
	"time"

	_ "embed"

	"github.com/spanditime/gofin/tools/gormigen/internal/config"
	"github.com/spanditime/gofin/tools/gormigen/internal/utils"
)

//go:embed migration.tmpl
var migrationTemplateFile string

func GenerateMigrationPackage(cfg config.MigrationsConfig, date time.Time, name string, checkNewerMigrationsExist bool) error {
	var migration utils.Migration
	migration.Datetime = date
	migration.FullName = name

	migrations, err := utils.ParseMigrations(cfg.Path)
	if err != nil {
		return errors.Join(err, errors.New("failed to parse migrations"))
	}

	if checkNewerMigrationsExist && NewerMigrationsExist(migrations, migration) {
		return errors.New("newer migrations exist")
	}
	migration.Number = getNextNumber(migrations, date)

	// generate new migration version
	return generateMigrationFile(migration, cfg.Path)
}

func generateMigrationFile(migration utils.Migration, dir string) error {
	// create directory
	os.MkdirAll(dir, 0755)
	// create file
	file, err := os.Create(fmt.Sprintf("%s/%s.go", dir, migration.Dir()))
	if err != nil {
		return errors.Join(err, errors.New("failed to create migration file"))
	}
	defer file.Close()
	// execute template
	return template.Must(template.New("migration").Parse(migrationTemplateFile)).Execute(file, migration)
}

func NewerMigrationsExist(migrations []utils.Migration, test utils.Migration) bool {
	for _, migration := range migrations {
		if migration.Datetime.After(test.Datetime) {
			return true
		}
	}
	return false
}

func dateFromTime(date time.Time) time.Time {
	return time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
}

// set the number of the test migration to the number of the last migration + 1
// migrations must be sorted by datetime
func getNextNumber(migrations []utils.Migration, datetime time.Time) int {
	if len(migrations) == 0 {
		return 1
	}
	date := dateFromTime(datetime)
	highestNumber := 0
	for _, migration := range migrations {
		currentDate := dateFromTime(migration.Datetime)
		if currentDate.After(date) {
			break
		}
		if currentDate.Equal(date) {
			if migration.Number > highestNumber {
				highestNumber = migration.Number
			}
			continue
		}
	}
	return highestNumber + 1
}
