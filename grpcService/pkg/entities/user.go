package entities

type User struct {
	Name  string
	Pass  string
	Age   uint32
	Email string
}

var (
	FieldsReferenceSql = map[string]string{
		"Name":  "first_name",
		"Pass":  "pass",
		"Age":   "age",
		"Email": "email",
	}
)
