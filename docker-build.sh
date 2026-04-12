#!/bin/bash
# Builds the UI and embeds it before docker build
cd ui
npm install
npm run build
cd ..
mkdir -p internal/managementasset/dist
cp ui/dist/index.html internal/managementasset/dist/index.html
docker build -t cliproxyapiplus .
