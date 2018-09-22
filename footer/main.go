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
	the <a href=https://webconverger.com/>stateless Web kiosk software</a> I
	develop. Webconverger typically replaces Windows on PCs and is deployed in
	public and business environments for ease of deployment and privacy. Once
	installed it auto-updates making it painless to maintain. Try it where you
	exclusively use the only viable open platform... the Web!</p></fieldset>

	<footer><p><a href=https://github.com/kaihendry/natalian/blob/mk/{{ .Mdwn }}>Source</a></p>
	<p><a href="mailto:hendry+blog@iki.fi?subject=natalian.org/{{ .URL }}%20feedback">Email</a></p>
	<p><a href="https://validator.w3.org/nu/?doc=https://natalian.org/{{ .URL }}">Validate</a></p>

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
