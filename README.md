# LBRY CLI

[![Go](https://github.com/LBRYFoundation/lbry-cli/actions/workflows/go.yml/badge.svg)](https://github.com/LBRYFoundation/lbry-cli/actions/workflows/go.yml)
[![Codecov](https://codecov.io/gh/LBRYFoundation/lbry-cli/graph/badge.svg)](https://codecov.io/gh/LBRYFoundation/lbry-cli)
![Snapcraft Version](https://img.shields.io/snapcraft/v/lbry-cli/latest/stable)
![WinGet Package Version](https://img.shields.io/winget/v/LBRY.CLI)

The LBRY Command Line Interface (CLI) to interact with the LBRY Daemon.

## Build

```shell
go build -o lbry-cli
```

Or for Windows:

```shell
go build -o lbry-cli.exe
```

## Install

### Snap

```shell
snap install lbry-cli
```

### WinGet

```shell
winget install LBRY.CLI
```

## Run

```shell
lbry-cli version
```

## License
This project is MIT licensed. For the full license, see [LICENSE](LICENSE.md).
