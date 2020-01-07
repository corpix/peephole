{ pkgs ? import <nixpkgs> {} }:
with pkgs; buildGoPackage rec {
  name = "peephole-${version}";
  version = "1.0";

  buildInputs = [ git dep ];

  installPhase = ''
    source $stdenv/setup
    set -e

    mkdir -p              $bin/bin
    cp    go/bin/peephole $bin/bin
  '';

  src = ./.;
  goPackagePath = "github.com/corpix/peephole";
}
