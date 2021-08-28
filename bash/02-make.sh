#!/bin/bash

build() {
    ...
}

test() {
    ...
}

deploy() {
    ...
}

# This lets you do ./run.sh build foo bar at the command
# line.  The autocomplete scans for functions in .sh files
# and fills them in as the first arg.

"$@"
