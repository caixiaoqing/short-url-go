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
```bash
# Test create short url with this command:
curl -sX POST -H 'Content-Type: application/json' 'localhost:8080/shorten' -d '{"url":"http://a.very.long.url"}'
```
-----

2.Get original( long) url from short

> GET    /original/{short_url}
> 
```bash
# Test get original url with this command:
curl -sX GET -H 'Content-Type: application/json' 'localhost:8080/original/short1'
```
-----

3.Redirect to original( long) url from short

> GET    /{short_url}
> 
```bash
# Test redirect to original url with this command:
curl -sX GET -H 'Content-Type: application/json' 'localhost:8080/short1'
```
-----
