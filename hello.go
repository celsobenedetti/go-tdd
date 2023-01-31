package main

import (
	"fmt"
)

var helloPrefixes = map[string]string{
	"en": "Hello",
	"sp": "Hola",
	"fr": "Bonjour",
}

func Hello(params ...string) string {
	var name, lang string

	if len(params) > 0 && len(params[0]) > 0 {
		name = params[0]
	} else {
		name = "World"
	}

	if len(params) > 1 {
		lang = params[1]
	} else {
		lang = "en"
	}

	return helloPrefixes[lang] + ", " + name + "!"
}

func main() {
	fmt.Println(Hello("World"))
}
