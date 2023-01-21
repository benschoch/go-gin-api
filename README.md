# recipes api
go 1.19, gin-gonic, mongoDB6
```shell
docker run -d --name go-mongo-rest -p 127.0.0.1:27017:27017 -e MONGO_INITDB_ROOT_USERNAME=root -e MONGO_INITDB_ROOT_PASSWORD=pass mongo
```
run app:
```shell
go run cmd/recipes-core-rest/main.go
```

### references:
https://github.com/gin-gonic/gin  
https://github.com/swaggo/gin-swagger  
https://github.com/golang-standards/project-layout  
https://gist.github.com/mbchoa/a9032d05198c1f76a680c94d6bc8f290