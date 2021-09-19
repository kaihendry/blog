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
	<fieldset><legend>Found any of my content interesting or useful?</legend>

	<!-- Create a button that your customers click to complete their purchase. -->
	<button id="checkout-button" class="pulse" role="link">Buy me a coffee ☕</button>
	<div id="error-message"></div>
	
	<script>
	  var checkoutButton = document.getElementById('checkout-button');
	  checkoutButton.addEventListener('click', function () {
		// When the customer clicks on the button, redirect
		// them to Checkout.

		var stripe = Stripe('pk_live_ZaFSnYjHJihqC6qadpgHLOgl');

		stripe.redirectToCheckout({
		  items: [{sku: 'sku_EkCsw8Oco9oUMQ', quantity: 1}],
	
		  // Note that it is not guaranteed your customers will be redirected to this
		  // URL *100%* of the time, it's possible that they could e.g. close the
		  // tab between form submission and the redirect.
		  successUrl: window.location.protocol + '//natalian.org/thank-you.html',
		  cancelUrl: window.location.protocol + '//natalian.org/oh-no.html',
		  submitType: 'donate',
		  clientReferenceId: window.location.href
		})
		.then(function (result) {
		  if (result.error) {
			var displayError = document.getElementById('error-message');
			displayError.textContent = result.error.message;
		  }
		});
	  });
	</script>
	
	</fieldset>

	<footer>
	<nav><a href="https://github.com/kaihendry/natalian/blob/main/{{ .Mdwn }}">Source</a> /
	<a href="mailto:hendry+blog@iki.fi?subject=natalian.org/{{ .URL }}%20feedback">Email</a> /
	<a href="https://validator.w3.org/nu/?doc=https://natalian.org/{{ .URL }}">Validate</a></nav>
	</footer>
	</article>

<script async src="https://js.stripe.com/v3"></script>
<script async src="/stats.js"></script>
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
