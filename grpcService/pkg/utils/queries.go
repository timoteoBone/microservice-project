package utils

var (
	CreateUserQuery  string = "INSERT INTO USER (first_name, id, pass, age, email) VALUES (?,?,?,?,?)"
	GetUserQuery     string = "SELECT first_name, age, email FROM USER WHERE id=?"
	GetPasswordQuery string = "SELECT pass FROM USER WHERE id = ?"
)
