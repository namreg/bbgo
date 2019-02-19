# bbgo 
[![Build Status](https://travis-ci.org/namreg/bbgo.svg?branch=master)](https://travis-ci.org/namreg/bbgo)

[![GitHub](https://img.shields.io/github/license/mashape/apistatus.svg)](https://github.com/namreg/bbgo/blob/master/LICENSE)

BBGO is a fast [bbcode](https://en.wikipedia.org/wiki/BBCode) compiler for Go with supporting custom tags.

## Usage

```go
bbg := bbgo.New()
fmt.Println(bbg.Parse("[b]Hello World[/b]"))

// Output:
// <b>Hello World</b>
```

## Supported BBCode Syntax
```
[tag]basic tag[/tag]
[tag1][tag2]nested tags[/tag2][/tag1]

[tag=value]tag with value[/tag]
[tag arg=value]tag with named argument[/tag]
[tag="quote value"]tag with quoted value[/tag]

[tag=value foo="hello world" bar=baz]multiple tag arguments[/tag]
```

## Default Tags
 * `[b]text[/b]` --> `<b>text</b>` (b, i, u, and s all map the same)
 * `[url]link[/url]` --> `<a href="link">link</a>`
 * `[url=link]text[/url]` --> `<a href="link">text</a>`
 * `[img]link[/img]` --> `<img src="link">`
 * `[img=link]alt[/img]` --> `<img alt="alt" title="alt" src="link">`
 * `[color=red]text[/color]` --> `<span style="color: red;">text</span>`
 * `[quote]text[/quote]` --> `<blockquote><cite>Quote</cite>text</blockquote>`
 * `[quote=Somebody]text[/quote]` --> `<blockquote><cite>Somebody said:</cite>text</blockquote>`
 * `[quote name=Somebody]text[/quote]` --> `<blockquote><cite>Somebody said:</cite>text</blockquote>`
 * `[code][b]anything[/b][/code]` --> `<pre>[b]anything[/b]</pre>`
 * `[list][*] item 1[*] item 2[*] item 3[/list]` --> `<ul><li> item 1</li><li> item 2</li><li> item 3</li></ul>`

## Adding Custom Tags
```go
bbg.RegisterTag("color", bbgo.Processor(func(ctx *context.Context, tag node.Tag, w io.Writer) {
    switch t := tag.(type) {
	case *node.OpeningTag:
        // write here logic for opening tag
    case *node.ClosingTag:
        // write here logic for opening tag
    case *node.SelfClosingTag:
        // write here logic for self-closing tag
	}
}))
```
The default tags can also be modified by calling `bbg.RegisterTag`:
