with import <nixpkgs> {};
stdenv.mkDerivation {
  name = "nix-cage-shell";
  buildInputs = [
    go
    gocode
    godef
    dep
  ];
  shellHook = ''
    export GOPATH=~/projects
  '';
}
