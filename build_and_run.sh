#!/bin/bash
# This script builds the Go project and runs the program.

go build -o program main.go
if [ $? -eq 0 ]; then
    echo "Build successful. Running the program..."
    ./program
else
    echo "Build failed. Please check the errors above."
fi
# Clean up the built program
rm -f program
echo "Cleaned up the built program."
# End of script
# Usage: ./build_and_run.sh
# Make sure to give execute permission to the script before running it:
# chmod +x build_and_run.sh