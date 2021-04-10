#!/bin/bash

set -ex

pros=$(git diff README.md | grep "+|" | awk 'NR==1{printf("%s", $2)} NR>1{ printf(", %s", $2) }')
test -n "${pros}" && commit="Add ${pros} solutions."
echo "${commit}"

git add *
git commit -m "${commit}"

