package main

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/eduncan911/podcast"
)

func main() {
	c := NewConfigFromFile("config.yaml")
	now := time.Now().UTC()
	p := podcast.New(c.Title, c.Website, c.Description, &now, &now)
	p.AddSubTitle(c.SubTitle)
	p.AddAuthor(c.Author, c.Email)
	p.IOwner = &podcast.Author{Name: c.Author, Email: c.Email}
	p.AddAtomLink(c.Feed)
	p.AddImage(c.Image)
	p.AddSummary(c.Description)
	for _, c := range c.Categories {
		main := strings.Split(c, "/")[0]
		sub := strings.Split(c, "/")[0:]
		p.AddCategory(main, sub)
	}
	if c.Explicit {
		p.IExplicit = "Yes"
	} else {
		p.IExplicit = "No"
	}
	// TODO
	// read an MP3 and get the size and length.
	for _, e := range c.Episodes {
		d := time.Now().UTC()
		item := podcast.Item{
			Title:       e.Title,
			Link:        c.Website,
			GUID:        e.Link,
			Description: e.Description,
			PubDate:     &d,
		}
		item.AddImage(e.Image)
		item.AddSummary(e.Description)
		item.AddEnclosure(e.Link, podcast.MP3, e.Size)
		item.AddDuration(e.Duration)
		if _, err := p.AddItem(item); err != nil {
			log.Fatal(item.Title, ": error", err.Error())
		}
	}
	if err := p.Encode(os.Stdout); err != nil {
		log.Fatalf("Unable to encode feed: %v", err.Error())
	}
}
