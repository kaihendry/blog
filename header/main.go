package main

import (
	"flag"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/kaihendry/blog"
)

type Post struct {
	Title string
}

var p Post

func main() {

	flag.Parse()
	mdwn := flag.Arg(0)

	fName := filepath.Base(mdwn)
	extName := filepath.Ext(mdwn)
	bName := fName[:len(fName)-len(extName)]

	m := blog.GetKey(mdwn, "title")

	title := m["title"]

	if title == "" {
		title = strings.Replace(bName, "_", " ", -1)
	}

	p = Post{Title: title}

	t, err := template.New("foo").Parse(`<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8" />
<link href="/style.css" rel="stylesheet">
<meta name=viewport content="width=device-width, initial-scale=1">
<title>{{ .Title }}</title>
</head>
<body>
`)

	if err != nil {
		panic(err)
	}

	err = t.Execute(os.Stdout, p)

	if err != nil {
		panic(err)
	}

}
