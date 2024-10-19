package database

const createTables = `
-- create secrets table
CREATE TABLE IF NOT EXISTS secrets (
	Key					TEXT PRIMARY KEY,
	Value				TEXT NOT NULL
);

-- Create login table
CREATE TABLE IF NOT EXISTS login (
  ID					INTEGER PRIMARY KEY AUTOINCREMENT,
  Name					TEXT NOT NULL,
  Folder				TEXT NOT NULL,
  Username				TEXT NOT NULL,
  Password				TEXT NOT NULL,
  Authentication_key	TEXT NOT NULL,
  Url					TEXT NOT NULL,
  Note					TEXT NOT NULL
);
`

func Migrate() error {
	_, err := DB.Exec(createTables)
	return err
}
