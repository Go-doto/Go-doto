package main

import (
	"fmt"
	"github.com/Go-doto/Go-doto/pkg/cassandra/migration"
	_ "github.com/golang-migrate/migrate/source/file"
	"log"
	"os"
)

func main() {

	var cassandraAddress string
	if addr := os.Getenv("CASSANDRA_URL"); addr != "" {
		cassandraAddress = addr
	} else {
		cassandraAddress = fmt.Sprintf("cassandra://%v:%v/doto", "localhost", 9042)
	}

	config := migration.Config{
		CassandraAddress: cassandraAddress,
		MigrationsUrl:    "file://./pkg/cassandra/migrations",
	}
	migrate, err := migration.CreateMigration(config)

	if err != nil {
		log.Fatal(err)
	}
	args := os.Args
	if len(args) >= 2 {
		if args[1] == "up" {
			err = migrate.Up(0)
			if err != nil {
				log.Fatal(err)
			}
			os.Exit(0)
		}
		if args[1] == "down" {
			err = migrate.Down(0)
			if err != nil {
				log.Fatal(err)
			}
			os.Exit(0)
		}
	}
	fmt.Println("Use: `<command> up` or `<command> down`")
}
