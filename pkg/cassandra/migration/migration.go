package migration

import (
	"errors"
	"github.com/golang-migrate/migrate"
)

type Migration struct {
	Logger    Logger
	Migration *migrate.Migrate
}

func (m Migration) Up(steps int) error {
	defer m.Migration.Close()
	if steps < 0 {
		return errors.New("steps must be positive value")
	}
	var err error
	if steps == 0 {
		err = m.Migration.Up()
	} else {
		err = m.Migration.Steps(steps)
	}
	return err
}

func (m Migration) Down(steps int) error {
	defer m.Migration.Close()
	if steps < 0 {
		return errors.New("steps must be positive value")
	}
	var err error
	if steps == 0 {
		err = m.Migration.Steps(-1)
	} else {
		err = m.Migration.Steps(steps * -1)
	}
	return err
}
