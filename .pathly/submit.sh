#!/usr/bin/env bash
set -euo pipefail

chmod u+x ./test.sh
./test.sh

git add -A
git commit -am"completed exercise"
git push