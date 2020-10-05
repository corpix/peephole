peephole
---------

Simple proxy server. Project is under development.

What should be done before release:

- [ ] Configuration reload with `SIGHUP`
- [X] Statsd metrics
- [ ] Loadtest
- [X] Docker container

## Get

> You will need Go >=1.9
``` console
$ go get github.com/corpix/peephole
$ cd $GOPATH/src/github.com/corpix/peephole
```

## Run

> Will look for `config.yaml` in current directory by default.

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

- `PEEPHOLE_LISTEN` address in a format `0.0.0.0:9988` or `[::]:9988` to listen on
- `PEEPHOLE_CONFIG` relative or absolute path to a configuration file(`json`, `yaml`, `toml` [formats supported](https://github.com/corpix/formats#formats))
- `PEEPHOLE_LOG_LEVEL` set log level to one of `debug|info|warn|error`
- `PEEPHOLE_PROXY_METRICS_STATSDADDRESSES` comma-separated list of `ip:port` or `host:port` of statsd servers to report runtime telemetry

### Configuration file

Most simple example of configuration file is:

> This tool supports yaml configuration, however, json is more convenient to present and because json is fully interchangeable
> with yaml we will use it as an example.

``` yaml
{
  "Listen": "127.0.0.1:1338"
}
```

This proxy is configured to:

- listen on `127.0.0.1` port `1338`
- allow anonymous access
- allow connections anywhere

------------------------------------------

In some cases you might want to have two things:

- authentication
- whitelists
- blacklists(but they are not implemented just yet)

Here is an example of configuration describing the proxy for Telegram messenger, which was blocked in some countries.

This proxy is configured to:

> All coincidences with reality are accidental.

- listen on `127.0.0.1` port `1338`
- authenticate users against a set of accounts, where we have username `jarov` and password `g0t0gulag`
- deny anonymous access
- allow connections to the telegram networks addresses, where addresses are in [CIDR format](https://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing)
- allow connections to the telegram domains, where domains are [Go regexps](https://golang.org/pkg/regexp/syntax/#hdr-Syntax)

``` yaml
{
  "Listen": "127.0.0.1:1338",
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
        "149.154.160.0/20",
        "149.154.164.0/22",
        "149.154.168.0/22",
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

#### Metrics

If you have statsd endpoint available anywhere inside your infrastructure
then you could tell peephole to write runtime telemetry into this statsd
server/servers, here is a relevant peace of configuration:

``` yaml
{
  "Proxy": {
    "Metrics": {
      "StatsdAddresses": [ "127.0.0.1:8125" ]
    }
  }
}
```

## Credits

- Socks package is a fork of https://github.com/armon/go-socks5
