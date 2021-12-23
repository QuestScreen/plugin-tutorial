{
  description = "Tutorial / demo plugin for QuestScreen";
  inputs = {
    nixpkgs.url = github:NixOS/nixpkgs/nixos-21.11;
    flake-utils.url = github:numtide/flake-utils;
    questscreen.url = github:QuestScreen/QuestScreen/nix;
  };
  outputs = {self, nixpkgs, flake-utils, questscreen}: with flake-utils.lib; eachSystem allSystems (system:
    let
      pkgs = import nixpkgs { inherit system; };
      discworld-meta = {
        goImportPath = "github.com/QuestScreen/plugin-tutorial";
        cssFiles = [ "style.css" ];
        modules.calendar = {
          configName = "calendarConfig";
          config = {
            Font = {
              package = "github.com/QuestScreen/api/config";
              type = "FontSelect";
              yamlName = "font";
            };
            Background = {
              package = "github.com/QuestScreen/api/config";
              type = "BackgroundSelect";
              yamlName = "background";
            };
          };
        };
        templates.systems.discworld = {
          name = "Discworld";
          description = "Terry Pratchett's Discworld";
        };
        templates.groups = [{
          name = "Discworld Group";
          description = "Contains a „Main“ scene with base modules and the calendar";
          system = "discworld.discworld";
          scenes = [{
            name = "Main";
            template = "discworld.default";
          }];
        }];
        templates.scenes.default = {
          name = "Default";
          description = "A scene with base modules and the calendar";
          config = ''
            modules:
              base.background:
                enabled: true
              base.herolist:
                enabled: true
              base.overlays:
                enabled: true
              base.title:
                enabled: true
              discworld.calendar:
                enabled: true
          '';
        };
      };
      discworld-plugin = pkgs.stdenvNoCC.mkDerivation {
        pname = "questscreen-discworld-plugin";
        version = self.shortRev or "dirty-${self.lastModifiedDate}";
        src = self;
        phases = [ "unpackPhase" "installPhase" ];
        installPhase = ''
          mkdir -p $out
          cp -r -t $out * 
        '';
      };
    in rec {
      packages = rec {
        discworld = discworld-plugin // discworld-meta;
        questscreen-discworld = questscreen.lib.buildQuestScreen {
          inherit pkgs;
          pname = "questscreen-discworld";
          version = self.shortRev or "dirty-${self.lastModifiedDate}";
          plugins = {inherit discworld; };
        };
      };
      defaultPackage = packages.questscreen-discworld;
    });
}