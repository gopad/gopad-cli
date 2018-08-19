# Gopad: CLI client

[![Build Status](http://drone.gopad.tech/api/badges/gopad/gopad-cli/status.svg)](http://drone.gopad.tech/gopad/gopad-cli)
[![Build Status](https://ci.appveyor.com/api/projects/status/qj9esfrpi3e7nd7d?svg=true)](https://ci.appveyor.com/project/gopadz/gopad-cli)
[![Stories in Ready](https://badge.waffle.io/gopad/gopad-api.svg?label=ready&title=Ready)](http://waffle.io/gopad/gopad-api)
[![Join the Matrix chat at https://matrix.to/#/#gopad:matrix.org](https://img.shields.io/badge/matrix-%23gopad%3Amatrix.org-7bc9a4.svg)](https://matrix.to/#/#gopad:matrix.org)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/8b885bc21c374fa3a5e661b3ad9d9a65)](https://www.codacy.com/app/gopad/gopad-cli?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=gopad/gopad-cli&amp;utm_campaign=Badge_Grade)
[![Go Doc](https://godoc.org/github.com/gopad/gopad-cli?status.svg)](http://godoc.org/github.com/gopad/gopad-cli)
[![Go Report](http://goreportcard.com/badge/github.com/gopad/gopad-cli)](http://goreportcard.com/report/github.com/gopad/gopad-cli)
[![](https://images.microbadger.com/badges/image/gopad/gopad-cli.svg)](http://microbadger.com/images/gopad/gopad-cli "Get your own image badge on microbadger.com")


**This project is under heavy development, it's not in a working state yet!**

Within this repository we are building the command-line client to interact with the [Gopad API](https://github.com/gopad/gopad-api) server, for further information take a look at our [documentation](https://gopad.tech).


## Install

You can download prebuilt binaries from the GitHub releases or from our [download site](http://dl.gopad.tech/cli). You are a Mac user? Just take a look at our [homebrew formula](https://github.com/gopad/homebrew-gopad).


## Development

Make sure you have a working Go environment, for further reference or a guide take a look at the [install instructions](http://golang.org/doc/install.html). This project requires Go >= v1.8.

```bash
go get -d github.com/gopad/gopad-cli
cd $GOPATH/src/github.com/gopad/gopad-cli
make clean generate build

./gopad-cli -h
```


## Security

If you find a security issue please contact gopad@webhippie.de first.


## Contributing

Fork -> Patch -> Push -> Pull Request


## Authors

* [Thomas Boerger](https://github.com/tboerger)


## License

Apache-2.0


## Copyright

```
Copyright (c) 2018 Thomas Boerger <thomas@webhippie.de>
```
