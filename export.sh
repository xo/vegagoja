#!/bin/bash

SRC=$(realpath $(cd -P "$(dirname "${BASH_SOURCE[0]}")" && pwd))

for f in $(find $SRC/testdata -type f -iname \*.svg|sort -h); do
  echo "EXPORTING ${f#$SRC/} -> ${f#$SRC/}.export.png"
  inkscape \
    --export-area-page \
    --export-background='#ffffff' \
    --export-type=png \
    -o "${f}.export.png" \
   "$f"
done

