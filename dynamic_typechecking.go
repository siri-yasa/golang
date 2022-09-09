package main

import (
	"fmt"
	"reflect"
)

type (
	ID     string
	Person struct {
		name string
	}
)

// func main() {
func Println2(x interface{}) {
	fmt.Println(true)
	fmt.Println(2022)
	fmt.Println(2.15)
	fmt.Println(4 + 5i)
	fmt.Println("Hello world")
	fmt.Println(ID("345678"))
	fmt.Println([4]byte{})
	fmt.Println([]byte{})
	fmt.Println(map[string]int{})
	fmt.Println(Person{name: "sireesha"})
	fmt.Println(&Person{name: "sireesha"})
	fmt.Println(make(chan int))

	if v, ok := x.(ID); ok {
		fmt.Printf("x has type ID, which I defined. The value is: %v\n", v)
	} else {
		fmt.Printf("'%T' is not the type I want\n", x)
	}

	switch x.(type) {
	case bool:
		fmt.Print("This is a boolean value: ", x.(bool))
	case int:
		fmt.Print("This is my nice int value: ", x.(int))
	case float64:
		fmt.Print(x.(float64))
	case complex128:
		fmt.Print(x.(complex128))
	case string:
		fmt.Print(x.(string))
	case Person:
		fmt.Print(x.(Person))
	case chan int:
		fmt.Print(x.(chan int))
	default:
		fmt.Print("Unknown type")
	}

	fmt.Print("\n")
}

func main() {
	var x interface{}
	x = 3.14

	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x) // x.(<type>)

	fmt.Printf("x: type = %v, value = %v\n", t, v)
	good := x
	fmt.Printf("good: type = %T, value = %v\n", good, good)

	x = &struct{ name string }{}

	t = reflect.TypeOf(x)
	v = reflect.ValueOf(x) // x.(<type>)
	fmt.Printf("x: type = %v, value = %v\n", t, v)
	super := x
	fmt.Printf("super: type = %T, value = %v\n", super, super)
}
