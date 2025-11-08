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
      packages = forAllSystems ({ pkgs }: {
        default = let
          inherit (pkgs)
            lib stdenv fetchFromGitHub buildGoModule installShellFiles
            pkg-config alsa-lib;
        in buildGoModule (finalAttrs: {
          pname = "go-pray";
          version = "0.1.14";

          src = fetchFromGitHub {
            owner = "0xzer0x";
            repo = "go-pray";
            tag = "v${finalAttrs.version}";
            hash = "sha256-8DMGw/lHawORwELIedcml2Kew/q5EGL1Jbf70amHuJU=";
          };

          vendorHash = "sha256-qMTg2Vsk0nte1O8sbNWN5CCCpgpWLvcb2RuGMoEngYE=";

          nativeBuildInputs = [ pkg-config installShellFiles ];

          buildInputs = [ alsa-lib ];

          ldflags = [
            "-X github.com/0xzer0x/go-pray/internal/version.version=${finalAttrs.version}"
            "-X github.com/0xzer0x/go-pray/internal/version.buildTime=1980-01-01T00:00:00Z"
          ];

          # NOTE: Create temporary config file to supress missing config error
          postInstall = lib.optionalString
            (stdenv.buildPlatform.canExecute stdenv.hostPlatform) ''
              GOPRAY_TMPCONF="$(mktemp --suffix=.yml)"
              printf 'calculation: { method: "UAQ" }\nlocation: { lat: 0, long: 0 }\n' >"''${GOPRAY_TMPCONF}"
              installShellCompletion --cmd go-pray \
                --bash <($out/bin/go-pray --config="''${GOPRAY_TMPCONF}" completion bash) \
                --fish <($out/bin/go-pray --config="''${GOPRAY_TMPCONF}" completion fish) \
                --zsh <($out/bin/go-pray --config="''${GOPRAY_TMPCONF}" completion zsh)
              rm "''${GOPRAY_TMPCONF}"
            '';

          meta = {
            description = "Prayer times CLI to remind you to Go pray";
            homepage = "https://github.com/0xzer0x/go-pray";
            changelog =
              "https://github.com/0xzer0x/go-pray/releases/tag/v${finalAttrs.version}";
            license = lib.licenses.gpl3Plus;
            maintainers = with lib.maintainers; [ _0xzer0x ];
            platforms = lib.platforms.linux;
            mainProgram = "go-pray";
          };
        });
      });
    };
}
