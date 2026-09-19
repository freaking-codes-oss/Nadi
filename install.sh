#!/data/data/com.termux/files/usr/bin/sh
set -eu
command -v go >/dev/null 2>&1 || { echo 'Install Go first: pkg install golang'; exit 1; }
go build -o "$PREFIX/bin/nadi" ./cmd/nadi
chmod 700 "$PREFIX/bin/nadi"
echo 'Nadi installed to $PREFIX/bin/nadi'
