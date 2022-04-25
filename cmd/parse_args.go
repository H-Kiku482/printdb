package cmd

// parce command line arguments

import (
	"errors"
	"flag"
)

type cmdArgs struct {
	Help     bool
	User     string
	Password bool
	Host     string
	Port     string
	Database string
	Cat      bool
	Text     string
	Markdown string
}

func ParseAndSetCmdArgs(args []string) (*cmdArgs, error) {
	ca := new(cmdArgs)

	user := flag.String("u", "root", "Select MySQL user.")
	password := flag.Bool("p", false, "Use password for printing.")
	host := flag.String("H", "127.0.0.1", "Select MySQL Host.")
	port := flag.String("P", "3306", "Select MySQL port.")
	cat := flag.Bool("c", true, "Print tables on terminal.")
	text := flag.String("t", "", "Print tables on *.txt file.")
	markdown := flag.String("m", "", "Print tables on *.md file.")

	flag.Parse()

	ca.User = *user
	ca.Password = *password
	ca.Host = *host
	ca.Port = *port
	ca.Cat = *cat
	ca.Text = *text
	ca.Markdown = *markdown
	ca.Database = flag.Arg(0)

	if ca.Database == "" {
		e := errors.New("please input the database name")
		return nil, e
	}

	return ca, nil
}
