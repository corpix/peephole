{ config, lib, pkgs, ... }:

with lib;

let
  name = "peephole";
  cfg = config.services."${name}";
  pkg = (pkgs.callPackage ./default.nix { }).bin;
in {
  options = with types; {
    services."${name}" = {
      enable = mkEnableOption "Peephole, crypto currency exchange data spy";
      user = mkOption {
        default = name;
        type = str;
        description = ''
          User name to run service from.
        '';
      };
      group = mkOption {
        default = name;
        type = str;
        description = ''
          Group name to run service from.
        '';
      };
      logger = mkOption {
        default = { level = "info"; format = "json"; };
        type = submodule ({ ... }: {
          options = {
            level = mkOption {
              type = enum [ "debug" "info" "error" ];
              default = "error";
              description = ''
                Logging verbosity level.
              '';
            };
            format = mkOption {
              type = enum [ "json" "text" ];
              default = "json";
              description = ''
                Logging format.
              '';
            };
          };
        });
      };
      listen = mkOption {
        default = "127.0.0.1:1338";
        type = str;
        description = ''
          Address to listen for proxy requests.
        '';
      };
      proxy = mkOption {
        type = submodule ({ ... }: {
          options = {
            accounts = mkOption {
              default = { };
              type = attrsOf str;
              description = ''
                Accounts in a form <user> = <password>.
              '';
            };

            whitelist = mkOption {
              default = { };
              type = submodule ({ ... }: {
                options = {
                  addresses = mkOption {
                    type = listOf str;
                    default = [];
                    description = ''
                      List of addresses clients are allowed to connect to.
                    '';
                  };
                  domains = mkOption {
                    type = listOf str;
                    default = [];
                    description = ''
                      List of domains clients are allowed to connect to.
                    '';
                  };
                };
              });
            };

            metrics = mkOption {
              type = submodule ({ ... }: {
                options = {
                  statsdAddresses = mkOption {
                    type = listOf str;
                    default = [];
                    description = ''
                      List of host:port statsd addresses to report metrics to.
                    '';
                  };
                };
              });
              default = { };
              description = ''
                Metrics configuration which describes statsd server addresses, etc.
              '';
            };
          };
        });
        default = { };
        description = ''
          Proxy configuration which describes whitelisting, authentication, etc.
        '';
      };
    };
  };

  config = mkIf cfg.enable {
    users.extraUsers = mkIf (name == cfg.user) {
      "${name}" = {
        name = name;
        group = cfg.group;
      };
    };

    users.extraGroups = mkIf (name == cfg.group) {
      "${name}" = { inherit name; };
    };

    systemd.services."${name}" = let
      configuration = {
        Logger.Format                 = cfg.logger.format;
        Logger.Level                  = cfg.logger.level;
        Listen                        = cfg.listen;
        Proxy.Accounts                = cfg.proxy.accounts;
        Proxy.Whitelist.Addresses     = cfg.proxy.whitelist.addresses;
        Proxy.Whitelist.Domains       = cfg.proxy.whitelist.domains;
        Proxy.Metrics.StatsdAddresses = cfg.proxy.metrics.statsdAddresses;
      };
    in {
      enable = true;

      wantedBy = [ "multi-user.target" ];
      after    = [ "network.target" ];

      serviceConfig = {
        Type = "simple";
        User = name;
        Group = name;
        ExecStart = "${pkg}/bin/${name} -c ${pkgs.writeText "config.json" (builtins.toJSON configuration)}";
        Restart = "on-failure";
        RestartSec = 1;
      };
    };
  };
}
