package main

import (
	"fmt"
	"io"
	"net/http"
)
const url = "https://lco.dev"

func main(){
	fmt.Println("Web request")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	fmt.Println("Response :", response)

	defer response.Body.Close() // Caller's responsibility to close the connection

	// io.ReadAll(response.Body)
	databytes, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	content := string(databytes)
	fmt.Println("content:", content)


}