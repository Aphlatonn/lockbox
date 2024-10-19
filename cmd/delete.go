package cmd

import (
	"lockbox/config"
	"strings"

	"bufio"
	"fmt"
	"os"
	"path"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:     "delete",
	Short:   "Delete a passwords store",
	Long:    `Delete a passwords store file with the provided name.`,
	Aliases: []string{"remove"},
	Run: func(cmd *cobra.Command, args []string) {
		// Get the config
		conf := config.GetConfig()

		// check if the passwords store is provided in args
		if len(args) > 0 {
			passwordsStore := strings.TrimSpace(strings.Join(args, "_"))
			if err := deletePasswordsStore(path.Join(conf.PasswordsStorePath, passwordsStore+passwordFileExt)); err != nil {
				fmt.Println(err)
				return
			}
		} else {
			// Create a new Scanner
			scanner := bufio.NewScanner(os.Stdin)

			// Prompt for data store name
			fmt.Print("💾 Enter passwords store name: ")
			scanner.Scan()
			passwordsStoreName := scanner.Text()

			if passwordsStoreName != "" {
				passwordsStoreName = strings.Replace(strings.TrimSpace(passwordsStoreName), " ", "_", 0)
				if err := deletePasswordsStore(path.Join(conf.PasswordsStorePath, passwordsStoreName+passwordFileExt)); err != nil {
					fmt.Println(err)
					return
				}
			}
		}
	},
}

func deletePasswordsStore(filePath string) error {
	// Create the database file if it doesn't exist
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return fmt.Errorf("There is no passwords store with this name")
	}

	// Create a new Scanner
	scanner := bufio.NewScanner(os.Stdin)

	// Prompt for data store name
	fmt.Print("🚨 Are you sure you want to delete this passwords store (yes/no): ")
	scanner.Scan()
	confirmation := scanner.Text()

	// check if the user confirmed
	if confirmation == "yes" {
		err := os.Remove(filePath)
		if err != nil {
			return err
		}
		fmt.Println("✅ Password store deleted successfully")
	}
	return nil
}
