#!/usr/bin/env bash
readonly PORGNAME=$(basename $0)

usage = '$PORGNAME needs to have "$PICTURE_DIR" set to execute'
dir=$PICTURE_DIR
if [ ! -d $dir ]; then
    echo "$usage"
    exit 1
fi
file=$(ls -1 $dir/*.jpg  | sort --random-sort | head -1)
echo $file
sqlite3 ~/Library/Application\ Support/Dock/desktoppicture.db "update data set value = '${file}'" && killall Dock

