{ config, lib, pkgs, ... }:

with lib;
let
  yamlFormat = pkgs.formats.yaml { };
  cfg = config.services.go-pray;
in {
  options = {
    services.go-pray = {
      enable = mkEnableOption "Prayer times notification daemon";

      package = mkOption {
        type = types.package;
        default = pkgs.callPackage ../../pkgs/go-pray.nix { };
        defaultText = literalExpression "pkgs.go-pray";
        description = "Package providing {command}`go-pray`.";
      };

      settings = mkOption {
        type = yamlFormat.type;
        default = { };
        example = lib.literalExpression ''
          {
            language = "en";
            calculation.method = "MWL";
            timezone = "Europe/London";
            location = {
              lat = 51.5176;
              long = -0.1612;
            };
            notification = {
              icon = "xclock";
              title = "Prayer Time";
              body = "Time for {{ .CalendarName }} prayer 🕌";
            };
          }
        '';
      };
    };
  };

  config = mkIf cfg.enable {
    home.packages = [ cfg.package ];

    xdg.configFile."go-pray/config.yml" = mkIf (cfg.settings != { }) {
      source = (yamlFormat.generate "go-pray-config.yml" cfg.settings);
    };

    # NOTE: Create systemd user unit to autostart daemon
    systemd.user.services.go-pray = {
      Unit = {
        Description = "Prayer times notification daemon";
        Documentation = "https://github.com/0xzer0x/go-pray";
        After = [ "graphical-session-pre.target" ];
        PartOf = [ "graphical-session.target" ];
      };

      Service = let configFile = "${config.xdg.configHome}/go-pray/config.yml";
      in {
        Type = "exec";
        Environment = [ "PATH=${config.home.profileDirectory}/bin" ];
        ExecCondition = [ "/bin/sh -c 'test -f ${configFile}'" ];
        ExecStart = [ "${cfg.package}/bin/go-pray daemon" ];
      };

      Install = { WantedBy = [ "graphical-session.target" ]; };
    };
  };
}
