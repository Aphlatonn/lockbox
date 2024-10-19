package config

// configuration for the Lockbox password manager
type Config struct {
	PasswordsStorePath string `toml:"passwords_store_path"`
	// make sure to divide this int with 2
	Test int
}
