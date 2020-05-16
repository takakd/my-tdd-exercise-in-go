#!/usr/bin/env bash

if [[ $# -le 0 ]]; then
    echo "usage: sh commit (green|red) comment"
    exit 0
fi

if [[ "$1" = "red" ]] || [[ "$1" = "green" ]]; then
#    REV=$(date +%Y%m%d)
#    echo "* master" | grep "\* master"
#    if [[ $? -ne 0 ]]; then
#        git checkout -b ${REV}
#    fi

    git add --all

    if [[ "$2" = "" ]]; then
        git commit -m "${1}"
    else
        git commit -m "${1} note:${2}"
    fi
fi