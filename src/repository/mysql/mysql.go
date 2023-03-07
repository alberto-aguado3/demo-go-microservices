package mysql

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	config "demo/cmd/configuration"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	sqlconfig "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func NewMysqlDatabase(c *config.SqlConfig) *SqlDatabase {

	cfg := mysql.NewConfig()
	cfg.User = c.User
	cfg.Passwd = c.Password
	cfg.Net = "tcp"
	cfg.Addr = fmt.Sprintf("%s:%s", c.Host, c.Port)
	cfg.DBName = c.DbName
	cfg.ParseTime = true

	url := cfg.FormatDSN()

	db, err := sql.Open("mysql", url)
	if err != nil {
		panic(err)
	}

	return &SqlDatabase{
		Db:      db,
		Timeout: c.Timeout,
	}
}

// type SqlDriver interface {
// }

type SqlDatabase struct {
	Db      *sql.DB
	Timeout time.Duration
}

func MigrateSQL(db SqlDatabase) error {
	driver, err := sqlconfig.WithInstance(db.Db, &sqlconfig.Config{})
	if err != nil {
		return err
	}

	dir, err := os.Getwd()
	fmt.Println(dir, err)
	m, err := migrate.NewWithDatabaseInstance("file://../src/repository/mysql/migrations", "demo", driver)
	if err != nil {
		return err
	}

	err = m.Force(0)
	fmt.Println(err)

	if err = m.Up(); err != nil {
		if err != migrate.ErrNoChange {
			return fmt.Errorf("SQL migrate up: %s\n", err.Error())
		}
	}
	return nil
}
