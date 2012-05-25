#! /bin/bash

# Builder command to update README, Git and PKGBUILD.
# Don't upload automatically built packages without checking them.

echo "$(./brocket -L)" > ./README
grep -A7 "# Information" ./brocket > ./tmp
. ./tmp
rm ./tmp

oldversion=$(\packer -Si brocket | grep 'Version' | cut -c 18-20)
if [[ $version == $oldversion ]]; then
    pkgrel=$(($(\packer -Si brocket | grep 'Version' | cut -c 22)+1))
else
    pkgrel=1
fi

md5sum=$(md5sum ./brocket | cut -c 1-32)

echo "# Maintainer: $maintainer <$author at gmail dot com>
pkgname=$name
pkgver=$version
pkgrel=$pkgrel
pkgdesc='A launcher for X11 WMs that attempts to prevent multiple instances.'
arch=(any)
url=$site
license=('GPL')
groups=()
depends=('bash' 'wmctrl')
optdepends=()
provides=()
conflicts=()
replaces=()
backup=()
options=()
install=
source=(brocket)
md5sums=('$md5sum')

build() {
  install -Dm 755 $srcdir/brocket ${pkgdir}/usr/bin/brocket
}" > ./PKGBUILD

echo "##### namcap results:"
echo "$(namcap ./PKGBUILD)"
echo
echo "##### makepkg --source results:"
$(makepkg -f --source)
echo
echo "##### Git results:"
$(git add PKGBUILD README brocket builder)
echo "$(git status)"
echo
echo "You can now 'git commit' and 'git push -u origin master'"
echo "and 'aurup ./$name-$version-$pkgrel.src.tar.gz'"