#!/bin/bash

echo "Building..."
go build -v -ldflags '-s -w'
echo "Build complete. Output file: Phantom_Ronin"