package mysql

import (
	"database/sql"
	"fmt"
	"time"

	config "demo/cmd/configuration"

	"github.com/go-sql-driver/mysql"
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
