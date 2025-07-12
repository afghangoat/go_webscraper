package main

import "fmt"
import "html"
import (
    "io"
    "log"
    "net/http"
)

func get_page_content(link string)(string) {
    res, err := http.Get(link)
    if err != nil {
        log.Fatal(err)
    }
    content, err := io.ReadAll(res.Body)
    res.Body.Close()
    if err != nil {
        log.Fatal(err)
    }
    return string(content)
}

func main() {
	fmt.Printf("This is a very basic webscraper which gets the content of a specific website link.\n")
	
	var i string
	//"https://afghangoat.hu/index.html"
	
	fmt.Printf("Enter the domain you want to scrape:\n")
	fmt.Scan(&i)
	var content1 string=get_page_content(i)
	
	doc :=html.EscapeString(content1)
	fmt.Println(doc)
	
	//links := make([]string)
	//links = appends(links, link)
	//fmt.Println("%s",content1)
}
