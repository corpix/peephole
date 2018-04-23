peephole
---------

[![Build Status](https://travis-ci.org/corpix/peephole.svg?branch=master)](https://travis-ci.org/corpix/peephole)

Simple proxy server.

### Running peephole

#### No isolation

``` console
$ go run ./peephole/peephole.go --debug
```

#### Docker

``` console
$ docker-compose up peephole
```

Build a binary release:

``` console
$ GOOS=linux make
$ ./build/peephole
```


