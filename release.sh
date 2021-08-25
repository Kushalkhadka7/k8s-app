#!/bin/bash

set -e

green="\033[32m"
reset="\033[0m"

apps=$(git diff --name-only | sort -u | grep -oP "apps\/.+?\/" | cat | uniq)


if (( ${#apps[@]} )); then
  printf "Publishing images...\n"

  # for app in $apps; do
  #   printf "$green> $app\n"
  # done

  APP_DOCKER_REGISTRY="crkushal7"

  for app in $apps; do
    app_name=${app/apps\//}
    app_name=${app_name::-1}

    git fetch --tags
    last_image_tag=$(git tag --sort=-version:refname | grep -P "^$app_name@\d+.\d+.\d+$" | head -n 1)

    last_version=${last_image_tag#*@}

    printf "$last_version"

    if [ -z "$last_version" ]; then
      new_version="1.0.0"
    else
      new_version=$(semver bump release "$last_version")
      new_version=$(semver bump patch "$last_version")
    fi

    github_tag="$app_name@$new_version"
    image_tag="$APP_DOCKER_REGISTRY/$app_name:$new_version"

    printf "$green> Building image $image_tag\n"
    docker build -t $image_tag --target dev $app

    git tag -a $github_tag -m "my version 1.4"

    git push origin --tags

    docker login -u "crkushal7" -p "d45444eb-d5d9-4025-a482-d7482517e55a"

    docker push $image_tag
  done

  git add apps

  git commit -m "Initial"

  git push origin master
else
  printf "Skipping no changes detected...\n"
fi
