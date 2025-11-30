package config

import (
	"errors"
	"os"

	"gopkg.in/yaml.v3"
)

type MigrationsConfig struct {
	Path string `yaml:"path"`
}

type OutputConfig struct {
	Filename    string `yaml:"filename"`
	PackageName string `yaml:"package_name"`
}

type InitMigrationConfig struct {
	Package string `yaml:"package"`
}

type Config struct {
	Migrations    MigrationsConfig     `yaml:"migrations"`
	Output        OutputConfig         `yaml:"output"`
	InitMigration *InitMigrationConfig `yaml:"init_migration"`
}

var defaultConfig = &Config{
	Migrations: MigrationsConfig{
		Path: "db/migrations",
	},
	Output: OutputConfig{
		Filename:    "db/migrations.generated.go",
		PackageName: "db_migrations",
	},
}

func GetConfig() (*Config, error) {
	cfg := defaultConfig
	// check if gormigen.yml or gormigen.yaml exists
	_, erryml := os.Stat("./gormigen.yml")
	if os.IsNotExist(erryml) {
		return nil, errors.New("gormigen.yml not found, run 'gormigen init' to create it")
	}
	// read gormigen.yml
	config, err := os.ReadFile("./gormigen.yml")
	if err != nil {
		return nil, errors.Join(err, errors.New("failed to read gormigen.yml"))
	}
	err = yaml.Unmarshal(config, &cfg)
	if err != nil {
		return nil, errors.Join(err, errors.New("failed to unmarshal gormigen.yml"))
	}
	return cfg, nil
}
