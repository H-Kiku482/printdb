package cmd

// print on terminal

import (
	"errors"
	"fmt"
)

func (pt *PrintTables) CatTables() error {
	tables, err := pt.getTableRows()
	if err != nil {
		return err
	}

	for tables.Next() {
		var output string
		var tableName string
		if err := tables.Scan(&tableName); err != nil {
			return errors.New("failed to open stream")
		}

		getTableSQL := "SHOW COLUMNS FROM `" + tableName + "`;"
		tableColumns, err := pt.db.Query(getTableSQL)
		if err != nil {
			return errors.New("failed to open stream")
		}

		output += tableName + "\n"

		tableList, err := pt.getOneTableListAsString(tableName, tableColumns)
		if err != nil {
			return err
		}
		output += *tableList

		fmt.Println(output)
	}

	return nil
}
