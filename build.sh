#!/bin/bash
set -euo pipefail

PROJECT_DIR="$(cd "$(dirname "$0")" && pwd)"
VERSION="${1:-$(date '+%Y%m%d')}"
DIST_DIR="$PROJECT_DIR/dist"
FISHMAN_NAME="fishman-${VERSION}"

echo "================================="
echo " FishMan Multi-Platform Build"
echo " Version: $VERSION"
echo "================================="
echo ""

# Clean old dist
rm -rf "$DIST_DIR"
mkdir -p "$DIST_DIR"

# ========== macOS ARM64 ==========
echo "[1/2] Building for macOS ARM64..."
CGO_ENABLED=1 GOOS=darwin GOARCH=arm64 \
  go build -trimpath \
    -ldflags="-s -w -X 'main.version=${VERSION}'" \
    -o "$DIST_DIR/${FISHMAN_NAME}-mac-arm" \
    "$PROJECT_DIR/main.go"
echo "  -> dist/${FISHMAN_NAME}-mac-arm  ($(ls -lh "$DIST_DIR/${FISHMAN_NAME}-mac-arm" | awk '{print $5}'))"

# ========== Windows AMD64 ==========
echo "[2/2] Building for Windows AMD64..."
CC=x86_64-w64-mingw32-gcc
CGO_ENABLED=1 GOOS=windows GOARCH=amd64 CC="$CC" \
  go build -trimpath \
    -ldflags="-s -w -H windowsgui -extldflags=-static -X 'main.version=${VERSION}'" \
    -o "$DIST_DIR/${FISHMAN_NAME}-x64.exe" \
    "$PROJECT_DIR/main.go"
echo "  -> dist/${FISHMAN_NAME}-x64.exe  ($(ls -lh "$DIST_DIR/${FISHMAN_NAME}-x64.exe" | awk '{print $5}'))"

echo ""
echo "================================="
echo " Done! Files in dist/"
echo "================================="
