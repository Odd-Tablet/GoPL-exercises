package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

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

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func search(str []string, strips []strip) {
	if len(str) >= 1 {
		for j := 0; j < len(str); j++ {
			fmt.Printf("Search term %s:\n", str[j])
			for i := 0; i < len(strips); i++ {
				x := strips[i]
				matched, err := regexp.MatchString("(?i)"+str[j], x.Transcript)
				check(err)
				if matched {
					fmt.Printf("Image: %s\n", x.Img)
					fmt.Printf("URL: http://xkcd.com/%s/\n", strconv.Itoa(x.Num))
					fmt.Printf("Transcript: %s\n", x.Transcript)
				}
			}
		}
	}
}
func main() {
	path := "./json.txt"
	dat, err := ioutil.ReadFile(path)
	check(err)
	var all []strip
	err = json.Unmarshal(dat, &all)
	check(err)
	args := os.Args[1:]
	if len(args) >= 1 {
		search(args, all)
	} else {
		fmt.Printf("Enter a search term: ")
		in := bufio.NewReader(os.Stdin)
		line, err := in.ReadString('\n')
		check(err)
		cleanLine := strings.Split(line, "\n")
		words := strings.Split(cleanLine[0], " ")
		search(words, all)
	}
}
