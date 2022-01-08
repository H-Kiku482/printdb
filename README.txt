Usage:  printtables DATABASE
   or:  printtables -u root -p DATABASE -m FILE.md
   or:  printtables DATABASE -t FILE.txt

Options:
    --help, -h          Print this message and exit.
    --user, -u          Select MySQL user. Default is root.
    --password, -p      Use password for printing.
    --host, -H          Select this option when the MySQL host is not 'localhost'.
    --port, -P          Select this option when the MySQL port is not '3306'.
    --markdown, -m      Print as markdown file.
    --text, -t          Print as text file.
    --cat, -c           Display on terminal.
