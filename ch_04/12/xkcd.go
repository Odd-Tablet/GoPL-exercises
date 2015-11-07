package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	path := "./json.txt"
	dat, err := ioutil.ReadFile(path)
	check(err)
	type strip struct {
		Month      string
		Num        int
		News       string
		Link       string
		Year       string
		Safe_title string
		Transcript string
		Alt        string
		Img        string
		Title      string
		Day        string
	}
	var all []strip
	if err := json.Unmarshal(dat, &all); err != nil {
		log.Fatalf("JSON unmarshaling failed: %s", err)
	}
	fmt.Printf("%+v\n", all)
}
