# Golang REST API

Environment variables:
- configFilePath: string
- logFilePath: string
- jwtTokenSecret: string

Allows Queries:
```
http://localhost:3000/api/users?username=Antonio
```

## Generate Docs

```shell script
# Get swag
go get -u github.com/swaggo/swag/cmd/swag

# Generate docs
swag init --parseDependency --output docs
```

Run and go to [http://localhost:3000/docs/index.html](http://localhost:3000/docs/index.html)


## Run with Docker
   
1. **Build**

```shell script
make build
docker build . -t api-rest
```

2. **Run**

```shell script
docker run -p 3000:3000 api-rest
```

3. **Test**

```shell script
go test -v ./test/...