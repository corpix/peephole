with import <nixpkgs>{};
{ pkgs ? import <nixpkgs> {} }:

buildGo19Package rec {
  name = "formats-unstable-${version}";
  version = "development";

  buildInputs = with pkgs; [ git dep ];
  installPhase = ''
    source $stdenv/setup
    set -e

    mkdir -p             $bin/bin
    cp    go/bin/formats $bin/bin
  '';

  goDeps = ./deps.nix;

  src = ./.;
  goPackagePath = "github.com/corpix/formats";
}
