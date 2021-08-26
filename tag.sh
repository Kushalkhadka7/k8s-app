#!/bin/bash

git fetch --tags
last_image_tag=$(git tag --sort=-version:refname | grep -P "^$app_name@\d+.\d+.\d+$" | head -n 1)

printf "$last_image_tag"
