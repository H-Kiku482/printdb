package cmd

// parce command line arguments

import (
	"errors"
	"unicode/utf8"
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

func (ca *cmdArgs) checkOptionalValue(args []string, i int) (*string, error) {
	var value string

	if i < len(args)-1 {
		value = args[i+1]
		if (value[0:1] == "\"") && (value[:utf8.RuneCountInString(value)] == "\"") {
			r := value[1 : utf8.RuneCountInString(value)-1]
			return &r, nil
		} else if (value[0:1] == "'") && (value[:utf8.RuneCountInString(value)] == "'") {
			r := value[1 : utf8.RuneCountInString(value)-1]
			return &r, nil
		} else if value[0:1] == "-" {
			return nil, errors.New("invalid variable \"" + value + "\"")
		} else {
			return &value, nil
		}
	} else {
		return nil, errors.New("invalid variable \"" + value + "\"")
	}
}

func ParseAndSetCmdArgs(args []string) (*cmdArgs, error) {
	argc := len(args) - 1
	argv := args[1:]

	ca := new(cmdArgs)

	ca.Help = false
	ca.User = "root"
	ca.Password = false
	ca.Host = "127.0.0.1"
	ca.Port = "3306"
	ca.Database = ""
	ca.Cat = true
	ca.Text = ""
	ca.Markdown = ""

	for i := 0; i < argc; i++ {
		if argv[i][0:1] == "-" {
			if argv[i][1:] == "h" || argv[i][1:] == "-help" {
				ca.Help = true
				return ca, nil
			} else if argv[i][1:] == "u" || argv[i][1:] == "-user" {
				str, err := ca.checkOptionalValue(argv, i)
				if err != nil {
					return ca, err
				}
				i++
				ca.User = *str
			} else if argv[i][1:] == "p" || argv[i][1:] == "-password" {
				str, _ := ca.checkOptionalValue(argv, i)
				if str == nil {
				} else if *str == "true" {
					i++
					ca.Password = true
				} else if *str == "false" {
					i++
					ca.Password = false
				} else {
					ca.Password = true
				}
			} else if argv[i][1:] == "H" || argv[i][1:] == "-host" {
				str, err := ca.checkOptionalValue(argv, i)
				if err != nil {
					return ca, err
				}
				i++
				ca.Host = *str
			} else if argv[i][1:] == "P" || argv[i][1:] == "-port" {
				str, err := ca.checkOptionalValue(argv, i)
				if err != nil {
					return ca, err
				}
				i++
				ca.Port = *str
			} else if argv[i][1:] == "c" || argv[i][1:] == "-cat" {
				str, _ := ca.checkOptionalValue(argv, i)
				if str == nil {
					i++
					ca.Cat = true
				} else if *str == "true" {
					i++
					ca.Cat = true
				} else if *str == "false" {
					i++
					ca.Cat = false
				} else {
					return ca, errors.New("invalid variable \"" + *str + "\"")
				}
			} else if argv[i][1:] == "t" || argv[i][1:] == "-text" {
				str, err := ca.checkOptionalValue(argv, i)
				if err != nil {
					return ca, err
				}
				i++
				ca.Text = *str
			} else if argv[i][1:] == "m" || argv[i][1:] == "-markdown" {
				str, err := ca.checkOptionalValue(argv, i)
				if err != nil {
					return ca, err
				}
				i++
				ca.Markdown = *str
			} else {
				return ca, errors.New("unknown option \"" + argv[i] + "\"")
			}
		} else {
			if ca.Database != "" {
				return ca, errors.New("too many arguments: " + ca.Database + ", " + argv[i])
			}
			if (argv[i][0:1] == "\"") && (argv[i][:utf8.RuneCountInString(argv[i])] == "\"") {
				ca.Database = argv[i][1 : utf8.RuneCountInString(argv[i])-1]
			} else if (argv[i][0:1] == "'") && (argv[i][:utf8.RuneCountInString(argv[i])] == "'") {
				ca.Database = argv[i][1 : utf8.RuneCountInString(argv[i])-1]
			} else {
				ca.Database = argv[i]
			}
		}
	}

	if ca.Database == "" {
		return ca, errors.New("pleace select the database name")
	}

	return ca, nil
}
