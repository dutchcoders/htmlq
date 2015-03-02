package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) != 0 {
		fmt.Println("htmlq: Query HTML from StdIn using css selectors")
		fmt.Println("\nUsage:")
		fmt.Println("curl https://www.google.com/ | htmlq \"a\"")
		return
	}

	doc, err := goquery.NewDocumentFromReader(os.Stdin)
	if err != nil {
		fmt.Println("Error %s", err)
		return
	}

	doc.Find(strings.Join(os.Args[1:], " ")).Each(func(i int, s *goquery.Selection) {
		fmt.Println(s.Text())
	})

}
