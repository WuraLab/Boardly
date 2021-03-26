# Development



### Run Locally

To run Boardly, clone the repo and run 

- Start the database via docker-compose

```
docker-compose up -d
```

- Run boardly 

````
go run src/backend/cmd/boardly.go
````



### Testing Locally

- Start the test database via docker-compose

````
Start the database via docker-compose
````

- Make sure `.test.env` exists with the right env vars.
- run test

```
TEST_DIR=$(pwd) go test ./...
```





