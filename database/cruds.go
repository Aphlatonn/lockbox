package database

// create secret data func
func CreateSecretData(key string, value string) (err error) {
	_, err = DB.Exec("INSERT INTO secrets (Key, Value) VALUES (?, ?)", key, value)
	return
}

// read secret data func
func ReadSecretData(key string) (value string, err error) {
	err = DB.QueryRow("SELECT Value FROM secrets WHERE Key = ?", key).Scan(&value)
	return
}

// update secret data func
func UpdateSecretData(key string, value string) (err error) {
	_, err = DB.Exec("UPDATE secrets SET Value = ? WHERE Key = ?", value, key)
	return
}
