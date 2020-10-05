{ pkgs ? import <nixpkgs> {} }:
with pkgs; buildGoModule rec {
  name = "peephole";
  src = ./..;
  vendorSha256 = null;
}
