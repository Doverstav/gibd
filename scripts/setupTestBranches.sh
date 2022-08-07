#!/bin/bash

echo "Setting up test branches"

# Setup branch that is only local
git branch localOnly

# Setup branch with remote existing
git branch remoteExisting
git push -u origin remoteExisting

# Setup branch that has remote gone
git branch remoteDeleted
git push -u origin remoteDeleted
git push -d origin remoteDeleted

# Setup branch that gives merged warning
git checkout -b mergeWarning
git commit -m "Test" --allow-empty

# Return to master
git checkout master