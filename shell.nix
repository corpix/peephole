with import <nixpkgs> {};
stdenv.mkDerivation {
  name = "nix-cage-shell";
  buildInputs = [
    go
    gocode
    godef
    dep
    delve
  ];
  shellHook = ''
    export GOPATH=~/projects
  '';
}
