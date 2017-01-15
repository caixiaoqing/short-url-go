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

Below is a usage shown from the root of the api url,
![enter image description here](https://lh3.googleusercontent.com/-bV8oWk9tmYs/WHs3WbEtxZI/AAAAAAAAAC8/gNnkRXOqh1okUM1Cv7wVacmzueN1w4y2ACLcB/s0/usage.png "usage.png")

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

## Dependencies
- Golang v1.7+
- External Go packages dependencies

```bash
# production
go get github.com/gorilla/mux           # HTTP router
```

## Cross compile for the target server
To run the web app in the server, you need to know the server's [operating system and architecture](https://golang.org/doc/install/source#environment).
```bash
GOOS=linux GOARCH=amd64 go build github.com/caixiaoqing/short-url-go
```
Note: There is a sample build and script to run and test the short-url-go web service for linux x86 system (I build it as 386, 32-bit x86, it should be compatible with 64-bit x86 also)

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

