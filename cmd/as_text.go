package cmd

// print as text file

import (
	"errors"
	"os"
)

func (pt *PrintTables) PrintAsText(filepath string) error {
	tables, err := pt.getTableRows()
	if err != nil {
		return err
	}

	f, err := os.Create(filepath)
	if err != nil {
		return errors.New("failed to connect to database")
	}
	defer f.Close()

	f.WriteString(pt.database + "\n")

	for tables.Next() {
		var tableName string
		if err := tables.Scan(&tableName); err != nil {
			return errors.New("failed to open stream")
		}

		getTableSQL := "SHOW COLUMNS FROM `" + tableName + "`;"
		tableColumns, err := pt.db.Query(getTableSQL)
		if err != nil {
			return errors.New("failed to open stream")
		}

		f.WriteString(tableName + "\n")

		tableList, err := pt.getOneTableListAsString(tableName, tableColumns)
		if err != nil {
			return err
		}
		f.WriteString(*tableList)
	}

	return nil
}
