#!/bin/bash

set -e

cur=$(pwd);
for f in $(find . -type d ! -path "./example*" ! -path "./releases*" ! -path "./vendor*" ! -path "./.git*" ! -path "./.idea*"); do
    printf "\033[1m## $cur/$f\033[0m\n"
    cd "$cur/$f"
    go test "$@"
    echo
done

printf "\033[1m## All OK\033[0m\n"
