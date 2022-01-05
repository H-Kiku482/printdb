package cmd

// help message

func HelpText() string {
	return `Usage:  printdbtables [OPTION]... DATABASE
   or:  printdbtables [OPTION]... DATABASE -m FILE.md -c false
   or:  printdbtables [OPTION]... DATABASE -t FILE.txt

Options:
    --help, -h          Print this message and exit.
    --user, -u          Select MySQL user. Default is root.
    --password, -p      Use password for print tables.
    --host, -h          Select this option if MySQL host is not '127.0.0.1'.
    --port, -P          Sekect this option if MySQL port is not '3306'.
    --markdown, -m      Print as markdown file.
    --text, -t          Print as markdown file.
    --cat, -c           Display on terminal.`
}
