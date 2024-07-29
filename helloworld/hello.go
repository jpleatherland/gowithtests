package main

import "fmt"

func Greet(language string) string {
	switch language{
	case "spanish":
		return "Hola, "
	default:
		return "Hello, "
	}
}

func Hello(language, name string) string {
	greeting := Greet(language)
	if name == "" {
		return greeting + "world"
	}
	return greeting + name
}

func main() {
	fmt.Println(Hello("", "Chris"))
}
