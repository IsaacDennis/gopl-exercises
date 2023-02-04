package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Movie struct {
	Title string `json:"title"`
	Year  int    `json:"released"`
	Color bool   `json:"color"`
}

func main() {
	content, _ := ioutil.ReadFile("./test.json")
	fmt.Println(string(content))
	movies := []Movie{}
	err := json.Unmarshal(content, &movies)
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	fmt.Printf("%#v", movies)
}
