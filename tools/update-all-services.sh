#!/bin/bash

set -ex
git clone --depth=1 https://github.com/aws/aws-sdk-go-v2.git
cd aws-sdk-go-v2/service
echo "services:" > ../../all-services.yaml
ls | egrep -v '^internal$' | perl -pE 's/^/  /;s/$/:/' >> ../../all-services.yaml
