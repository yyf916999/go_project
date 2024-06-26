package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	response, error := http.Get("http://localhost")
	if error != nil {
		fmt.Println(error)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}
}
