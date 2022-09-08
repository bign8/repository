#! /bin/bash -ex

# Upstream issue to track simplifying this
# https://github.com/golang/go/issues/50745

go test -v $(go list -f '{{.Dir}}/...' -m | xargs)
# go test ./...
# for d in impl/*/; do
#     pushd $d > /dev/null
#     go test ./...
#     popd > /dev/null
# done
