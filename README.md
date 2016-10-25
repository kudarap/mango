# ManGo

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
    - `/test?filter.id=1` return row with an id is 1
- sorting
    - `/test?order=id` ascending by default
    - `/test?order=id,asc` ascending order
    - `/test?order=id,desc` descending order
- pagination
    - `/test?slice=0,10` first 10 rows
- fields (NOT SUPPORTED YET)
    - `/test?fields=id`
    - `/test?fields=id,title`