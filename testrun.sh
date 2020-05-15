#!/usr/bin/env bash

CMD="go test ./tddexercise"
INTERVAL=3

watch -n ${INTERVAL} ${CMD}