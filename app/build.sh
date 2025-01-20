#!/bin/bash

# Set Go project directory (change this to your project's directory if needed)
PROJECT_DIR=$(pwd)

# Output directory
OUTPUT_DIR="../builds"

# Make sure the output directory exists
echo "Creating output directory at $OUTPUT_DIR..."
mkdir -p $OUTPUT_DIR

# Clean up old builds
echo "Cleaning old builds..."
rm -f $OUTPUT_DIR/SkyeBrowser-linux $OUTPUT_DIR/SkyeBrowser-macos $OUTPUT_DIR/SkyeBrowser-windows.exe

# Build for Linux (x86_64)
echo "Building for Linux..."
GOOS=linux GOARCH=amd64 go build -o $OUTPUT_DIR/SkyeBrowser-linux main.go

# Make the Linux binary executable
chmod +x $OUTPUT_DIR/SkyeBrowser-linux

# Build for macOS (x86_64)
echo "Building for macOS..."
GOOS=darwin GOARCH=amd64 go build -o $OUTPUT_DIR/SkyeBrowser-macos main.go

# Build for Windows (x86_64)
echo "Building for Windows..."
GOOS=windows GOARCH=amd64 go build -o $OUTPUT_DIR/SkyeBrowser-windows.exe main.go

echo "Build completed!"

# Optionally, create a zip archive of the output files
# echo "Creating zip archive..."
# zip -r $OUTPUT_DIR/SkyeBrowser.zip $OUTPUT_DIR/*

# End of script
