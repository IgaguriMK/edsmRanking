package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const (
	UrlBase = "https://www.edsm.net/en/top/"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		log.Fatal("Need ranking name.")
	}
	rankingName := args[1]

	doc, err := goquery.NewDocument(UrlBase + rankingName)
	if err != nil {
		log.Fatal("goquery error: ", err)
	}

	lineCount := 0

	doc.Find("table.table-hover tr").Each(func(index int, s *goquery.Selection) {
		// Skip <tr> <th> ... line.
		if index == 0 {
			return
		}

		tds := s.Find("td")
		if tds.Length() < 4 {
			log.Printf("too short td [%d]\n", tds.Length())
			return
		}

		rankStr := tds.Eq(1).Text()
		cmdr := tds.Eq(2).Text()
		cntStr := tds.Eq(3).Text()

		rankStr = strings.TrimSpace(rankStr)
		cmdr = strings.TrimSpace(cmdr)
		cntStr = strings.TrimSpace(cntStr)
		cntStr = strings.Replace(cntStr, ",", "", -1)

		fmt.Printf("%s\t%s\t%s\n", rankStr, cmdr, cntStr)
		lineCount++
	})

	if lineCount < 100 {
		log.Fatal("Too few output.")
	}
}
