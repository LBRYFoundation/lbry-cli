# LBRY CLI

[![Go](https://github.com/LBRYFoundation/lbry-cli/actions/workflows/go.yml/badge.svg)](https://github.com/LBRYFoundation/lbry-cli/actions/workflows/go.yml)
[![Codecov](https://codecov.io/gh/LBRYFoundation/lbry-cli/graph/badge.svg)](https://codecov.io/gh/LBRYFoundation/lbry-cli)
![Snapcraft Version](https://img.shields.io/snapcraft/v/lbry-cli/latest/stable?logo=snapcraft)
![WinGet Package Version](https://img.shields.io/winget/v/LBRY.CLI)
![Flathub Version](https://img.shields.io/flathub/v/org.lbry.cli?logo=flathub)
![Homebrew Formula Version](https://img.shields.io/homebrew/v/lbry-cli?logo=homebrew)
![Chocolatey Version](https://img.shields.io/chocolatey/v/lbry-cli?logo=chocolatey)
![Scoop Version](https://img.shields.io/scoop/v/lbry-cli)
![Conda Version](https://img.shields.io/conda/vn/conda-forge/lbry-cli?logo=anaconda)

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

### Flatpak (using Flathub)

```shell
flatpak install flathub org.lbry.cli
```

## Run

```shell
lbry-cli version
```

## License
This project is MIT licensed. For the full license, see [LICENSE](LICENSE.md).
