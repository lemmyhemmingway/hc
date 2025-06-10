#!/bin/bash
set -e

ROOT="$(cd "$(dirname "$0")/.." && pwd)"
cd "$ROOT"

cd ui
npm install
npm run build
cd ..

go build -o healthcheck
./healthcheck
