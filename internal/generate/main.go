package generate

import (
	"errors"

	"github.com/spanditime/gofin/tools/gormigen/internal/config"
	"github.com/spanditime/gofin/tools/gormigen/internal/utils"
)

const (
	UpFunctionName   = "Up"
	DownFunctionName = "Down"
)

func Generate(cfg *config.Config) error {
	// parse migrations directories
	migrations, err := utils.ParseMigrations(cfg.Migrations.Path)
	if err != nil {
		return errors.Join(err, errors.New("failed to parse migrations"))
	}
	// generate migration manager file
	err = GenerateMigrationManagerFile(migrations, cfg)
	if err != nil {
		return errors.Join(err, errors.New("failed to generate migration manager file"))
	}
	return nil
}
