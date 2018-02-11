with import <nixpkgs>{};
{ pkgs ? import <nixpkgs> {} }:

buildGo19Package rec {
  name = "go-boilerplate-unstable-${version}";
  version = "development";

  buildInputs = with pkgs; [ git glide ];

  src = ./.;
  goPackagePath = "github.com/corpix/go-boilerplate";
}
