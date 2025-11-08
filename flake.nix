{
  description = "Prayer times CLI to remind you to Go pray";

  inputs = { nixpkgs.url = "github:nixos/nixpkgs/nixos-unstable"; };

  outputs = { self, nixpkgs }:
    let
      systems = [ "x86_64-linux" "aarch64-linux" ];
      forAllSystems = f:
        nixpkgs.lib.genAttrs systems
        (system: f { pkgs = nixpkgs.legacyPackages.${system}; });
    in {
      packages = forAllSystems
        ({ pkgs }: { default = pkgs.callPackage ./nix/pkgs/go-pray.nix { }; });

      homeManagerModules = {
        go-pray = ./nix/modules/home-manager/go-pray.nix;
        default = self.homeManagerModules.go-pray;
      };
    };
}
