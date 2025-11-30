package utils

import (
	"os"

	"golang.org/x/mod/modfile"
)

func GetModuleFromGoMod() (string, error) {
	// read go.mod
	goMod, err := os.ReadFile("go.mod")
	if err != nil {
		return "", err
	}
	// parse go.mod
	goModFile, err := modfile.Parse("go.mod", goMod, nil)
	if err != nil {
		return "", err
	}
	// TODO: handle replace directive, maybe add it to the module path
	return goModFile.Module.Mod.Path, nil
}
