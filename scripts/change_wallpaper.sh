#!/usr/bin/env bash
walls=$PICTURE_DIR
wall=$(ls -1 "$walls" | sort --random-sort | head -1)
osascript -e 'tell application "Finder" to set desktop picture to POSIX file "'"$walls/$wall"'"'
