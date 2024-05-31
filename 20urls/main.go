package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghb23456"

func main() {
	fmt.Println("Handling URLS")
	fmt.Println("url: ", myurl)

	// parsing

	result, _ := url.Parse(myurl)

	fmt.Println("result: ", result.Scheme)
	fmt.Println("result: ", result.Host)
	fmt.Println("result: ", result.Path)
	fmt.Println("result: ", result.Port())
	fmt.Println("result: ", result.RawQuery)

	qparams := result.Query()
	fmt.Println("params: ", qparams)

	for _, value := range qparams {
		fmt.Println("value of params", value)
	}

	partsOfUrl :=  &url.URL{
		Scheme: "https",
		Host: "lco.dev",
		Path: "/tutcss",
		RawPath: "user=abhi",
	}

	anotherUrl := partsOfUrl.String()
	fmt.Println("anotherurl:", anotherUrl)

}