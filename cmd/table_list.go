package cmd

// return string like 'show columns from TABLE' sql

import (
	"database/sql"
	"errors"
	"unicode/utf8"
)

func (pt *PrintTables) getOneTableListAsString(tableName string, tableColumns *sql.Rows) (*string, error) {
	var returnValue string

	var tableColumnInfo tableColumnInfo

	r := newRecords()

	for tableColumns.Next() {
		tableColumns.Columns()

		if err := tableColumns.Scan(&tableColumnInfo.field, &tableColumnInfo.dataType, &tableColumnInfo.null, &tableColumnInfo.key, &tableColumnInfo.defaultValue, &tableColumnInfo.extra); err != nil {
			return nil, errors.New("failed to connect to database")
		}
		if tableColumnInfo.defaultValue.String == "" {
			tableColumnInfo.defaultValue.String = "NULL"
		}

		r.columnsCount = r.columnsCount + 1

		r.field = append(r.field, tableColumnInfo.field.String)
		r.dataType = append(r.dataType, tableColumnInfo.dataType.String)
		r.null = append(r.null, tableColumnInfo.null.String)
		r.key = append(r.key, tableColumnInfo.key.String)
		r.defaultValue = append(r.defaultValue, tableColumnInfo.defaultValue.String)
		r.extra = append(r.extra, tableColumnInfo.extra.String)

		if c := utf8.RuneCountInString(tableColumnInfo.field.String); r.fieldCount < c {
			r.fieldCount = c
		}
		if c := utf8.RuneCountInString(tableColumnInfo.dataType.String); r.dataTypeCount < c {
			r.dataTypeCount = c
		}
		if c := utf8.RuneCountInString(tableColumnInfo.null.String); r.nullCount < c {
			r.nullCount = c
		}
		if c := utf8.RuneCountInString(tableColumnInfo.key.String); r.keyCount < c {
			r.keyCount = c
		}
		if c := utf8.RuneCountInString(tableColumnInfo.defaultValue.String); r.defaultValueCount < c {
			r.defaultValueCount = c
		}
		if c := utf8.RuneCountInString(tableColumnInfo.extra.String); r.extraCount < c {
			r.extraCount = c
		}
	}

	listBorder := "+-"
	for i := 0; i < r.fieldCount; i++ {
		listBorder += "-"
	}
	listBorder += "-+-"
	for i := 0; i < r.dataTypeCount; i++ {
		listBorder += "-"
	}
	listBorder += "-+-"
	for i := 0; i < r.nullCount; i++ {
		listBorder += "-"
	}
	listBorder += "-+-"
	for i := 0; i < r.keyCount; i++ {
		listBorder += "-"
	}
	listBorder += "-+-"
	for i := 0; i < r.defaultValueCount; i++ {
		listBorder += "-"
	}
	listBorder += "-+-"
	for i := 0; i < r.extraCount; i++ {
		listBorder += "-"
	}
	listBorder += "-+"

	returnValue += listBorder + "\n"

	headerLine := "| "
	headerLine += r.fieldHeader
	for i := 0; i < r.fieldCount-utf8.RuneCountInString(r.fieldHeader); i++ {
		headerLine += " "
	}
	headerLine += " | "
	headerLine += r.dataTypeHeader
	for i := 0; i < r.dataTypeCount-utf8.RuneCountInString(r.dataTypeHeader); i++ {
		headerLine += " "
	}
	headerLine += " | "
	headerLine += r.nullHeader
	for i := 0; i < r.nullCount-utf8.RuneCountInString(r.nullHeader); i++ {
		headerLine += " "
	}
	headerLine += " | "
	headerLine += r.keyHeader
	for i := 0; i < r.keyCount-utf8.RuneCountInString(r.keyHeader); i++ {
		headerLine += " "
	}
	headerLine += " | "
	headerLine += r.defaultValueHeader
	for i := 0; i < r.defaultValueCount-utf8.RuneCountInString(r.defaultValueHeader); i++ {
		headerLine += " "
	}
	headerLine += " | "
	headerLine += r.extraHeader
	for i := 0; i < r.extraCount-utf8.RuneCountInString(r.extraHeader); i++ {
		headerLine += " "
	}
	headerLine += " |"

	returnValue += headerLine + "\n"
	returnValue += listBorder + "\n"

	for i := 0; i < r.columnsCount; i++ {
		line := "| "
		line += r.field[i]
		for j := 0; j < r.fieldCount-utf8.RuneCountInString(r.field[i]); j++ {
			line += " "
		}
		line += " | "
		line += r.dataType[i]
		for j := 0; j < r.dataTypeCount-utf8.RuneCountInString(r.dataType[i]); j++ {
			line += " "
		}
		line += " | "
		line += r.null[i]
		for j := 0; j < r.nullCount-utf8.RuneCountInString(r.null[i]); j++ {
			line += " "
		}
		line += " | "
		line += r.key[i]
		for j := 0; j < r.keyCount-utf8.RuneCountInString(r.key[i]); j++ {
			line += " "
		}
		line += " | "
		line += r.defaultValue[i]
		for j := 0; j < r.defaultValueCount-utf8.RuneCountInString(r.defaultValue[i]); j++ {
			line += " "
		}
		line += " | "
		line += r.extra[i]
		for j := 0; j < r.extraCount-utf8.RuneCountInString(r.extra[i]); j++ {
			line += " "
		}
		line += " |"
		returnValue += line + "\n"
	}

	r.columnsCount = 0

	returnValue += listBorder + "\n"

	return &returnValue, nil
}
