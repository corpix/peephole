with import <nixpkgs>{};
{ pkgs ? import <nixpkgs> {} }:

buildGoPackage rec {
  name = "peephole-${version}";
  version = "1.0";

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
