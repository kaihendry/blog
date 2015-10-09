package blog

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

type Post struct {
	PostDate    time.Time
	URL         string
	Title       string
	Description string
}

type Posts struct {
	Posts []Post
}

func (p Posts) Len() int {
	return len(p.Posts)
}

func (p Posts) Less(i, j int) bool {
	return p.Posts[i].PostDate.Before(p.Posts[j].PostDate)
}

func (p Posts) Swap(i, j int) {
	p.Posts[i], p.Posts[j] = p.Posts[j], p.Posts[i]
}

var p []Post

func visit(mdwn string, f os.FileInfo, err error) error {
	if !f.IsDir() {
		// fmt.Printf("Visiting: %s\n", mdwn)

		if filepath.Ext(mdwn) == ".mdwn" {

			fName := filepath.Base(mdwn)
			extName := filepath.Ext(mdwn)
			bName := fName[:len(fName)-len(extName)]
			url := fmt.Sprintf("/%s/", path.Join(filepath.Dir(mdwn), bName))

			desc := GetKey(mdwn, "description")["description"]
			m := GetKey(mdwn, "title")

			title := m["title"]
			if title == "" {
				title = strings.Replace(bName, "_", " ", -1)
			}

			var a, b, c int
			fmt.Sscanf(mdwn, "%d/%d/%d/", &a, &b, &c)
			date := fmt.Sprintf("%d-%02d-%02d", a, b, c)
			t, err := time.Parse("2006-01-02", date)
			//fmt.Println("Date:", t)
			if err != nil {
				panic(err)
			}

			//fmt.Println("Title:", title)
			//fmt.Println("URL:", url)

			p = append(p, Post{PostDate: t, URL: url, Title: title, Description: desc})

			//} else {
			//	fmt.Fprintln(os.Stderr, "Skipping", mdwn)
		}

	}
	return nil
}

func OrderedList() []Post {

	err := filepath.Walk(".", visit)
	if err != nil {
		panic(err)
	}
	return p

}
