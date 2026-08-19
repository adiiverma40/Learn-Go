package main

import "fmt"

func main() {
	fmt.Println("Maps in goLang")
	m := make(map[string]string)
	m["rb"] = "Ruby"
	m["py"] = "Python"
	m["js"] = "JavaScript"
	m["go"] = "Go"
	fmt.Println(m)
	fmt.Println("go Stands for ", m["go"])
	delete(m, "go")
	fmt.Println(m)

	// Intresting loop in maps
	// 
	for key, value := range m {
		fmt.Println(key, "Stands for ", value)
	}
	
	for _ , value := range m {
		fmt.Println(value)
	}
}