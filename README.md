# Gopad: CLI client

[![General Workflow](https://github.com/gopad/gopad-cli/actions/workflows/general.yml/badge.svg)](https://github.com/gopad/gopad-cli/actions/workflows/general.yml) [![Join the Matrix chat at https://matrix.to/#/#gopad:matrix.org](https://img.shields.io/badge/matrix-%23gopad-7bc9a4.svg)](https://matrix.to/#/#gopad:matrix.org) [![Codacy Badge](https://app.codacy.com/project/badge/Grade/7a3bf170b3524feeb3ed129b02c80759)](https://app.codacy.com/gh/gopad/gopad-cli/dashboard?utm_source=gh&utm_medium=referral&utm_content=&utm_campaign=Badge_grade) [![Go Reference](https://pkg.go.dev/badge/github.com/gopad/gopad-cli.svg)](https://pkg.go.dev/github.com/gopad/gopad-cli) [![GitHub Repo](https://img.shields.io/badge/github-repo-yellowgreen)](https://github.com/gopad/gopad-cli)

> [!CAUTION]
> This project is in active development and does not provide any stable release
> yet, you can expect breaking changes until our first real release!

Within this repository we are building the command-line client to interact with
the [Gopad API][api] server.

## Install

You can download prebuilt binaries from the [GitHub releases][releases] or from
our [download site][downloads]. Besides that we also prepared repositories for
DEB and RPM packages which can be found at [Baltorepo][baltorepo]. If you prefer
to use containers you could use our images published on [GHCR][ghcr],
[Docker Hub][dockerhub] or [Quay][quay]. You are a Mac user? Just take a look
at our [homebrew formula][homebrew]. If you need further guidance how to
install this take a look at our [documentation][docs].

## Build

If you are not familiar with [Nix][nix] it is up to you to have a working
environment for Go (>= 1.24.0) as the setup won't we covered within this guide.
Please follow the official install instructions for [Go][golang]. Beside that
we are using [go-task][gotask] to define all commands to build this project.

```console
git clone https://github.com/gopad/gopad-cli.git
cd gopad-cli

task generate build
./bin/gopad-cli -h
```

If you got [Nix][nix] and [Direnv][direnv] configured you can simply execute
the following commands to get al dependencies including [go-task][gotask] and
the required runtimes installed:

```console
cat << EOF > .envrc
use flake . --impure --extra-experimental-features nix-command
EOF

direnv allow
```

## Security

If you find a security issue please contact
[gopad@webhippie.de](mailto:gopad@webhippie.de) first.

## Contributing

Fork -> Patch -> Push -> Pull Request

## Authors

-   [Thomas Boerger](https://github.com/tboerger)

## License

Apache-2.0

## Copyright

```console
Copyright (c) 2018 Thomas Boerger <thomas@webhippie.de>
```

[api]: https://github.com/gopad/gopad-cli
[releases]: https://github.com/gopad/gopad-cli/releases
[downloads]: https://dl.gopad.eu
[homebrew]: https://github.com/gopad/homebrew-gopad
[ghcr]: https://github.com/orgs/gopad/packages
[dockerhub]: https://hub.docker.com/r/gopad/gopad-cli/tags/
[quay]: https://quay.io/repository/gopad/gopad-cli?tab=tags
[docs]: https://gopad.eu/
[nix]: https://nixos.org/
[golang]: http://golang.org/doc/install.html
[gotask]: https://taskfile.dev/installation/
[direnv]: https://direnv.net/
