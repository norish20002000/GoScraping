package main

import (
    "os"
    "fmt"
    "github.com/PuerkitoBio/goquery"
    "io/ioutil"
)

func main() {
    doc, err := goquery.NewDocument("https://tv.yahoo.co.jp/listings/23/?&st=4&s=1&va=24&vb=1&vc=1&vd=0&ve=0&d=20180815")
    if err != nil {
        fmt.Print("document not found. ")
        os.Exit(1)
    }

    res, err := doc.Find("body").Html()
    if err != nil {
        fmt.Print("dom get failed")
    }
    ioutil.WriteFile("goquery.html", []byte(res), os.ModePerm)

   program := ""
   doc.Find("img").Each(func(_ int, s *goquery.Selection) {
	    str, _ := s.Attr("src")
	program += str + "\n"
   })

   fmt.Print(program)
}

