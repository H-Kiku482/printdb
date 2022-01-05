package cmd

import "database/sql"

type tableColumnInfo struct {
	field        sql.NullString
	dataType     sql.NullString
	null         sql.NullString
	key          sql.NullString
	defaultValue sql.NullString
	extra        sql.NullString
}
