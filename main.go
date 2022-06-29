package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gocolly/colly"
)

func scraper() {
	filename := "data.csv"
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()
	c := colly.NewCollector(
		colly.AllowedDomains("internshala.com"),
	)
	c.OnHTML(".internship_meta", func(e *colly.HTMLElement) {
		writer.Write([]string{
			e.ChildText("a"),
			e.ChildText("span"),
		})
	})

	//loops 200 pages
	for i := 0; i < 20; i++ {
		fmt.Println("web scaping page", i)
		c.Visit("https://internshala.com/internships/page-" + strconv.Itoa(i))
		break
	}
	log.Printf("it finnaly  works\n")
	log.Panic(c)
}
func main() {
	scraper()
}
