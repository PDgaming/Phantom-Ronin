#!/bin/bash

echo "Building..."
go build -v -ldflags '-s -w' -o ./build/Phantom_Ronin
echo "Build complete. Output file: Phantom_Ronin"