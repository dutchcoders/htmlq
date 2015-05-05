package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	kingpin "gopkg.in/alecthomas/kingpin.v1"
)

func main() {
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) != 0 {
		fmt.Println("htmlq: Query HTML from StdIn using css selectors")
		fmt.Println("\nUsage:")
		fmt.Println("curl https://www.google.com/ | htmlq --filter \"a\"")
		return
	}

	filter := kingpin.Flag("filter", "filter").Short('f').Required().String()
	arg := kingpin.Arg("arg", "arg").String()
	kingpin.Parse()

	doc, err := goquery.NewDocumentFromReader(os.Stdin)
	if err != nil {
		fmt.Println("Coud not read input: %s", err.Error())
		return
	}

	doc.Find(*filter).Each(func(i int, s *goquery.Selection) {
		if *arg == "" {
			fmt.Println(strings.TrimSpace(s.Text()))
			return
		}

		// attribute
		if val, ok := s.Attr((*arg)[1:]); ok {
			fmt.Println(val)
		}
	})
}
