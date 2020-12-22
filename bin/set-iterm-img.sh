#!/usr/bin/env bash
echo "set-iterm-img.sh"
echo $1
osascript bgImgIterm.scpt "$1"
