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

	t, err := template.New("foo").Parse(`
	<fieldset><legend>Advertisement</legend><p>If you like this, you might like
the opensource software <a href=https://webconverger.com/>Web kiosk
software</a> I develop. It's very useful in public and business environments
for ease of deployment and privacy.</p></fieldset>


	<footer><p><a href=https://github.com/kaihendry/natalian/blob/mk/{{ .Mdwn }}>Source</a></p>
	<p><a href="mailto:hendry+blog@iki.fi?subject=natalian.org/{{ .URL }}%20feedback">Email</a></p>
	<p><a href="https://validator.nu/?doc=http://natalian.org/{{ .URL }}">Validate</a></p>

	</footer>
	</article>

<script async src=/stats.js></script>
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
