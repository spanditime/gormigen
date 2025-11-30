package generate

import (
	"errors"
	"os"
	"text/template"

	_ "embed"

	"github.com/spanditime/gofin/tools/gormigen/internal/config"
	"github.com/spanditime/gofin/tools/gormigen/internal/utils"
)

//go:embed migration_manager.tmpl
var migrationManagerTemplate string

func GenerateMigrationManagerFile(migrations []utils.Migration, cfg *config.Config) error {
	// write file
	file, err := os.Create(cfg.Output.Filename)
	if err != nil {
		return err
	}
	defer file.Close()
	// get module from go.mod
	module, err := utils.GetModuleFromGoMod()
	if err != nil {
		return errors.Join(err, errors.New("failed to get module from go.mod"))
	}

	// generate data for template
	var data = struct {
		Module     string
		Config     *config.Config
		Migrations []utils.Migration
	}{
		Module:     module,
		Config:     cfg,
		Migrations: migrations,
	}
	// execute template
	err = template.Must(template.New("migration_manager").Parse(migrationManagerTemplate)).Execute(file, data)
	if err != nil {
		return errors.Join(err, errors.New("failed to execute template"))
	}
	return nil
}
