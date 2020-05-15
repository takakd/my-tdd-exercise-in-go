#!/usr/bin/env bash

if [[ $# -le 0 ]]; then
    echo "usage: sh commit (green|red)"
    exit 0
fi

if [[ "$1" -eq "red" ]] || [[ "$1" -eq "green" ]]; then
#    REV=$(date +%Y%m%d)
#    echo "* master" | grep "\* master"
#    if [[ $? -ne 0 ]]; then
#        git checkout -b ${REV}
#    fi

    git add --all
    git commit -m "${1}"
fi