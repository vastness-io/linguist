#!/bin/bash
echo "Updating go port of Github's Linguist"
git clone https://github.com/github/linguist ./pkg/linguist/data/linguist
go generate ./pkg/linguist
go generate ./pkg/linguist/data
rm -rf ./pkg/linguist/data/linguist
echo 'Complete'