with import <nixpkgs> {};
stdenv.mkDerivation {
  name = "gluttony-shell";
  buildInputs = [
    go
    gocode
    glide
    godef
  ];
  shellHook = ''
    export GOPATH=~/projects
  '';
}
