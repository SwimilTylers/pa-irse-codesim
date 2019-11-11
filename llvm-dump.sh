#!/bin/bash

temp=$(mktemp -t)
clang -O3 -emit-llvm -S "$1" -o "$temp"
cat "$temp"
