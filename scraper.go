package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gocolly/colly"
)

type Dictionary map[string]string

func main(){
	//args := os.Args
	url := "https://kevo-1.github.io" //args[1]
	collector := colly.NewCollector()
	
	collector.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})
	
	collector.OnResponse(func(r *colly.Response) {
		fmt.Println("Got a response from", r.Request.URL)
	})
	
	collector.OnError(func(r *colly.Response, err error) {
		fmt.Println("Error occured!", err)
	})

	skills := []string{}

	collector.OnHTML("span.skill-tag", func(h *colly.HTMLElement) {
		skill := h.Text
		println("Fetched skill", skill)

		skills = append(skills, skill)
	})

	collector.OnScraped(func(r *colly.Response) {
		data := map[string][]string{
			"skills": skills,
		}

		jsonData, err := json.MarshalIndent(data, "", "  ")
		if err != nil {
			fmt.Println("JSON marshal error:", err)
			return
		}

		err = os.WriteFile("skills.json", jsonData, 0644)
		if err != nil {
			fmt.Println("File write error:", err)
			return
		}

		fmt.Println("skills.json saved successfully!")
	})

	collector.Visit(url)
}
