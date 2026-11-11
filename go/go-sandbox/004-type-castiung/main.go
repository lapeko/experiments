package main

import "fmt"

type myInt int
type myFloat float32

func main() {
	var a any = myInt(3)
	switch v := a.(type) {
	case myFloat:
		{
			fmt.Println("myFloat", v)
		}
	case myInt:
		{
			fmt.Println("myInt", v)
		}
	default:
		{
			fmt.Println("Unknown type")
		}
	}
}
