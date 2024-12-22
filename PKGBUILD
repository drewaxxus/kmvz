# Maintainer: Andrew P <andreipiscoinicholas@gmail.com>
# License: MIT
# URL: https://github.com/drewaxxus/kmvz

pkgname=kmvz
pkgver=0.8
pkgrel=1
arch=('x86_64')
depends=('go')
makedepends=('git')
source=('https://github.com/drewaxxus/kmvz.git')
sha256sums=('2be52d4c4cba3fc64f41aa4475b4c5d5eb320b5229fb6d682367856870a66892')

prepare() {
  cd "${srcdir}"
  rm -rf "${pkgname}" # Remove the existing directory
  git clone "${source[0]}"
  cd "${pkgname}"
  echo '{ "aur_api_url": "https://aur.archlinux.org/rpc/" }' > config.json
}

build() {
  cd "${pkgname}"
  GOOS=linux GOARCH=amd64 go build -o "${pkgdir}/${pkgname}" ./cmd/kmvz 
}

package() {
  cd "${pkgdir}"
  install -Dm644 "${pkgname}" "${pkgdir}/${pkgname}"
  install -Dm644 "${pkgname}/config.json" "${pkgdir}/${pkgname}/config.json"
}
