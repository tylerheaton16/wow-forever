#!/usr/bin/env bash
# Builds the double-click launchers into ../dist (always embeds the current page).
set -euo pipefail
cd "$(dirname "$0")"
cp ../index.html .
mkdir -p ../dist
GOOS=windows GOARCH=amd64 go build -ldflags "-H windowsgui -s -w" -o ../dist/waylaid-crates.exe .
GOOS=darwin  GOARCH=arm64 go build -ldflags "-s -w" -o ../dist/waylaid-crates-mac .
echo "Built dist/waylaid-crates.exe and dist/waylaid-crates-mac"
