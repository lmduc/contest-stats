package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// PostgresConfig holds the configurations to Postgres
type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	Ssl      string
}

// Info returns the Postgres configuration in string
func (pc *PostgresConfig) Info() string {
	res := ""

	if len(pc.Host) > 0 {
		res += fmt.Sprintf(" host=%s", pc.Host)
	}

	if len(pc.Port) > 0 {
		res += fmt.Sprintf(" port=%s", pc.Port)
	}

	if len(pc.User) > 0 {
		res += fmt.Sprintf(" user=%s", pc.User)
	}

	if len(pc.Password) > 0 {
		res += fmt.Sprintf(" password=%s", pc.Password)
	}

	if len(pc.Name) > 0 {
		res += fmt.Sprintf(" dbname=%s", pc.Name)
	}

	if pc.Ssl == "false" {
		res += " sslmode=disable"
	}

	return res[1:]
}

type postgresDatabase struct {
	cfg *PostgresConfig
	db  *gorm.DB
}

// Connect to the Postgres
func (pg *postgresDatabase) Connect() error {
	db, err := gorm.Open("postgres", pg.cfg.Info())
	if err != nil {
		panic(err)
	}
	pg.db = db
	return nil
}

// Close the connection to the Postgres
func (pg *postgresDatabase) Close() error {
	return pg.db.Close()
}

// Connection returns the connection to the Postgres
func (pg *postgresDatabase) Connection() *gorm.DB {
	return pg.db
}

// NewPostgresDatabase returns a database connecting to Postgres
func NewPostgresDatabase(cfg *PostgresConfig) Database {
	return &postgresDatabase{
		cfg: cfg,
	}
}
