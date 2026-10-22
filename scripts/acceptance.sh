#!/usr/bin/env bash
set -euo pipefail

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$ROOT"
version="$(tr -d '[:space:]' < VERSION)"

echo "1/8 format"
files=$(gofmt -l .)
if [ -n "$files" ]; then printf 'Go files requiring gofmt:\n%s\n' "$files"; exit 1; fi

echo "2/8 vet"
go vet ./...

echo "3/8 race tests"
go test -race -count=1 ./...

echo "4/8 dependency boundary"
modules=$(go list -m all | wc -l | tr -d ' ')
test "$modules" = "1"

echo "5/8 build/version"
mkdir -p bin
go build -o ./bin/app-shell-demo ./cmd/app-shell-demo
test "$(./bin/app-shell-demo -version)" = "app-shell-demo ${version}"

echo "6/8 asset budget"
asset_bytes=$(($(wc -c < shell/assets/shell.css) + $(wc -c < shell/assets/pwa.js)))
echo "first-party CSS + optional JS: ${asset_bytes} bytes"
test "$asset_bytes" -le 32768

echo "7/8 neutral HTTP acceptance"
port=18080
./bin/app-shell-demo -addr "127.0.0.1:${port}" >/tmp/go-app-shell-demo.log 2>&1 &
pid=$!
trap 'kill "$pid" 2>/dev/null || true' EXIT
for _ in $(seq 1 50); do if curl -fsS "http://127.0.0.1:${port}/" >/tmp/gas-home.html; then break; fi; sleep .1; done
curl -fsS "http://127.0.0.1:${port}/records/alpha" >/tmp/gas-detail.html
curl -fsS "http://127.0.0.1:${port}/manifest.webmanifest" >/tmp/gas-manifest.json
curl -fsS "http://127.0.0.1:${port}/service-worker.js" >/tmp/gas-sw.js
grep -q '<main id="main-content"' /tmp/gas-home.html
grep -q 'aria-current="page"' /tmp/gas-home.html
grep -q 'data-compact-pane="detail"' /tmp/gas-detail.html
grep -q '"name": "App Shell Demo"' /tmp/gas-manifest.json
grep -q 'event.request.method !== "GET"' /tmp/gas-sw.js

echo "8/8 render benchmark"
go test -run '^$' -bench BenchmarkRender -benchmem ./shell

echo "PASS"
