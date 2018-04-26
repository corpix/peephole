with import <nixpkgs> {};
stdenv.mkDerivation {
  name = "nix-cage-shell";
  buildInputs = [
    go
    gocode
    godef
    dep
    delve
    go-langserver
    nodePackages.tern
  ];
  shellHook = ''
    # To fix shitty golang tooling
    # Not all tools work good with vendor
    [ -d $(pwd)/.vendor ] || {
      mkdir -p $(pwd)/.vendor
      ln    -s $(pwd)/vendor  $(pwd)/.vendor/src
    }

    export GOPATH=$HOME/projects:$HOME/vendor
    export GOROOT=${go}/share/go
  '';
}
