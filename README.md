A REST-ful API server which provides url shortener service written in Go (golang).

-----

## Features
1.Create short url from a long url 

> POST    /shorten
> Request body -  {"url":"http://a.very.long.url"}
> 
> Response : 
{
  "short": "http://localhost:8080/original/short1",
  "redirect": "http://localhost:8080/short1"
}

-----

2.Get original( long) url from short

> GET    /original/{short_url}
> 

-----

3.Redirect to original( long) url from short

> GET    /{short_url}
> 

-----

## Build and run the server
To run an instance of a server example:
```bash
#!/bin/bash

cd $GOPATH

# get go-package and put it in your go-workspace
go get github.com/caixiaoqing/short-url-go

# go to package root folder
cd $GOPATH/src/github.com/caixiaoqing/short-url-go

# install dependencies
go get

# run test
go test api_test/api_test.go
# output e.g.
# ok  	command-line-arguments	0.011s

# run Server
go run main.go
```
-----

## Test
After server started, you can test it with below commands:
```bash
# Test create short url with this command:
curl -sX POST -H 'Content-Type: application/json' 'localhost:8080/shorten' -d '{"url":"http://a.very.long.url"}'
# output e.g.
# {"short":"http://localhost:8080/original/000001","redirect":"http://localhost:8080/000001"}
```
```bash
# Test get original url with this command:
curl -sX GET -H 'Content-Type: application/json' 'localhost:8080/original/000001'
# output e.g.
# {"original":"http://a.very.long.url"}
```
```bash
# Test redirect to original url with this command:
curl -sX GET -H 'Content-Type: application/json' 'localhost:8080/000001'
# output e.g.
# <a href="http://a.very.long.url">Temporary Redirect</a>.
```
-----

## Dependencies
- Golang v1.7+
- External Go packages dependencies

```bash
# production
go get github.com/gorilla/mux           # HTTP router
```

----
