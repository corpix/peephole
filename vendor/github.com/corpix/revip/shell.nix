let nixpkgs = <nixpkgs>;
    config = {};
in with import nixpkgs { inherit config; }; let
  shellWrapper = writeScript "shell-wrapper" ''
    #! ${stdenv.shell}
    set -e

    exec -a shell ${fish}/bin/fish --login --interactive --init-command='
      set -x root '"$root"'
      set config $root/.fish.conf
      set presonal_config $root/.personal.fish.conf
      if test -e $personal_config
        source $personal_config
      end
      if test -e $config
        source $config
      end
    ' "$@"
  '';
in stdenv.mkDerivation rec {
  name = "nix-shell";
  buildInputs = [
    glibcLocales bashInteractive man
    nix cacert curl utillinux coreutils
    git jq tmux findutils gnumake

    go gopls
  ];
  shellHook = ''
    export root=$(pwd)

    if [ -f "$root/.env" ]
    then
      source "$root/.env"
    fi

    export LANG="en_US.UTF-8"
    export SHELL="${shellWrapper}"
    export NIX_PATH="nixpkgs=${nixpkgs}"

    exec "$SHELL"
  '';
}
