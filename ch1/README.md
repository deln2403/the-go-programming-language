

## 1.3 Finding Duplicate lines
### a)
`go build dup1.go`
`cat test_abc.txt | ./dup1`
### b) `go run dup2.go test_abc.txt test_123.txt`
### c)`go run dup3.go test_abc.txt test_123.txt`

## 1.4 Animated GIFs
`go run lissajous.go > out.gif`

## 1.5 Fetching a URL
`go run fetch.go  http://www.google.com > google.html`

## 1.6 Fetching URLs Concurrently
`go run fetchall.go https://golang.org http://gopl.io https://godoc.org`
