#!/bin/sh

main() {
  local release=${1:-10.1.0}
  local env="node-v${release}-linux-x64"

  if [ "${NODEJS_HOME}" = "${PWD}/.vendor/${release}" ]; then
    echo "Node.js is active"
  else
    if [ ! -d ".vendor/${env}" ]; then
      echo "Installing Node.js ${release}.."
      echo

      local package="${env}.tar.xz"

      wget -c "https://nodejs.org/dist/v${release}/${package}" -O "/tmp/${package}" &&
      tar -xf "/tmp/${package}" -C .vendor/

      echo
      echo "Done"
    fi

    NODEJS_HOME="${PWD}/.vendor/${env}"
    PATH="${PWD}/node_modules/.bin:${NODEJS_HOME}/bin:${PATH}"

    echo "Node.js (${release}) activated"
  fi
}

main $@

