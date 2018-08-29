

## 1.3 Finding Duplicate lines
### a) dup1
`go build dup1.go`
`cat test_abc.txt | ./dup1`
### b) dup2
`go run dup2.go test_abc.txt test_123.txt`
### c) dup3
`go run dup3.go test_abc.txt test_123.txt`

## 1.4 Animated GIFs
`go run lissajous.go > out.gif`

## 1.5 Fetching a URL
`go run fetch.go  http://www.google.com > google.html`

## 1.6 Fetching URLs Concurrently
`go run fetchall.go https://golang.org http://gopl.io https://godoc.org`

## 1.7 A Web Server
### a) server1
`go run server1.go`
`curl localhost:8080/some/test/path`
### b) server2
`go run server1.go`
`curl localhost:8080`
`curl localhost:8080/count`
### c) server3
`go run server3.go`
`curl localhost:8080`
### d) server4
`go run server4.go`
In browser, navigate to http://localhost:8080
