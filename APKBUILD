makedepends="go"
license="MIT"
arch="all"
url=https://lbry.org
pkgdesc="The LBRY Command Line Interface (CLI) to interact with the LBRY Daemon."
pkgver=0.0.0
pkgrel=0
pkgname=lbry-cli
source="$pkgname-$pkgver.tar.gz::https://github.com/LBRYFoundation/lbry-cli/archive/refs/heads/master.tar.gz"

options="net"
builddir="$srcdir/lbry-cli-master"
maintainer="LBRY Foundation <board@lbry.org>"

build() {
	go build -o lbry-cli
}

check() {
	go test ./...
}

package() {
	install -Dm755 "$builddir/lbry-cli" "$pkgdir/usr/bin/lbry-cli"
}

sha512sums="
d763830597c22719b782c1649c1e879502c0d3acd26896393b3ef51fad4e071734928f7942dae5644fe4e53a8f00044fda53486abd635eeb931dd59ba68e1e2d  lbry-cli-0.0.0.tar.gz
"
