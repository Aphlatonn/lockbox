package database

type Login struct {
	ID                int
	Name              string
	Folder            string
	UserName          string
	Password          string
	AuthenticationKey string
	Url               string
	Note              string
}
