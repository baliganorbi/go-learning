#!/bin/bash
# This script verifies the provenance of a Go project build artifact using SLSA Verifier.
# Make sure to have slsa-verifier installed and configured properly.

# Usage: ./verify-provenance.sh source-tag

# Check if the source-tag is provided
if [ -z "$1" ]; then
    echo "Usage: $0 <source-tag>"
    exit 1
fi
# Set the source tag from the command line argument
SOURCE_TAG=$1

set -e
# Check if slsa-verifier is installed
if ! command -v slsa-verifier &> /dev/null; then
    echo "slsa-verifier could not be found. Please install it first."
    exit 1
fi
# Check if the provenance file exists
if [ ! -f program-linux-amd64.intoto.jsonl ]; then
    echo "Provenance file program-linux-amd64.intoto.jsonl does not exist."
    exit 1
fi
# Check if the artifact exists
if [ ! -f program-linux-amd64 ]; then
    echo "Artifact program-linux-amd64 does not exist."
    exit 1
fi
# Verify the artifact using SLSA Verifier
echo "Starting verification of the artifact..."

slsa-verifier verify-artifact program-linux-amd64 \
  --provenance-path program-linux-amd64.intoto.jsonl \
  --source-uri github.com/baliganorbi/go-learning \
  --source-tag $SOURCE_TAG