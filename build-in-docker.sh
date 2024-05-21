#!/bin/bash
set -e
git pull
if [[ $GIT_REF != "" ]]; then
  git checkout --detach $GIT_REF
fi
make clean
make
echo 'Completed. Please extract /app/aws-sdk-client-go from this container!'
echo 'For example, run the following command:'
echo 'docker cp $(docker ps -lq):/app/aws-sdk-client-go .'
echo ''
