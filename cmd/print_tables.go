package cmd

import (
	"database/sql"
	"errors"

	_ "github.com/go-sql-driver/mysql"
)

type PrintTables struct {
	db       *sql.DB
	user     string
	password string
	host     string
	database string
}

func NewPrintTables(user string, password string, host string, port string, database string) (*PrintTables, error) {
	db, err := sql.Open("mysql", user+":"+password+"@tcp("+host+":"+port+")/"+database)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	pt := new(PrintTables)
	pt.db = db
	pt.user = user
	pt.password = password
	pt.host = host
	pt.database = database
	return pt, nil
}

func (pt *PrintTables) getTableRows() (*sql.Rows, error) {
	tables, err := pt.db.Query("select TABLE_NAME as 'teble' from information_schema.tables where table_schema = '" + pt.database + "';")
	if err != nil {
		return nil, errors.New("failed to connect to database")
	}
	return tables, nil
}
