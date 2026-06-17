// Reflection in Go is a powerful feature that allows you to inspect and manipulate the structure of types at runtime. It is part of the reflect package in the standard library. With reflection, you can dynamically access fields, call methods, and even modify values of variables. This can be particularly useful for tasks such as serialization, deserialization, and building generic functions. In this example, we will demonstrate how to use reflection to inspect the fields and methods of a struct type.

package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	p := Person{Name: "Alice", Age: 30}

	v := reflect.ValueOf(p)

	fmt.Printf("Type: %s\n", v.Type())

	for i := 0; i < v.NumField(); i++ {
		fmt.Printf("Field %d: %s = %v\n", i, v.Type().Field(i).Name, v.Field(i).Interface())
	}

	vPtr := reflect.ValueOf(&p).Elem()
	vPtr.FieldByName("Name").SetString("Bob")
	vPtr.FieldByName("Age").SetInt(25)

	fmt.Printf("Updated Person: %+v\n", p)

}
