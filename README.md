# Gopad: CLI client

[![General Workflow](https://github.com/gopad/gopad-cli/actions/workflows/general.yml/badge.svg)](https://github.com/gopad/gopad-cli/actions/workflows/general.yml) [![Join the Matrix chat at https://matrix.to/#/#gopad:matrix.org](https://img.shields.io/badge/matrix-%23gopad-7bc9a4.svg)](https://matrix.to/#/#gopad:matrix.org) [![Codacy Badge](https://app.codacy.com/project/badge/Grade/7a3bf170b3524feeb3ed129b02c80759)](https://app.codacy.com/gh/gopad/gopad-cli/dashboard?utm_source=gh&utm_medium=referral&utm_content=&utm_campaign=Badge_grade) [![Go Reference](https://pkg.go.dev/badge/github.com/gopad/gopad-cli.svg)](https://pkg.go.dev/github.com/gopad/gopad-cli) [![GitHub Repo](https://img.shields.io/badge/github-repo-yellowgreen)](https://github.com/gopad/gopad-cli)

Within this repository we are building the command-line client to interact with
the [Gopad API][api] server.

## Install

You can download prebuilt binaries from the [GitHub releases][releases] or from
our [download site][downloads]. If you prefer to use containers you could use
our images published on [Docker Hub][dockerhub] or [Quay][quay]. You are a Mac
user? Just take a look at our [homebrew formula][homebrew]. If you need further
guidance how to install this take a look at our [documentation][docs].

## Development

Make sure you have a working Go environment, for further reference or a guide
take a look at the [install instructions][golang]. This project requires
Go >= v1.18, at least that's the version we are using.

```console
git clone https://github.com/gopad/gopad-cli.git
cd gopad-cli

make generate build

./bin/gopad-cli -h
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
[downloads]: https://dl.gopad.eu/cli
[homebrew]: https://github.com/gopad/homebrew-gopad
[dockerhub]: https://hub.docker.com/r/gopad/gopad-cli/tags/
[quay]: https://quay.io/repository/gopad/gopad-cli?tab=tags
[docs]: https://gopad.eu/
[golang]: http://golang.org/doc/install.html
