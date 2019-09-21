package migration

import (
	"github.com/golang-migrate/migrate"
	"log"
)

type Logger struct {
	VerboseLog bool
	migrate.Logger
}

func (l Logger) Printf(format string, v ...interface{}) {
	log.Printf(format, v)
}

func (l Logger) Verbose() bool {
	return l.VerboseLog
}
