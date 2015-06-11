package main

import (
	"flag"
	"html/template"
	"os"
	"path/filepath"
	"strings"

	"github.com/kaihendry/blog"
)

type Post struct {
	Title string
	URL   string
}

var p Post

func main() {

	flag.Parse()
	mdwn := flag.Arg(0)

	fName := filepath.Base(mdwn)
	extName := filepath.Ext(mdwn)
	bName := fName[:len(fName)-len(extName)]
	url := mdwn[:len(mdwn)-len(extName)]

	m := blog.GetKey(mdwn, "title")

	title := m["title"]

	if title == "" {
		title = strings.Replace(bName, "_", " ", -1)
	}

	p = Post{Title: title, URL: url}

	t, err := template.New("foo").Parse(`<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8" />
<link href="/style.css" rel="stylesheet">
<meta name=viewport content="width=device-width, initial-scale=1">
<link href="http://natalian.org/{{ .URL }}/" rel=canonical>
<title>{{ .Title }}</title>
</head>
<body>
<h1><a href=/>natalian.org</a><a class=posturl href="/{{ .URL }}/">/<br>{{ .Title }}</a></h1>
`)

	if err != nil {
		panic(err)
	}

	err = t.Execute(os.Stdout, p)

	if err != nil {
		panic(err)
	}

}
