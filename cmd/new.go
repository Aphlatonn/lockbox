package cmd

import (
	"lockbox/config"
	"lockbox/database"
	"lockbox/utils"

	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use:     "new",
	Short:   "Create a new password store",
	Long:    `Create a new password store SQLite file with the provided name and master password.`,
	Aliases: []string{"setup"},
	Run: func(cmd *cobra.Command, args []string) {
		var (
			passStoreName string
			masterPass    string
			hint          string
		)

		// what is the passwords store name
		whatPassStoreName := huh.NewInput().
			Title("✍️ What name would you like to use for this passwords store?").
			Prompt("? ").
			Placeholder("my_store").
			EchoMode(huh.EchoModeNormal).
			CharLimit(80).
			Value(&passStoreName)
		if err := whatPassStoreName.Run(); err != nil {
			fmt.Println(err)
			return
		}

		// check if the name is provided
		if passStoreName == "" {
			fmt.Println("⚠️ passwords store name is required")
			return
		}
		passStoreName += ".lb.db"

		// what is the master for taht passwords store
		whatMasterPassword := huh.NewInput().
			Title("🔐 What master password would you like to use?").
			Description("Make sure its a strong and complicated password").
			Prompt("? ").
			Placeholder("my_store").
			EchoMode(huh.EchoModePassword).
			CharLimit(80).
			Value(&masterPass)
		if err := whatMasterPassword.Run(); err != nil {
			fmt.Println(err)
			return
		}

		// check if the master password is provided
		if masterPass == "" {
			fmt.Println("⚠️ master password is required")
			return
		}

		// what is the hint for that passwords store
		whatHint := huh.NewText().
			Title("🧩 What Hint would you like to use? in case you forgot your master password:").
			Description("use 'help hint' to get more info").
			Placeholder("my_store").
			Value(&hint)
		if err := whatHint.Run(); err != nil {
			fmt.Println(err)
			return
		}

		// Get the config
		conf := config.GetConfig()

		// Create the directory if it doesn't exist
		dataStorePath := path.Join(conf.PasswordsStorePath, passStoreName)
		if err := os.MkdirAll(filepath.Dir(dataStorePath), os.ModePerm); err != nil {
			fmt.Println("Error creating directory:", err)
			return
		}

		// Create the database file if it doesn't exist
		if _, err := os.Stat(dataStorePath); os.IsNotExist(err) {
			// Create an empty file to initialize the database
			file, err := os.Create(dataStorePath)
			if err != nil {
				fmt.Println("Error creating database file:", err)
				return
			}
			defer file.Close()
		}

		// Open database connection
		err := database.OpenConnection(dataStorePath)
		if err != nil {
			fmt.Println("Error opening database:", err)
			return
		}
		defer database.CloseConnection()

		// exec the migrations
		if err := database.Migrate(); err != nil {
			fmt.Printf("Error migrating database: %v\n", err)
			return
		}

		// Hash the master password using bcrypt
		hashedMasterPassword, err := utils.HashMasterPassword(masterPass)
		if err != nil {
			fmt.Println("Error hashing password:", err)
			return
		}

		// Save the hashed password
		if err := database.CreateSecretData("master_password", hashedMasterPassword); err != nil {
			fmt.Println("Error saving master password:", err)
			return
		}

		// Save the hint
		if err := database.CreateSecretData("hint", hint); err != nil {
			fmt.Println("Error saving the hint:", err)
			return
		}

		fmt.Printf("✅ Password store created")
	},
}
