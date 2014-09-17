#!/bin/bash

set -x
set -e

ROOT="$(cd "$( dirname "${BASH_SOURCE[0]}" )/.." && pwd)"
cd $ROOT && go build -o ~/.bin/newp
