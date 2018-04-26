{
  # Example configuration to proxy telegram messenger traffic.
  # It will allow an authenticated user with login "jarov" and
  # password "g0t0gulag"(all coincidences with reality are accidental)
  # to connect to hosts which was specified in Targets.

  #Logger.Formatter = "json";

  Addr = "127.0.0.1:1338";

  #Accounts = { "jarov" = "g0t0gulag"; };

  Addresses = [
    # telegram ipv4
    "91.108.4.0/22"
    "91.108.8.0/22"
    "91.108.12.0/22"
    "91.108.16.0/22"
    "91.108.56.0/22"
    "91.108.56.0/23"
    "91.108.56.0/24"
    "149.154.160.0/20"
    "149.154.160.0/22"
    "149.154.164.0/22"
    "149.154.168.0/22"
    "149.154.168.0/23"
    "149.154.170.0/23"

    # telegram ipv6
    "2001:67c:4e8::/48"
    "2001:b28:f23d::/48"
    "2001:b28:f23e::/48"
    "2001:b28:f23f::/48"
  ];

  Domains = [
    "^(.*\.)?t\.me$"
    "^(.*\.)?telegram\.org$"
  ];
}
