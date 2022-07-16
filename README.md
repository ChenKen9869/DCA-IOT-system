# Farm Management Platform Backend

write by :

​	**golang**

​	**gin**

​	**gorm**

server port : 5930

database: 

​	**mysql** for management information. 

​	**mongodb** for sensor data.

### run the server

```go
go run .
```

### or build to run on linux

```shel
$ENV:GOOS="linux"
go build .
./go-backend
```

### see more detail of api : (after run the server)

 http://localhost:5930/swagger/index.html
