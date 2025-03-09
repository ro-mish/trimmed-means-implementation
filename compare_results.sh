#!/bin/bash

echo "=== Running Go Implementation ==="
echo "--------------------------------"
go run main.go
echo
echo "=== Running R Implementation ==="
echo "--------------------------------"
Rscript compare_means.R 