{ pkgs, lib, config, inputs, ... }:

{
  # https://devenv.sh/packages/
  packages = with pkgs; [ go-task gcc pkg-config alsa-lib.dev ];

  # https://devenv.sh/languages/
  languages.go.enable = true;

  scripts = {
    go-pray.exec = ''
      go run . "''${@}"
    '';
  };

  tasks = {
    "go-pray:build" = {
      exec = ''
        ${pkgs.go-task}/bin/task build
      '';
      execIfModified = [ "cmd" "internal" "main.go" "go.mod" "go.sum" ];
    };
  };
}
