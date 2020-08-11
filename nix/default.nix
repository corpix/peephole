{ pkgs ? import <nixpkgs> {} }:
with pkgs; buildGoModule rec {
  pname = "peephole";
  version = "1.0";

  src = ./..;

  vendorSha256 = null;
}
