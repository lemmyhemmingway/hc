#!/bin/bash
set -e

ROOT="$(cd "$(dirname "$0")/.." && pwd)"
cd "$ROOT"

cd ui
npm install
npm start &
ANG_PID=$!
cd ..

go run main.go &
GO_PID=$!

cleanup() {
  kill $ANG_PID $GO_PID
}
trap cleanup EXIT

wait $GO_PID
