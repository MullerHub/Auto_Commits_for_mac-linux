set -euo pipefail

SRC="main.go"
OUT_DIR="./bin"
WINDOWS_OUT="$OUT_DIR/Luna.exe"
LINUX_OUT="$OUT_DIR/Luna_linux_amd64"
MAC_OUT="$OUT_DIR/luna-mac"

mkdir -p "$OUT_DIR"

echo "Building linux/amd64..."
GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o "$LINUX_OUT" "$SRC"

echo "Building windows/amd64..."
GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o "$WINDOWS_OUT" "$SRC"

echo "Buildando para Mac (Apple Silicon M1/M2)..."
# GOOS=darwin (Mac) e GOARCH=arm64 (Chip M1/M2/M3)
GOOS=darwin GOARCH=arm64 go build -ldflags "-s -w" -o "$MAC_OUT" "$SRC"

echo "✅ Sucesso! O executável está em: $MAC_OUT"
echo "Para rodar, use: ./bin/luna-mac"

echo "Build success. Artifacts:"
ls -lh "$OUT_DIR"
