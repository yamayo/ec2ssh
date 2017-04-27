#!/bin/sh

for OS in "linux" "darwin"; do
  for ARCH in "386" "amd64"; do
    VERSION="$(git describe --tags $1)"
    GOOS=$OS GOARCH=$ARCH go build -ldflags "-X main.version=$VERSION" -o ec2ssh
    ARCHIVE="ec2ssh-$VERSION-$OS-$ARCH.tar.gz"
    tar -czf $ARCHIVE ec2ssh
    echo $ARCHIVE
  done
done
