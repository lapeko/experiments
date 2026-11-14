#!/bin/bash

NAME="$1"

if [[ -z "${NAME}" ]]; then
    echo "empty name provided" >&2
    exit 1
fi

echo "Hello, $NAME!"
