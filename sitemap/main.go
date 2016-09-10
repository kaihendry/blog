package main

import (
	"fmt"
	"sort"

	"github.com/kaihendry/blog"
)

func main() {

	posts := blog.OrderedList()
	sort.Sort(sort.Reverse(posts))

	for _, v := range posts {
		fmt.Printf("https://natalian.org%s\n", v.URL)
	}

}
