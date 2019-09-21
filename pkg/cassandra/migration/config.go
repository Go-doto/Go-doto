package migration

import (
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/cassandra"
	"net/url"
	"strings"
)

type Config struct {
	CassandraAddress string //"cassandra://localhost:9042/keyspacename"
	MigrationsUrl    string //file://./pkg/cassandra/migrations
}

func CreateMigration(c Config) (*Migration, error) {
	cassandraInstance := &cassandra.Cassandra{}

	driver, err := cassandraInstance.Open(c.CassandraAddress)

	if err != nil {
		return nil, err
	}

	u, err := url.Parse(c.CassandraAddress)
	if err != nil {
		return nil, err
	}

	keySpace := strings.TrimPrefix(u.Path, "/")
	m, err := migrate.NewWithDatabaseInstance(
		c.MigrationsUrl,
		keySpace, driver)

	migration := &Migration{
		Logger:    Logger{},
		Migration: m,
	}
	if err != nil {
		return migration, err
	}
	m.Log = &Logger{}
	return migration, nil
}
