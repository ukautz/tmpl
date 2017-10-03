#!/bin/bash

set -e

dir=$(realpath $(dirname $0))
[ ! -d "$dir/releases" ] && mkdir "$dir/releases"

version=$1
if [ "$version" == "" ]; then
    echo "Usage: $0 <version>"
    exit 1
elif [[ ! $version =~ ^[0-9]+\.[0-9]+\.[0-9]+$ ]]; then
    echo "Version must comply with '[0-9]+\.[0-9]+\.[0-9]+'"
    exit 2
fi

set +e
if ! git diff-index --quiet HEAD --; then
    echo "Uncommited changes found. Cannot release"
    exit 3
fi
set -e

function read_latest_version() {
    version=$(git tag -l | sort -V -r | head -n 1)
    [ "$version" == "" ] && version="0.0.0"
    echo $version
}

function version_gt() {
    test "$(printf '%s\n' "$@" | sort -V | head -n 1)" != "$1";
}

last_version=$(read_latest_version)
if version_gt $last_version $version; then
    echo "Last version $last_version is higher then provided version $version"
    exit 4
elif [ "$last_version" == "$version" ]; then
    echo "Version $version is current released"
    exit 5
fi

function gen_template() {
    echo $1 | tmpl -d - -D json -t "$dir/$2.tmpl" -o "$dir/$2"
}

git commit --allow-empty -m "Release $version"
git tag $version

build_time=$(date --rfc-3339=seconds -u)
build_commit=$(git rev-parse HEAD)
data='{"version":"'$version'","time":"'$build_time'","commit":"'$build_commit'"}'
gen_template "$data" "version.go"

#--network host \
docker build \
    --network host \
    -t build/github-ukautz-tmpl:latest \
    -f Dockerfile.release .

docker run --rm -ti \
    --name build-github-ukautz-tmpl \
    --volume $dir:/source \
    --volume $dir/releases:/releases \
    build/github-ukautz-tmpl:latest \
    /source/build.sh "$(id -u):$(id -g)"

git checkout version.go



