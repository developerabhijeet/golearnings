package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)
func main() {
	fmt.Println("Web Verb")
	PerformGetRequest()
}

func PerformGetRequest() {
	const myUrl = "http://localhost:8000/get"

	response, err := http.Get(myUrl)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status code: ", response.Status)
	fmt.Println("Content Length:", response.ContentLength)

	var responseString strings.Builder
	content, _ := io.ReadAll(response.Body)
	// fmt.Println("Content:", string(content))
	byteCount, _ := responseString.Write(content)
	fmt.Println("ByteCount is:", byteCount)
	fmt.Println("response String", responseString.String())
}