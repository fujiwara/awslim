#!/bin/bash
set -e
git pull
if [[ $GIT_REF != "" ]]; then
  git checkout --detach $GIT_REF
fi
make clean
make
echo 'Completed. Please extract /app/awslim from this container!'
echo 'For example, run the following command:'
echo 'docker cp $(docker ps -lq):/app/awslim .'
echo ''
