{
  description = "A CLI and server for interacting with WorldAnvil and local markdown files.";

  inputs.nixpkgs.url = "https://flakehub.com/f/NixOS/nixpkgs/*";

  outputs = inputs: let
    goVersion = 24;

    supportedSystems = [
      "x86_64-linux"
      "aarch64-linux"
      "x86_64-darwin"
      "aarch64-darwin"
    ];
    forEachSupportedSystem = f:
      inputs.nixpkgs.lib.genAttrs supportedSystems (
        system:
          f {
            pkgs = import inputs.nixpkgs {
              inherit system;
              overlays = [inputs.self.overlays.default];
            };
          }
      );
  in {
    overlays.default = final: prev: {
      go = final."go_1_${toString goVersion}";
    };

    formatter = forEachSupportedSystem ({pkgs}: pkgs.alejandra);

    devShells = forEachSupportedSystem (
      {pkgs}: let
        generateOpenApiClient = let
          git = pkgs.lib.getExe pkgs.git;
          go = pkgs.lib.getExe pkgs.go;
          swagger-cli = pkgs.lib.getExe pkgs.swagger-cli;
          openapi-generator-cli = pkgs.lib.getExe pkgs.openapi-generator-cli;

          pkgsCfg = {
            disallowAdditionalPropertiesIfNotPresent = false;
            enumClassPrefix = true;
            withGoCodegenComment = true;
            generateInterfaces = true;
            isGoSubmodule = true;
            withGoMod = false;
          };

          stringProps = builtins.concatStringsSep "," (pkgs.lib.mapAttrsToList (k: v: "${k}=${toString v}") pkgsCfg);
        in
          pkgs.writeShellScriptBin "generate-openapi-client" ''
            #!/usr/bin/env bash
            set -euo pipefail

            cd "$(${git} rev-parse --show-toplevel)"

            output_dir="pkg/api/worldanvil"
            package_name="$(basename "$output_dir")"

            rm -rf "$output_dir"
            mkdir -p "$output_dir"

            ${go} run ./cmd/cli download-spec

            # ${swagger-cli} bundle -o spec/full.yml -t yaml -w 100 spec/openapi.yml

            ${openapi-generator-cli} generate \
              -i "spec/openapi.yml" \
              -g go \
              -o "$output_dir" \
              --skip-validate-spec \
              --additional-properties="packageName=$package_name,${stringProps}"

            cd "$output_dir"

            # Remove unnecessary files
            rm -rf .* api docs test git_push.sh
          '';
      in {
        default = pkgs.mkShell {
          packages = with pkgs; [
            go
            gopls
            gotools
            # https://github.com/golangci/golangci-lint
            golangci-lint

            # swagger-cli
            # swagger-codegen
            generateOpenApiClient
          ];
        };
      }
    );
  };
}
