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

	t, err := template.New("foo").Parse(`<footer><p><a href=https://github.com/kaihendry/natalian/blob/mk/{{ .Mdwn }}>Source</a></p>
	<p><a href="mailto:hendry+blog@iki.fi?subject=natalian.org/{{ .URL }}%20feedback">Email me feedback please!</a></p>
	<p><a href="https://validator.nu/?doc=http://natalian.org/{{ .URL }}">Validate me</a></p>
	</footer>
	</article>

<script>
  (function(i,s,o,g,r,a,m){i['GoogleAnalyticsObject']=r;i[r]=i[r]||function(){
  (i[r].q=i[r].q||[]).push(arguments)},i[r].l=1*new Date();a=s.createElement(o),
  m=s.getElementsByTagName(o)[0];a.async=1;a.src=g;m.parentNode.insertBefore(a,m)
  })(window,document,'script','//www.google-analytics.com/analytics.js','ga');

  ga('create', 'UA-195686-1', 'auto');
  ga('send', 'pageview');

function trackJavaScriptError(e) {
        var ie = window.event || {},
            errMsg = e.message || ie.errorMessage;
        var errSrc = (e.filename || ie.errorUrl) + ': ' + (e.lineno || ie.errorLine);
        ga('send', 'event', 'JavaScript Error', errMsg, errSrc, { 'nonInteraction': 1 });
}

if (window.addEventListener) {
  window.addEventListener('error', trackJavaScriptError, false);
} else if (window.attachEvent) {
  window.attachEvent('onerror', trackJavaScriptError);
} else {
  window.onerror = trackJavaScriptError;
}
</script>
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
