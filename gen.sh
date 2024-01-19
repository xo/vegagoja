#!/bin/bash

# github repos
REPO_VG=vega/vega
REPO_VL=vega/vega-lite

SRC=$(realpath $(cd -P "$(dirname "${BASH_SOURCE[0]}")" && pwd))

set -e

WORKDIR=$HOME/src/charts

mkdir -p $WORKDIR

BABEL_CONFIG=$(cat << __END__
{
  edge: '17',
  firefox: '60',
  chrome: '67',
  safari: '11.1'
},
useBuiltIns: 'entry',
corejs: { version: 3, proposals: false }
__END__
)

git_latest_tag() {
  git -C "$1" describe --abbrev=0 --tags
}

git_checkout_reset() {
  local dir="$WORKDIR/$1" name="$1" repo="$2"
  if [ ! -d "$dir" ]; then
    (set -x;
      git clone "$repo" "$dir"
    )
  fi
  (set -x;
    git -C "$dir" fetch origin
  )
  local ver=$(git_latest_tag "$dir")
  echo "$name $ver"
  echo "$ver" > "$SRC/${name}-version.txt"
  (set -x;
    git -C "$dir" reset --hard
    git -C "$dir" clean -f -x -d -e node_modules
    git -C "$dir" checkout "$ver" &> /dev/null
  )
}

git_checkout_reset vega      "https://github.com/${REPO_VG}.git"
git_checkout_reset vega-lite "https://github.com/${REPO_VL}.git"

pushd $WORKDIR/vega &> /dev/null
(set -x;
  yarn install
  yarn build
)
popd &> /dev/null

pushd $WORKDIR/vega-lite &> /dev/null

perl -0pi -e "s%targets:\s*{[^}]+}%targets: ${BABEL_CONFIG}%gms" babel.config.js
sed -i "1s%^%import structuredClone from '@ungap/structured-clone';\\n\\n%" src/util.ts

(set -x;
  yarn install
  yarn add -D \
    @babel/cli \
    @ungap/structured-clone \
    @types/ungap__structured-clone
  yarn build
  yarn run \
    babel \
    --config-file=./babel.config.js \
    --retain-lines \
    --out-dir=./out \
    ./build
)
popd &> /dev/null

(set -x;
  cp $WORKDIR/vega/packages/vega/build/vega.min.js $SRC
  cp $WORKDIR/vega-lite/out/vega-lite.min.js $SRC
)
