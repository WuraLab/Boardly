# Development



### Run Locally

To run Boardly, clone the repo and run 

- Start the database, api and web.

```
docker-compose up -d
```

### Testing Locally

- Start the test database via docker-compose

````
Start the database via docker-compose
$ docker-compose up db
````

- Make sure `.test.env` exists with the right env vars.
- run test

```
TEST_DIR=$(pwd) go test ./...
```





