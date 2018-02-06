#!/bin/bash
echo "Updating go port of Github's Linguist"
git clone https://github.com/github/linguist ./third_party/linguist/data/linguist
go generate ./third_party/linguist
go generate ./third_party/linguist/data
rm -rf ./third_party/linguist/data/linguist
echo 'Complete'