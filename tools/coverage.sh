#!/bin/bash
set -x

COVERAGE_DIR="coverage"

PKG_LIST=$(go list ./... 2>/dev/null | grep -v /vendor/)

# Create the coverage files directory
mkdir -p "$COVERAGE_DIR"

# Create a coverage file for each package
for package in ${PKG_LIST}; do
    go test -covermode=count -coverprofile "${COVERAGE_DIR}/${package##*/}.cov" "$package"
done

# Merge the coverage profile files
echo 'mode: count' > "${COVERAGE_DIR}"/coverage.all
tail -q -n +2 "${COVERAGE_DIR}"/*.cov >> "${COVERAGE_DIR}"/coverage.all

# Display the global code coverage
go tool cover -func="${COVERAGE_DIR}"/coverage.all

# If needed, generate HTML report
if [ "$1" == "html" ]; then
    go tool cover -html="${COVERAGE_DIR}"/coverage.all -o "${COVERAGE_DIR}"/coverage.html
fi

# Remove the coverage files directory
# rm -rf "$COVERAGE_DIR";