package main

import (
	"fmt"
)

type Field struct {
	name string
	dest interface{}
}

func Fields(fields ...Field) (string, []interface{}) {
	sqlString := ""
	destinations := make([]interface{},0)
	for _, field := range fields {
		if sqlString != "" {
			sqlString += ", "
		}
		sqlString += field.name
		destinations = append(destinations, field.dest)
		
	}
	return sqlString, destinations
}

func main() {

	data := struct {
		name string
		surname  int
	}{}

	fields, destinations := Fields(
		Field{"NAME", &data.name},
		Field{"SURNAME", &data.surname},
	)

	query := fmt.Sprintf("SELECT %s FROM CATS", fields)
	// ca.: query, dests := sqlm.Format(
	//	"SELECT %r FROM CATS", 
	//	Fields(
	//		Field{"NAME", &data.name},
	//		Field{"SURNAME", &data.surname},
	//	),
	// )
	fmt.Println(query, destinations)
}
