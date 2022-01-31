package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/H-Kiku482/printtables/cmd"
	"golang.org/x/term"
)

func main() {
	args, err := cmd.ParseAndSetCmdArgs(os.Args)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if args.Help {
		fmt.Println(cmd.HelpText())
		os.Exit(0)
	}

	var password = ""

	if args.Password {
		cancelChannel := make(chan os.Signal)
		signal.Notify(cancelChannel, os.Interrupt)
		defer signal.Stop(cancelChannel)

		terminalState, err := term.GetState(int(syscall.Stdin))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		go func() {
			<-cancelChannel
			term.Restore(int(syscall.Stdin), terminalState)
			fmt.Println("")
			os.Exit(1)
		}()

		fmt.Print("Enter password:")
		inputPassword, err := term.ReadPassword(int(os.Stdin.Fd()))
		password = string(inputPassword)
		fmt.Println("")

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	printTables, err := cmd.NewPrintTables(args.User, password, args.Host, args.Port, args.Database)
	defer printTables.CloseDB()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if args.Cat {
		if err := printTables.CatTables(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	if args.Markdown != "" {
		if err := printTables.PrintAsMarkdown(args.Markdown); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	if args.Text != "" {
		if err := printTables.PrintAsText(args.Text); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}
