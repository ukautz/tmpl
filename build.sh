#!/bin/sh

set -e

if [ "X$ISBUILD" != "X1" ]; then
    echo "This script is supposed to be called during build within docker"
    exit 1
fi

export PATH="$PATH:${GOROOT}/bin"
RELEASES=${RELEASES:-/releases}
SOURCE=${SOURCE:-/source}

NEWUID=$1
PLATFORMS='[["darwin", ["amd64"]], ["linux", ["amd64", "386"]], ["windows", ["amd64"]]]'
#PLATFORMS='[["linux", ["amd64"]]]'

cd $GOPATH/src/github.com/ukautz/tmpl
for platform in $(echo "$PLATFORMS" | jq -cr '.[]'); do
    export GOOS=$(echo "$platform" | jq -cr '.[0]')
    echo " > $GOOS"
    for arch in $(echo "$platform" | jq -cr '.[1] | .[]'); do
        export GOARCH=$arch
        echo " >> $arch"
        #set -x
        go build -ldflags "-s -w" -o "$RELEASES/tmpl.${GOOS}-${GOARCH}" main/main.go && \
            upx --brute "$RELEASES/tmpl.${GOOS}-${GOARCH}" && \
            chown $NEWUID "$RELEASES/tmpl.${GOOS}-${GOARCH}"
    done
done



