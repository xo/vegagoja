#/bin/bash

SRC=$(realpath $(cd -P "$(dirname "${BASH_SOURCE[0]}")" && pwd))

pushd $SRC/testdata &> /dev/null

for i in $(ls */*.json|sort -h); do
  if [ ! -f $i.svg ]; then
    echo "TestRender/$(sed -e 's/\.v[lg]\.json$//' <<< "$i")"
  fi
done

popd &> /dev/null
