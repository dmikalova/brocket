#!/bin/sh
set -eu
if [ "${DEBUG:=}" = true ]; then set -x; fi

DIR="$(realpath "$(dirname "$0")")"

ln -sf "${DIR}/brocket" "${HOME}/.bin/"
for FILE in "${DIR}/examples/"*; do
    echo $FILE
    ln -sf "${FILE}" "${HOME}/.bin/"
done
