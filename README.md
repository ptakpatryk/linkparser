# linkparser
## HTML Parser for anchor tags

Pass path to the `*.html` file using `-html` flag.
Alow `linkparser` to do the magic ✨

Returns slice of `Link` structs:
```go
 type Link struct {
	Href string
	Text string
}
```
