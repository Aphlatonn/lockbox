package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "lockbox",
	Short: "A simple and secure password manager.",
	Long:  `Lockbox is a command-line tool that allows you to securely store, retrieve, and manage your passwords.`,
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// add subcommands
	rootCmd.AddCommand(
		newCmd,
		deleteCmd,
		listCmd,
		generateCmd,
	)
}
