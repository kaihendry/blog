package main

import (
	"io/ioutil"
	"log"
	"sort"
	"time"

	"github.com/gorilla/feeds"
	"github.com/kaihendry/blog"
)

func main() {

	author := feeds.Author{"Kai Hendry", "hendry+natalian@iki.fi"}
	now := time.Now()
	feed := &feeds.Feed{
		Title:       "Kai Hendry's blog",
		Link:        &feeds.Link{Href: "https://natalian.org/"},
		Description: "Anglo African linux geek living in Singapore",
		Author:      &author,
		Created:     now,
	}

	posts := blog.OrderedList()
	sort.Sort(sort.Reverse(posts))

	maximum := 20
        if (maximum > len(posts)) {
                maximum = len(posts)
        }
		
	for _, v := range posts[:maximum] {
		if v.Description == "" {
			log.Println("Warning:", v.URL, "has no tl;dr")
		}
		i := feeds.Item{
			Title:       v.Title,
			Link:        &feeds.Link{Href: "https://natalian.org" + v.URL},
			Description: v.Description,
			Created:     v.PostDate,
		}
		feed.Add(&i)
	}

	atom, _ := feed.ToAtom()
	rss, _ := feed.ToRss()

	//fmt.Println(atom)
	ioutil.WriteFile("index.atom", []byte(atom), 0644)
	ioutil.WriteFile("index.rss", []byte(rss), 0644)

}
