#!/bin/bash

# Build for Windows
export GOOS=windows
export GOARCH=amd64
export CGO_ENABLED=1
export CC=x86_64-w64-mingw32-gcc

echo "Building for Windows (amd64)..."
go build -v -ldflags '-s -w -H=windowsgui' -o Phantom_Ronin.exe
echo "Build complete. Output file: Phantom_Ronin.exe"

unset GOOS GOARCH CGO_ENABLED CC

# Build for Linux
echo "Building for Linux (amd64)..."
go build -v -ldflags '-s -w'
echo "Build complete. Output file: Phantom_Ronin"