package generate

import (
	"errors"
)

type Mode string

const (
	ModeAuto     Mode = "auto"
	ModeInplace  Mode = "inplace"
	ModePushBack Mode = "push_back"
	ModeDryRun   Mode = "dry_run"
)

func FixCollisions(mode Mode) error {
	return errors.New("not implemented")
}
