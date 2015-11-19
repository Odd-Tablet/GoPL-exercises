package main

//Web server for XKCD search tool
import (
	//"bufio"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	//"os"
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
func search(str []string, strips []strip, w http.ResponseWriter) {
	for j := 0; j < len(str); j++ {
		//fmt.Printf("Search term %s:\n", str[j])
		for i := 0; i < len(strips); i++ {
			x := strips[i]
			matched, err := regexp.MatchString("(?i)"+str[j], x.Transcript)
			check(err)
			if matched {
				//fmt.Fprintf(w, "<a href = %s>Image: %s\n</a>", x.Img, x.Img)
				fmt.Fprintf(w, "URL: <a href = 'http://xkcd.com/%s'>http://xkcd.com/%s/</a><br/>", strconv.Itoa(x.Num), strconv.Itoa(x.Num))
				fmt.Fprintf(w, "<p>Transcript: %s\n</p>", x.Transcript)
			}
		}
	}
}
func main() {
	http.HandleFunc("/results", resultHandler)
	http.HandleFunc("/", searchHandler)
	log.Fatal(http.ListenAndServe("motmb.rocks:8080", nil))
	/*
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
	*/
}
func searchHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "search")
	//http.Redirect(w, r, "/results", http.StatusFound)
}
func resultHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "results")
	fmt.Fprintf(w, "Search Value: %s<br/>", r.FormValue("body"))
	path := "./json.txt"
	dat, err := ioutil.ReadFile(path)
	check(err)
	var all []strip
	err = json.Unmarshal(dat, &all)
	check(err)
	words := strings.Split(r.FormValue("body"), " ")
	search(words, all, w)
}
func renderTemplate(w http.ResponseWriter, tmpl string) {
	t, err := template.ParseFiles(tmpl + ".html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
