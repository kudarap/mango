# Puto

## objective
- learning GoLang & implement patterns
- JSON RESTful API
- component base structure

## usage
### setup
- run server
    - `go run app/main.app`

### consume
- listing
    - `curl -X GET localhost:8000/test`
- detail
    - `curl -X GET localhost:8000/test/1`
- adding
    - `curl -X POST -d '{"title":"dragon","description":"some description"}' localhost:8000/test`
- removing
    - `curl -X DELETE localhost:8000/test/1`
- updating
    - `curl -X PUT -d '{"title":"dragon","description":"some description"}' localhost:8000/test/1`

### options
- filters
    - `/test?filters.id=1`
- fields
    - `/test?fields=id`
    - `/test?fields=id,title`
- sort
    - `/test?sort.asc=id`
    - `/test?sort.desc=id`
- page
    - `/test?page.limit=10`
    - `/test?page.limit=10&page.offset=10`
