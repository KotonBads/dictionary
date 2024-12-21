#!/usr/bin/bash

for dir in cmd/*; do
    if [ -d "$dir" ]; then
        for file in "$dir"/*; do
            if [ -f "$file" ]; then
                echo "Building $file"
                base=$(basename "${file%.*}")
                go build -o "bin/$base" "$file"
            fi
        done
    fi
done
