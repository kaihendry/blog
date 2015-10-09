package main

import (
	"html/template"
	"os"
	"sort"

	"github.com/kaihendry/blog"
)

func main() {

	currentYear := "1900"

	funcMap := template.FuncMap{
		"newYear": func(t string) bool {
			if t == currentYear {
				return false
			} else {
				currentYear = t
				return true
			}
		},
	}

	//fmt.Fprintln(os.Stderr, p)

	p := blog.OrderedList()
	posts := blog.Posts{p}
	sort.Sort(sort.Reverse(posts))
	t, err := template.New("foo").Funcs(funcMap).Parse(`<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8" />
<link href="/style.css" rel="stylesheet">
<meta name=viewport content="width=device-width, initial-scale=1">
<link rel="icon" type="image/svg+xml" href="http://hendry.iki.fi/kaihendry.svg">
<meta name="description" content="Kai Hendry's personal blog">
<meta property="og:description" content="Kai Hendry's personal blog">
<meta property="og:image" content="https://graph.facebook.com/664018124/picture?type=large">
<link rel="alternate" type="application/atom+xml" title="Atom feed" href="index.atom">
<title>Kai Hendry's blog</title>
</head>
<body>

<aside>
<p><span id=greet></span> <dfn><abbr title="An Anglo African of the East Coast of South Africa">natalian</abbr>.org</dfn> is a personal blog by <a href="http://hendry.iki.fi/">Kai Hendry</a>, born in Natal, South Africa 1978. <a href=http://dabase.com>dabase.com is a more technical blog</a> <a href="https://twitter.com/kaihendry">@kaihendry</a></p>
</aside>

<nav>
{{ range $i,$e := . }}
{{ if newYear (.PostDate.Format "2006")}}
{{ if gt $i 0 }}</ol>{{end}}
<h1>{{ .PostDate.Format "2006" }}</h1>
<ol class="index">{{ end }}
<li><time datetime="{{ .PostDate.Format "2006-01-02" }}">{{ .PostDate.Format "Jan 2" }}</time>
<a href="{{ .URL }}">{{ .Title }}</a></li>{{end}}
</ol>
</nav>

<footer>
<p><a href=https://github.com/kaihendry/natalian/blob/mk/Makefile>Generated with a Makefile</a> and a piece of <a href=https://github.com/kaihendry/blog>Golang</a></p>
<p><a href="https://validator.nu/?doc=http%3A%2F%2Fnatalian.org%2F">Valid HTML</a> &amp; <a href="https://developers.google.com/speed/pagespeed/insights/?url=http%3A%2F%2Fnatalian.org%2F">fast!</a></p>
</footer>

<script src=/stats.js></script>
</body>
</html>
`)

	if err != nil {
		panic(err)
	}

	err = t.Execute(os.Stdout, p)

	if err != nil {
		panic(err)
	}

}
