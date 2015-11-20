# puto toolkit #

## objective ##
- learning GoLang & implement patterns
- JSON RESTful API
- component base structure

## usage ##
- `go run app/main.app`
- `curl -X GET localhost:8000/test`
- `curl -X GET localhost:8000/test/1`
- `curl -X POST -d '{"title":"dragon","description":"some description"}' localhost:8000/test`
- `curl -X DELETE localhost:8000/test/1`
- `curl -X PUT -d '{"title":"dragon","description":"some description"}' localhost:8000/test/1`

## options ##
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
