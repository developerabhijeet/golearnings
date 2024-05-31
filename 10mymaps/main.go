package main
import "fmt"

func main() {
	fmt.Println("Maps")

	languages := make(map[string]string)

	languages["JS"] = "javascript"
	languages["rb"] = "ruby"
	languages["py"] = "python"

	fmt.Println("list of lang:", languages)
	fmt.Println("JS:", languages["py"])

	delete(languages, "rb")
	fmt.Println("updated lang:", languages)

	// Loops are interesting in Golang

	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}

	// if you don't need key then use _
	for _, value := range languages {
		fmt.Printf("value is %v\n", value)
	}

}