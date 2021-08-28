#!/bin/bash
function hello() {
    echo "Hello $1" # arguments are accessible through $1, $2,...
}

hello "$@"
