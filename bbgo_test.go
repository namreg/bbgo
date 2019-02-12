package bbgo_test

import (
	"testing"

	"github.com/namreg/bbgo"
)

func TestBBGO_Parse(t *testing.T) {

	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"empty string", ``, ``},
		{"check escape", `<script src="'>`, `&lt;script src=&#34;&#39;&gt;`},
		{"not defined tag", `[foo][b]hello[/b][/foo]`, `[foo]<b>hello</b>[/foo]`},
		{"b", `[b][b]hello[/b]`, `<b><b>hello</b>`},
		{"img", `[img]http://example.com/logo.png[/img]`, `<img src="http://example.com/logo.png" />`},
		{"img with title", `[img="bla bla bla"]http://example.com/logo.png[/img]`, `<img title="bla bla bla" src="http://example.com/logo.png" />`},
		{"img without src", `[img][/img]`, `<img src="" />`},
		{"quote", "[quote]hello[/quote]", `<blockquote>hello</blockquote>`},
		{"quote with attr", `[quote name=Someguy]hello[/quote]`, `<blockquote><cite>Someguy said:</cite>hello</blockquote>`},
		{"url", `[url]https://en.wikipedia.org[/url]`, `<a href="https://en.wikipedia.org">https://en.wikipedia.org</a>`},
		{"url with value", `[url=https://en.wikipedia.org]English Wikipedia[/url]`, `<a href="https://en.wikipedia.org">English Wikipedia</a>`},
		{"code", `[code][b]some[/b]\n[i]stuff[/i]\n[/quote][/code][b]more[/b]`, `<pre>[b]some[/b]\n[i]stuff[/i]\n[/quote]</pre><b>more</b>`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := bbgo.New()
			got := b.Parse(tt.input)
			if tt.want != got {
				t.Fatalf("want = %s, got = %s", tt.want, got)
			}
		})
	}
}
