with import <nixpkgs>{};
{ pkgs ? import <nixpkgs> {} }:

buildGo19Package rec {
  name = "peephole-unstable-${version}";
  version = "development";

  buildInputs = with pkgs; [ git dep ];

  installPhase = ''
    source $stdenv/setup
    set -e

    mkdir -p              $bin/bin
    cp    go/bin/peephole $bin/bin
  '';

  src = ./.;
  goPackagePath = "github.com/corpix/peephole";
}
