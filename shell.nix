let
  nixpkgs = builtins.fetchTarball {
    # https://github.com/NixOS/nixpkgs/tags
    # $ nix-prefetch-url --unpack url
    #   hash...
    url    = "https://github.com/NixOS/nixpkgs/archive/96d00ea0d9860b8bfba816b8ed21a98175a6a961.tar.gz";
    sha256 = "00zbcihpcvx8l5p7gl2n65z46r4rn5lkwrvgsf3nsl9gxv8vrm8z";
  };
in with import nixpkgs { };
let
  shellWrapper = writeScript "shell-wrapper" ''
    #! ${stdenv.shell}
    exec -a shell ${fish}/bin/fish "$@"
  '';
in stdenv.mkDerivation {
  name = "nix-shell";
  buildInputs = with pkgs; [
    glibcLocales man nix cacert coreutils git gnumake
    tmux jq docker skopeo
    curl utillinux bash-completion
    go gopls golangci-lint
    gnumake clang pkgconfig
  ];
  shellHook = ''
    export REPO_ROOT=$(git rev-parse --show-toplevel)

    export SHELL="${shellWrapper}"
    export LANG=en_US.UTF-8
    export NIX_PATH="nixpkgs=${nixpkgs}"

    unset GOPATH
  '';
}
