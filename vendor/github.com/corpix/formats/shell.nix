with import <nixpkgs> {};
stdenv.mkDerivation {
  name = "nix-cage-shell";
  buildInputs = [
    go
    gocode
    godef
    go2nix
    dep
    delve
    go-langserver
  ];
  shellHook = ''
    export GOPATH=$HOME/projects
    export GOROOT=${go}/share/go
  '';
}
