package cmd

// call with errors

func (pt *PrintTables) CloseDB() error {
	pt.db.Close()
	if err := pt.db.Close(); err != nil {
		return err
	}
	return nil
}
