package cmd

import "unicode/utf8"

type records struct {
	fieldHeader        string
	dataTypeHeader     string
	nullHeader         string
	keyHeader          string
	defaultValueHeader string
	extraHeader        string

	fieldCount        int
	dataTypeCount     int
	nullCount         int
	keyCount          int
	defaultValueCount int
	extraCount        int

	field        []string
	dataType     []string
	null         []string
	key          []string
	defaultValue []string
	extra        []string

	columnsCount int
}

func newRecords() *records {
	var r *records = new(records)

	r.fieldHeader = "Field"
	r.dataTypeHeader = "Type"
	r.nullHeader = "Null"
	r.keyHeader = "Key"
	r.defaultValueHeader = "Default"
	r.extraHeader = "Extra"

	r.fieldCount = utf8.RuneCountInString(r.fieldHeader)
	r.dataTypeCount = utf8.RuneCountInString(r.dataTypeHeader)
	r.nullCount = utf8.RuneCountInString(r.nullHeader)
	r.keyCount = utf8.RuneCountInString(r.keyHeader)
	r.defaultValueCount = utf8.RuneCountInString(r.defaultValueHeader)
	r.extraCount = utf8.RuneCountInString(r.extraHeader)

	r.columnsCount = 0

	return r
}
