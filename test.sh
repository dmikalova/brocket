#!/bin/sh
set -eu

cd "$(realpath "$(dirname "${0}")")"
go install

# /home/dmikalova/go/bin/brocket run-or-raise --class "firefox"
# /home/dmikalova/go/bin/brocket run-or-raise --class "dolphin"
# /home/dmikalova/go/bin/brocket run-or-raise --class "konsole"
# /home/dmikalova/go/bin/brocket run-or-raise --class "code" --list true
# /home/dmikalova/go/bin/brocket run-or-raise --class "bitwarden" --cmd "/snap/bitwarden/current/bitwarden"

# time sh -c 'for i in $(seq 100); do
#     /home/dmikalova/go/bin/brocket run-or-raise --class "konsole"
#     /home/dmikalova/go/bin/brocket run-or-raise --class "code"
# done'

# time sh -c 'for i in $(seq 100); do
#     /home/dmikalova/.bin/brocket.sh "konsole"
#     /home/dmikalova/.bin/brocket.sh "code"
# done'
