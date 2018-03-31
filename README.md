## Intro
Prototype of website.

## Compile
**Make sure your Go development environment is ready.**

1. Download dependency:
``` shell
$ go get -u github.com/julienschmidt/httprouter
$ go get -u gopkg.in/russross/blackfriday.v2
```
2. Download code:
``` shell
$ git clone git@github.com:htmc-io/website-proto.git
```
3. Build:
``` shell
$ cd website-proto
$ go build -o ./main ./src
```

## Run
``` shell
$ cd website-proto
$ ./main 8080 # run server on port 8080, if no port is specified, use 80
```

## Test
Open web browser, accesss http://localhost:8080 .




