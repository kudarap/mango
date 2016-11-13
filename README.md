# ManGo
[![Build Status](https://circleci.com/gh/javinc/mango.svg?style=shield&circle-token=607278cc890cea8c92e97be98eee9b1748c7f75c)](https://circleci.com/gh/javinc/mango) [![Build Status](https://travis-ci.org/javinc/mango.svg?branch=master)](https://travis-ci.org/javinc/mango) [![Go Report Card](https://goreportcard.com/badge/github.com/javinc/mango)](https://goreportcard.com/report/github.com/javinc/mango)

## feature
- JSON RESTful API
- component base structure
- Rethink as default database
- JWT as default authentication

## usage
### setup
- run server
    - start up rethinkdb server
    - `go run main.go`

### consume
- listing
    - `curl -X GET localhost:8000/test`
- detail
    - `curl -X GET localhost:8000/test/1`
- create
    - `curl -X POST -d '{"title":"dragon","description":"some dragon"}' localhost:8000/test`
- update
    - `curl -X PUT -d '{"title":"dragon","description":"some dragon"}' localhost:8000/test/1`
- remove
    - `curl -X DELETE localhost:8000/test/1`

### options
- filters
    - `/test?filter.title=dragon` return row with an id is 1
- sorting
    - `/test?order=title` ascending by default
    - `/test?order=title,asc` ascending order
    - `/test?order=title,desc` descending order
- pagination
    - `/test?slice=0,10` first 10 rows
- fields (NOT SUPPORTED YET)
    - `/test?fields=id`
    - `/test?fields=id,title`
