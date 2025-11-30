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
	fullName := name

	migrations, err := utils.ParseMigrations(cfg.Path)
	if err != nil {
		return errors.Join(err, errors.New("failed to parse migrations"))
	}

	if checkNewerMigrationsExist && NewerMigrationsExist(migrations, date) {
		return errors.New("newer migrations exist")
	}
	number := getNextNumber(migrations, date)

	migration := utils.NewMigration(date, number, fullName)

	// generate new migration version
	return generateMigrationFile(migration, cfg.Path)
}

func generateMigrationFile(migration utils.Migration, dir string) error {
	// create directory
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return errors.Join(err, errors.New("failed to create migration directory"))
	}
	// create file
	file, err := os.Create(fmt.Sprintf("%s/%s.go", dir, migration.Dir()))
	if err != nil {
		return errors.Join(err, errors.New("failed to create migration file"))
	}
	defer file.Close()
	// execute template
	return template.Must(template.New("migration").Parse(migrationTemplateFile)).Execute(file, migration)
}

func NewerMigrationsExist(migrations []utils.Migration, datetime time.Time) bool {
	date := utils.DateFromTime(datetime)
	for _, migration := range migrations {
		if migration.Datetime().After(date) {
			return true
		}
	}
	return false
}

// set the number of the test migration to the number of the last migration + 1
// migrations must be sorted by datetime
func getNextNumber(migrations []utils.Migration, datetime time.Time) int {
	if len(migrations) == 0 {
		return 1
	}
	date := utils.DateFromTime(datetime)
	highestNumber := 0
	for _, migration := range migrations {
		currentDate := migration.Datetime()
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
