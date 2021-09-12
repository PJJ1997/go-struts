package main

import (
	"fmt"
	"net/http"

	"github.com/fatih/structs"
)

type Server struct {
	Name        string `json:"name,omitempty"`
	ID          int
	Enabled     bool
	users       []string // not exported
	http.Server          // embedded
}

func main() {
	server := &Server{
		Name:    "gopher",
		ID:      123456,
		Enabled: true,
	}

	// ========================== Struct Func ==========================
	// Convert a struct to a map[string]interface{}
	// => {"Name":"gopher", "ID":123456, "Enabled":true}
	fmt.Println("========================== Struct Func ==========================")
	m := structs.Map(server)
	fmt.Println(m)

	// Convert the values of a struct to a []interface{}
	// => ["gopher", 123456, true]
	v := structs.Values(server)
	fmt.Println(v)

	// Convert the names of a struct to a []string
	// (see "Names methods" for more info about fields)
	n := structs.Names(server)
	fmt.Println(n)

	// Return the struct name => "Server"
	s := structs.Name(server)
	fmt.Println(s)

	// Check if any field of a struct is initialized or not.
	h := structs.HasZero(server)
	fmt.Println(h)

	// Check if all fields of a struct is initialized or not.
	z := structs.IsZero(server)
	fmt.Println(z)

	// Check if server is a struct or a pointer to struct
	i := structs.IsStruct(server)
	fmt.Println(i)

	// ========================== Field Func ==========================
	t := structs.New(server)

	// Get the Field struct for the "Name" field
	fmt.Println("========================== Field Func ==========================")
	name := t.Field("Name")

	// Get the underlying value,  value => "gopher"
	value := name.Value().(string)
	fmt.Println(value)

	// Set the field's value
	name.Set("another gopher")
	fmt.Println(name.Value().(string))

	// Get the field's kind, kind =>  "string"
	fmt.Println(name.Kind())

	// Get the Field's tag value for tag name "json", tag value => "name,omitempty"
	tagValue := name.Tag("json")
	fmt.Println(tagValue)

	// Nested structs are supported too
	addrField := t.Field("Server").Field("MaxHeaderBytes")

	// Get the value for addr
	a := addrField.Value().(int)
	fmt.Println(a)

	// Or get all fields
	for _, f := range t.Fields() {
		fmt.Printf("field name: %+v\n", f.Name())

		if f.IsExported() {
			fmt.Printf("value   : %+v\n", f.Value())
			fmt.Printf("is zero : %+v\n", f.IsZero())
		}
	}
}
