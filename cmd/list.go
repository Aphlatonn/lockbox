package cmd

import (
	"fmt"
	"lockbox/config"
	"path"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

const passwordFileExt = ".lb.db"

var listCmd = &cobra.Command{
	Use:     "list",
	Short:   "Get a list of all password stores",
	Aliases: []string{"ls", "all"},
	Run: func(cmd *cobra.Command, args []string) {
		// Get the config
		conf := config.GetConfig()

		// Fetch all password store files
		passwordStores, err := fetchPasswordStores(conf.PasswordsStorePath)
		if err != nil {
			fmt.Printf("Error fetching password stores: %v\n", err)
			return
		}

		// Print the list of password stores
		printPasswordStores(passwordStores)
	},
}

// fetchPasswordStores retrieves all files ending with the password file extension
func fetchPasswordStores(storePath string) ([]string, error) {
	// Get a list of files with the .lb.db extension
	files, err := filepath.Glob(path.Join(storePath, "*"+passwordFileExt))
	if err != nil {
		return nil, err
	}

	// Trim the file extension and return only the file names
	for i, file := range files {
		files[i] = formatFileName(file)
	}
	return files, nil
}

// formatFileName trims the file extension and extracts the base file name
func formatFileName(filePath string) string {
	fileName := strings.TrimSuffix(filepath.Base(filePath), passwordFileExt)
	return "> " + fileName
}

// printPasswordStores prints the list of password stores in a formatted way
func printPasswordStores(stores []string) {
	if len(stores) == 0 {
		fmt.Println("No password stores found.")
	} else {
		fmt.Println("🔑 Password stores:\n", strings.Join(stores, "\n"))
	}
}
