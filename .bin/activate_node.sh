#!/bin/sh

main() {
  local release="10.3.0"
  local name="node-v$release-linux-x64"
  local env="$PWD/.vendor/node"

  if [ "$NODEJS_HOME" != "$env" ]; then
    if [ ! -d "$env" ]; then
      local package="$name.tar.xz"

      wget -c "https://nodejs.org/dist/v$release/$package" -O "/tmp/$package" &&
      tar -xf "/tmp/$package" -C .vendor/ &&
      mv ".vendor/$name" "$env"
    fi

    export NODEJS_HOME="$env"
    export PATH="$PWD/node_modules/.bin:$NODEJS_HOME/bin:$PATH"

    [ -d "node_modules" ] || npm install
  fi
}

main $@

