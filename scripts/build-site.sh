#!/usr/bin/env bash
set -euo pipefail

REPO_ROOT=$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)
HUGO_BIN=${HUGO_BIN:-hugo}
BUILD_OUTPUT=${BUILD_OUTPUT:-"$REPO_ROOT/public"}
MODULE_PROXY=${MODULE_PROXY:-https://proxy.golang.org}

BUILD_SCRATCH=
MODULE_SCRATCH=
if [[ -z ${BUILD_ROOT:-} ]]; then
  BUILD_SCRATCH=$(mktemp -d)
  BUILD_ROOT="$BUILD_SCRATCH/website-build"
fi
if [[ -z ${MODULE_ROOT:-} ]]; then
  MODULE_SCRATCH=$(mktemp -d)
  MODULE_ROOT="$MODULE_SCRATCH/hugo-modules"
fi

cleanup() {
  if [[ ${KEEP_BUILD_ROOT:-false} != true ]]; then
    [[ -z $BUILD_SCRATCH ]] || rm -rf "$BUILD_SCRATCH"
    [[ -z $MODULE_SCRATCH ]] || rm -rf "$MODULE_SCRATCH"
  fi
}
trap cleanup EXIT

prepare_module() {
  local module_path=$1
  local version=$2
  local target=$3
  local archive="$MODULE_ROOT/$(basename "$target").zip"
  local unpack="$MODULE_ROOT/$(basename "$target")-unpack"

  mkdir -p "$target" "$unpack"
  curl -fsSL --retry 3 "$MODULE_PROXY/$module_path/@v/$version.zip" -o "$archive"
  unzip -q -o "$archive" -d "$unpack"
  local source
  source=$(find "$unpack" -type f -name go.mod -print -quit)
  test -n "$source"
  cp -a "$(dirname "$source")"/. "$target"/
}

rm -rf "$BUILD_ROOT" "$BUILD_OUTPUT"
mkdir -p "$BUILD_ROOT" "$BUILD_OUTPUT" "$MODULE_ROOT"
tar --exclude=.git --exclude=public -C "$REPO_ROOT" -cf - . | tar -C "$BUILD_ROOT" -xf -

prepare_module github.com/gzu-ai/hugo-blox-builder/modules/blox-core v0.3.1 "$MODULE_ROOT/core"
prepare_module github.com/gzu-ai/hugo-blox-builder/modules/blox-seo v0.2.3 "$MODULE_ROOT/seo"

sed -i.bak '1c\
module github.com/gzu-ai/hugo-blox-builder/modules/blox-core' "$MODULE_ROOT/core/go.mod"
sed -i.bak '1c\
module github.com/gzu-ai/hugo-blox-builder/modules/blox-seo' "$MODULE_ROOT/seo/go.mod"
sed -i.bak 's#github.com/gui-ai/hugo-blox-builder/modules/blox-core#github.com/gzu-ai/hugo-blox-builder/modules/blox-core#g' "$MODULE_ROOT/seo/go.mod"
rm -f "$MODULE_ROOT/core/go.mod.bak" "$MODULE_ROOT/seo/go.mod.bak"

cd "$BUILD_ROOT"
go mod edit -replace=github.com/gzu-ai/hugo-blox-builder/modules/blox-core="$MODULE_ROOT/core"
go mod edit -replace=github.com/gzu-ai/hugo-blox-builder/modules/blox-seo="$MODULE_ROOT/seo"

HUGO_MODULE_PROXY="$MODULE_PROXY,direct" \
GOPROXY="$MODULE_PROXY,direct" \
"$HUGO_BIN" --minify --destination "$BUILD_OUTPUT"

test -f "$BUILD_OUTPUT/zh/index.html"
test -f "$BUILD_OUTPUT/en/index.html"
test -f "$BUILD_OUTPUT/admin/index.html"
