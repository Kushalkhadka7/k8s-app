#!/bin/bash

set -e

apps=$(git diff --name-only | sort -u | grep -oP "apps\/.+?\/" | cat | uniq)

APP_DOCKER_REGISTRY="crkushal7"

if (( ${#apps[@]} != 0 )); then
  printf "Publishing images...\n"

  for app in $apps; do
    app_name=${app/apps\//}
    app_name=${app_name::-1}

    git fetch --tags
    last_image_tag=$(git tag --sort=-version:refname | grep -P "^$app_name@\d+.\d+.\d+$" | head -n 1)

    last_version=${last_image_tag#*@}

    if [ -z "$last_version" ]; then
      new_app_version="1.0.0"
    else
      new_app_version=$(semver bump release "$last_version")
      new_app_version=$(semver bump patch "$last_version")
    fi

    github_tag="$app_name@$new_app_version"
    image_tag="$APP_DOCKER_REGISTRY/$app_name:$new_app_version"

    docker build -t $image_tag --target dev $app
  done
else
  printf "Skipping no changes detected...\n"
fi
