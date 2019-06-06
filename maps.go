package main

import "fmt"

func main() {
	m := make(map[string]interface{})
	m["name"] = "Matheus Souza"
	m["age"] = 21
	m["skills"] = []string{"php", "javscript"}
	fmt.Println(m)
}
