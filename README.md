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

> Will look for `config.json` in current directory by default.

``` console
$ go run ./peephole/peephole.go --debug
```

### Docker

> If you use something other than Linux then       <br/>
> You should run `make` like this `make GOOS=linux <br/>`
> Otherwise your container will not work

``` console
$ make
$ docker-compose up peephole
```

## Configuration

### Env variables

- `ADDR` address in a format `0.0.0.0:9988` or `[::]:9988` to listen on
- `CONFIG` relative or absolute path to a configuration file(`json`, `yaml`, `toml` [formats supported](https://github.com/corpix/formats#formats))
- `DEBUG` if set then run with `debug` log level

### Configuration file

Configuration file could be presented in formats: `json`, `yaml`, `toml` and maybe other formats which is supported [by this package](https://github.com/corpix/formats#formats). 

> We determine what format to use by file extension. <br/>
> In development `config.nix` is used as a «single point of trust»,
> all configuration files in other formats are generated from configuration in Nix language, but
> we will use config file in `json` as an example to keep things simple.

Here is an example of configuration describing the proxy for Telegram messenger, which was blocked in some countries.

This proxy is configured to:

> All coincidences with reality are accidental.

- listen on `127.0.0.1` port `1338`
- authenticate users against a set of accounts, where we have username `jarov` and password `g0t0gulag`
- deny anonymous access
- allow connections to the telegram networks addresses, where addresses are in [CIDR format](https://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing)
- allow connections to the telegram domains, where domains are [Go regexps](https://golang.org/pkg/regexp/syntax/#hdr-Syntax)

``` json
{
  "Addr": "127.0.0.1:1338",
  "Proxy": {
    "Accounts": {
      "jarov": "g0t0gulag"
    },
    "Whitelist": {
      "Addresses": [
        "91.108.4.0/22",
        "91.108.8.0/22",
        "91.108.12.0/22",
        "91.108.16.0/22",
        "91.108.56.0/22",
        "91.108.56.0/23",
        "91.108.56.0/24",
        "149.154.160.0/20",
        "149.154.160.0/22",
        "149.154.164.0/22",
        "149.154.168.0/22",
        "149.154.168.0/23",
        "149.154.170.0/23",
        "2001:67c:4e8::/48",
        "2001:b28:f23d::/48",
        "2001:b28:f23e::/48",
        "2001:b28:f23f::/48"
      ],
      "Domains": [
        "^(.*.)?t.me$",
        "^(.*.)?telegram.org$"
      ]
    }
  }
}
```

> In case of regexps: don't forget to pay your attention to `^` and `$`

## Credits

- Socks package is a fork of https://github.com/armon/go-socks5
