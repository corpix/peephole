peephole
---------

[![Build Status](https://travis-ci.org/corpix/peephole.svg?branch=master)](https://travis-ci.org/corpix/peephole)

Simple proxy server. Project is under development.

## Get

> You will need Go >=1.9
``` console
$ go get github.com/corpix/peephole
$ cd $GOPATH/src/github.com/corpix/peephole
```

## Run

``` console
$ go run ./peephole/peephole.go --debug
```

### Docker

> If you use something other than Linux then
> You should run `make` like this `make GOOS=linux`
> Otherwise your container will not work

``` console
$ make
$ docker-compose up peephole
```

## Configuration

### Env variables

- `ADDR``
- `CONFIG`
- `DEBUG`
