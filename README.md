# Set Up

```bash
$ go get github.com/DonghuiZhuo/casemgr
```

# Unit Test

```bash
$ cd ${PROJECT_DIR}
$ go test -v ./...
```

# Integ Test

```bash
$ cd ${PROJECT_DIR}
$ go run main.go

# in a different terminal
$ curl http://localhost:8080/assets/
```

# References
* [Build web application in Go](https://www.sohamkamani.com/blog/2017/09/13/how-to-build-a-web-application-in-golang/)
