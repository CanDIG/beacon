#!/bin/sh
set -e

cd "$(dirname "$0")"

git submodule update --init
if [ -d .git/modules/candig-server ]; then
    mv -T .git/modules/candig-server candig-server/.git
fi
git -C candig-server config core.worktree ../
