#!/bin/bash
set -e

echo "Setting up test branches"

# Ensure we start from master
echo "Checking out master"
git checkout master &> /dev/null

# Setup branch that is only local
echo "Creating local only branch"
git branch localOnly &> /dev/null

# Setup branch with remote existing
echo "Creating branch with remote existing"
git checkout -b remoteExisting &> /dev/null
git push -u origin remoteExisting &> /dev/null
git checkout master &> /dev/null

# Setup branch that has remote gone
echo "Creating branch with remote gone"
git checkout -b remoteDeleted &> /dev/null
git push -u origin remoteDeleted &> /dev/null
git push -d origin remoteDeleted &> /dev/null
git checkout master &> /dev/null

# Setup branch that gives merged warning
echo "Creating branch with merge warning"
git checkout -b mergeWarning &> /dev/null
git commit -m "Test" --allow-empty &> /dev/null
git checkout master &> /dev/null

# Return to master
echo "Returning to master"
git checkout master &> /dev/null