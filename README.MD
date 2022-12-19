#### Build and Run
````
go build ./...
go run ./cmd/navigation-app/main.go
````
#### Post Request:
```
curl --location --request POST 'http://localhost:8080/visit' \
--header 'Content-Type: application/json' \
--data-raw '{
    "userId": "1",
    "url": "page-1"
}'
```
#### GET Request
```
curl --location --request GET 'http://localhost:8080/info?url=page-1'
```
####     Next Steps: logging