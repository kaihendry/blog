package main

import (
	"flag"
	"html/template"
	"os"
	"path/filepath"
)

type Post struct {
	URL  string
	Mdwn string
}

var p Post

func main() {

	flag.Parse()
	mdwn := flag.Arg(0)

	extName := filepath.Ext(mdwn)
	bName := mdwn[:len(mdwn)-len(extName)]

	p = Post{URL: bName, Mdwn: mdwn}

	t, err := template.New("foo").Parse(`<footer><p><a href=https://github.com/kaihendry/natalian/blob/master/{{ .Mdwn }}>Source</a>
	&diam;
	<a href="mailto:hendry+blog@iki.fi?subject=natalian.org/{{ .URL }}%20feedback">Email me feedback please!</a></p>
	<p><a href="https://validator.nu/?doc=http://natalian.org/{{ .URL }}">Validate me</a></p>
	<p><a href="https://validator.w3.org/checklink?uri=http://natalian.org/{{ .URL }}&amp;hide_type=all&amp;depth=&amp;check=Check">Check links</a></p>
	</footer>
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
