#!/usr/bin/env bash

set -euo pipefail

godoc2md github.com/andy2046/tik \
  > $GOPATH/src/github.com/andy2046/tik/doc.md
