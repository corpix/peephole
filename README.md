go-boilerplate
---------

[![Build Status](https://travis-ci.org/corpix/go-boilerplate.svg?branch=master)](https://travis-ci.org/corpix/go-boilerplate)

## Bootstrap

I have wrote a bootstrap script for you. It is wrote in python. Here is a CLI help:

``` console
$ ./bootstrap -h
usage: bootstrap [-h] --name NAME [--user USER] [--host HOST] [--description DESCRIPTION] target

Bootstrap a project from the boilerplate

positional arguments:
  target                Target directory to create project in

optional arguments:
  -h, --help            show this help message and exit
  --name NAME           Project name
  --user USER           Project user/org to use in imports
  --host HOST           Project host to use in imports
  --description DESCRIPTION
                        Project description to hardcode into
```

To bootstrap a new project named `test` for a github user `corpix` you could run:

``` console
$ ./bootstrap --description 'Hello world' --name test --user corpix --host github.com $GOPATH/src/github.com/corpix/test
```

New project in `$GOPATH/src/github.com/corpix/test` is waiting for you :)

## Development

All development process accompanied by containers. Docker containers used for development, Rkt containers used for production.

> I am a big fan of Rkt, but it could be comfortable for other developers to use Docker for development and testing.

## Optional requirements

If you plan to use some containerisation then you should have:

- [docker](https://github.com/moby/moby)
- [docker-compose](https://github.com/docker/compose)

Or:

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
