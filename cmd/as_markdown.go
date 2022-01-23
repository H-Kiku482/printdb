package cmd

// print as markdown

import (
	"errors"
	"os"
)

func (pt *PrintTables) PrintAsMarkdown(filepath string) error {
	tables, err := pt.getTableRows()
	if err != nil {
		return err
	}

	f, err := os.Create(filepath)
	if err != nil {
		return errors.New("failed to open stream")
	}
	defer f.Close()

	var tableName string

	f.WriteString("# " + pt.database + "\n")

	for tables.Next() {
		if err := tables.Scan(&tableName); err != nil {
			return errors.New("failed to get table")
		}

		getTableSQL := "SHOW COLUMNS FROM `" + tableName + "`;"
		tableColumns, err := pt.db.Query(getTableSQL)
		if err != nil {
			return errors.New("failed to get table columns")
		}

		f.WriteString("\n")
		f.WriteString("## " + tableName + "\n")
		f.WriteString("\n")
		f.WriteString("| Field | Type | Null | Key | Default | Extra |\n")
		f.WriteString("| :-- | :-- | :-- | :-- | :-- | :-- |\n")

		var tableColumnInfo tableColumnInfo
		for tableColumns.Next() {
			tableColumns.Columns()
			err := tableColumns.Scan(&tableColumnInfo.field, &tableColumnInfo.dataType, &tableColumnInfo.null, &tableColumnInfo.key, &tableColumnInfo.defaultValue, &tableColumnInfo.extra)
			if err != nil {
				return errors.New("")
			}
			if tableColumnInfo.defaultValue.String == "" {
				tableColumnInfo.defaultValue.String = "NULL"
			}
			f.WriteString("| " + tableColumnInfo.field.String + " | " + tableColumnInfo.dataType.String + " | " + tableColumnInfo.null.String + " | " + tableColumnInfo.key.String + " | " + tableColumnInfo.defaultValue.String + " | " + tableColumnInfo.extra.String + " |\n")
		}
	}

	return nil
}
