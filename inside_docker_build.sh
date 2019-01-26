#!/bin/bash
for GOOS in darwin linux; do go build -v -o pushover-$GOOS-amd64; done
