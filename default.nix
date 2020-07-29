{ pkgs ? import <nixpkgs> {} }:
with pkgs; buildGoPackage rec {
  name = "peephole-${version}";
  version = "1.0";

  buildInputs = [ git dep ];

  installPhase = ''
    source $stdenv/setup
    set -e

    mkdir -p              $out/bin
    cp    go/bin/peephole $out/bin
  '';

  src = ./.;
  goPackagePath = "github.com/corpix/peephole";
}
