#!/bin/bash

curl -sX POST -H 'Content-Type: application/json' 'localhost:8080/shorten' -d '{"url":"http://a.very.long.url"}'

curl -sX GET -H 'Content-Type: application/json' 'localhost:8080/original/000001'

curl -sX GET -H 'Content-Type: application/json' 'localhost:8080/000001'

