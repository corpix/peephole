{ config, lib, pkgs, ... }:
with lib;
let
  name = "peephole";
  cfg = config.services."${name}";
  pkg = pkgs.callPackage ./default.nix { };
in {
  options = with types; {
    services."${name}" = {
      enable = mkEnableOption "Peephole, crypto currency exchange data spy";
      limitNoFile = mkOption {
        default = 16000;
        type = int;
        description = "FDs limit";
      };
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

      config = mkOption {
        type = attrs;
        default = { };
        description = ''
          Peephole configuration.
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
      configuration = recursiveUpdate
        (import ../config.nix)
        cfg.config;
    in {
      enable = true;

      wantedBy = [ "multi-user.target" ];
      after    = [ "network.target" ];

      serviceConfig = {
        Type = "simple";
        User = name;
        Group = name;
        ExecStart = "${pkg}/bin/${name} -c ${pkgs.writeText "config.yaml" (builtins.toJSON configuration)}";
        Restart = "on-failure";
        RestartSec = 1;
        LimitAS = "infinity";
        LimitRSS = "infinity";
        LimitCORE = "infinity";
        LimitNOFILE = cfg.limitNoFile;
      };
    };
  };
}
