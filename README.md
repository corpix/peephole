go-boilerplate
---------

[![Build Status](https://travis-ci.org/corpix/go-boilerplate.svg?branch=master)](https://travis-ci.org/corpix/go-boilerplate)

## Development

All development process accompanied by containers. Docker containers used for development, Rkt containers used for production.

> I am a big fan of Rkt, but it could be comfortable for other developers to use Docker for development and testing.

## Requirements

- [docker](https://github.com/moby/moby)
- [docker-compose](https://github.com/docker/compose)
- [jq](https://github.com/stedolan/jq)
- [rkt](https://github.com/coreos/rkt)
- [acbuild](https://github.com/containers/build)

### Running go-boilerplate

Build a binary release:

``` console
$ GOOS=linux make
# This will put a binary into ./build/go-boilerplate
```

#### Docker

``` console
$ docker-compose up go-boilerplate
```

#### Rkt

There is no rkt container for this service at this time.

#### No isolation

``` console
$ go run ./go-boilerplate/go-boilerplate.go --debug
```
