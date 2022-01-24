package utils

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/timoteoBone/microservice-project/grpcService/pkg/entities"
)

var (
	CreateUserQuery  string = "INSERT INTO USER (first_name, id, pass, age, email) VALUES (?,?,?,?,?)"
	GetUserQuery     string = "SELECT first_name, age, email FROM USER WHERE id=?"
	GetPasswordQuery string = "SELECT pass FROM USER WHERE id = ?"
	DeleteUserQuery  string = "DELETE FROM USER WHERE id = ?"
	UpdateUserQuery  string = "UPDATE USER SET first_name, pass, age, email WHERE id = ?"
)

func GenerateQuery(user entities.User) (string, []interface{}) {
	referencesToSQL := entities.FieldsReferenceSql

	var query strings.Builder
	query.WriteString("UPDATE USER SET")

	var args []interface{}

	v := reflect.ValueOf(user)
	typeOfS := v.Type()

	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).Interface() != "" {
			query.WriteString(referencesToSQL[typeOfS.Field(i).Name])
			args = append(args, v.Field(i))
		}
		fmt.Printf("Field: %s\tValue: %v\n", typeOfS.Field(i).Name, v.Field(i).Interface())
	}

	query.WriteString("WHERE id = ?")
	return query.String(), args

}
