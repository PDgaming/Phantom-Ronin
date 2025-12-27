#!/bin/bash

echo "Building..."
go build -v -ldflags '-s -w' -o ./dist/bin/Phantom_Ronin
echo "Build complete. Output file: Phantom_Ronin"