package generate

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"text/template"
	"time"

	_ "embed"

	"github.com/spanditime/gofin/tools/gormigen/pkg/config"
	"github.com/spanditime/gofin/tools/gormigen/pkg/utils"
)

//go:embed migration.tmpl
var migrationTemplateFile string

func GenerateMigrationPackage(name string) error {
	cfg, err := config.GetConfig()
	if err != nil {
		return errors.Join(err, errors.New("failed to get config"))
	}
	date := time.Now().Format("20060102")
	migrations, err := utils.ParseMigrations(cfg.Migrations.Path)
	if err != nil {
		return errors.Join(err, errors.New("failed to parse migrations"))
	}

	// get todays number

	todaysNumber := 1
	if len(migrations) != 0 {
		// get last migration version
		lastMigration := migrations[len(migrations)-1]

		todaysNumber, err = getTodaysNumber(lastMigration, date)
		if err != nil {
			return errors.Join(err, errors.New("failed to get todays migration number"))
		}
	}

	// format name to snake_case and leave only space and letters, and remove duplicates
	name_snake_case := strings.ToLower(strings.ReplaceAll(name, " ", "_"))
	name_snake_case = regexp.MustCompile(`[^a-z0-9\s_]`).ReplaceAllString(name_snake_case, "")
	name_snake_case = strings.Join(strings.Fields(name_snake_case), "_")
	if len(name_snake_case) != 0 && name_snake_case[len(name_snake_case)-1] == '_' {
		name_snake_case = name_snake_case[:len(name_snake_case)-1]
	}

	// generate new migration version
	version := fmt.Sprintf("v%s%04d", date, todaysNumber)
	folderName := fmt.Sprintf("%s/%s-%s", cfg.Migrations.Path, version, name_snake_case)

	// create directory
	os.MkdirAll(folderName, 0755)
	// create files
	file, err := os.Create(fmt.Sprintf("%s/migration.go", folderName))
	if err != nil {
		return errors.Join(err, errors.New("failed to create migration file"))
	}
	defer file.Close()

	// generate data for template
	var data = struct {
		PackageName string
		Name        string
	}{
		PackageName: version,
		Name:        name,
	}
	// execute template
	err = template.Must(template.New("migration").Parse(migrationTemplateFile)).Execute(file, data)
	if err != nil {
		return errors.Join(err, errors.New("failed to execute template"))
	}

	return nil
}

func getTodaysNumber(lastMigration utils.Migration, date string) (int, error) {
	lastMigrationDate := lastMigration.Version[0:8]
	if lastMigrationDate > date {
		return 0, errors.New("last migration date is greater than current date")
	} else if lastMigrationDate == date {
		todaysLastNumber, err := strconv.Atoi(lastMigration.Version[8:12])
		if err != nil {
			return 0, errors.Join(err, errors.New("failed to convert last migration number to int"))
		}
		return todaysLastNumber + 1, nil
	}
	return 1, nil
}
