# Statistics for Unicode names

Program `stats.go` computes basic statistics about names in the Unicode Character database provided by the [`runenames` package](https://godoc.org/golang.org/x/text/unicode/runenames).


```
$ go run stats.go
unicode.MaxRune =  1114111
runenames.UnicodeVersion =  10.0.0
Repeated character names (with counts):
    33	<control>
  6582	<CJK Ideograph Extension A>
 20971	<CJK Ideograph>
 11172	<Hangul Syllable>
   896	<Non Private Use High Surrogate>
   128	<Private Use High Surrogate>
  1024	<Low Surrogate>
  6400	<Private Use>
  6125	<Tangut Ideograph>
 42711	<CJK Ideograph Extension B>
  4149	<CJK Ideograph Extension C>
   222	<CJK Ideograph Extension D>
  5762	<CJK Ideograph Extension E>
  7473	<CJK Ideograph Extension F>
 65534	<Plane 15 Private Use>
 65534	<Plane 16 Private Use>
____________________________________________________________
 31523	characters with unique names
first:	U+0020	' '	SPACE
 last:	U+E01EF	'󠇯'	VARIATION SELECTOR-256
```
