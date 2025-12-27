{
  description = "Urverk backend";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs =
    {
      self,
      nixpkgs,
      flake-utils,
    }:
    flake-utils.lib.eachDefaultSystem (
      system:
      let
        pkgs = import nixpkgs { inherit system; };
        CGO = "0";
        go = pkgs.go;
      in
      {
        packages.default = pkgs.buildGoModule {
          pname = "urverk-backend";
          version = "0.1.0";
          src = ./.;

          vendorHash = "sha256-39n1qCvAeQJeYmGPII8/D3xprRYFTM2Epg3JWzu1hZQ=";

          env = {
            CGO_ENABLED = CGO;
          };

          doCheck = true;
        };

        devShells.default = pkgs.mkShell {
          packages = with pkgs; [
            go
            delve
            gotools
          ];

          env = {
            CGO_ENABLED = CGO;
          };

          shellHook = ''
            echo "🦫 $(${go}/bin/go version) ready!"
            echo "CGO_ENABLED: $CGO_ENABLED"
          '';
        };
      }
    );
}
