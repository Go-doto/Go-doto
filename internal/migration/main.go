package main

import (
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/cassandra"
	_ "github.com/golang-migrate/migrate/source/file"
	"log"
	"os"
)

var cassandraAddress string

type Logger struct {
	migrate.Logger
}

func (l Logger) Printf(format string, v ...interface{}) {
	log.Printf(format, v)
}

func (l Logger) Verbose() bool {
	return false
}

func init() {
	if addr := os.Getenv("CASSANDRA_URL"); addr != "" {
		cassandraAddress = addr
	} else {
		cassandraAddress = fmt.Sprintf("cassandra://%v:%v/doto", "localhost", 9042)
	}
}

func getMigrate() (*migrate.Migrate, error) {
	cassandraInstance := &cassandra.Cassandra{}

	driver, err := cassandraInstance.Open(cassandraAddress)

	if err != nil {
		return nil, err
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://./pkg/cassandra/migrations",
		"doto", driver)
	if err != nil {
		return m, err
	}
	m.Log = &Logger{}
	return m, nil
}

func up() {
	m, err := getMigrate()
	if err != nil {
		log.Fatal(err)
	}
	defer m.Close()
	err = m.Up()
	if errors.Is(err, migrate.ErrNoChange) {
		log.Println("no new migrations")
		os.Exit(0)
	}
	if err != nil {
		log.Fatal(err)
	}
	log.Println("finished")
}

func down() {
	m, err := getMigrate()
	if err != nil {
		log.Fatal(err)
	}
	defer m.Close()
	err = m.Steps(-1)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("finished")
}

func main() {
	var commands = map[string]func(){
		"up":   up,
		"down": down,
	}

	args := os.Args
	if len(args) >= 2 {
		if action, ok := commands[args[1]]; ok {
			action()
		}
	}

}
