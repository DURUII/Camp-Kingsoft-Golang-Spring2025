#!/bin/bash

# Check if golangci-lint is installed
if ! command -v golangci-lint &> /dev/null; then
    echo "Installing golangci-lint..."
    curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v2.2.1
    export PATH=$PATH:$(go env GOPATH)/bin
fi

echo "Running golangci-lint on all Go modules..."

# Find all directories with go.mod and run lint
find . -name 'go.mod' -exec dirname {} \; | while read dir; do
    # Skip directories without Go files
    if [ ! "$(find "$dir" -name "*.go" -type f | head -n 1)" ]; then
        echo "SKIP: $dir (no Go files)"
        continue
    fi
    
    echo "=== Linting: $dir ==="
    if (cd "$dir" && golangci-lint run); then
        echo "PASS: $dir"
    else
        echo "FAIL: $dir"
        exit 1
    fi
done

echo "=== Lint check completed ===" 