#!/usr/bin/env bash

# ensure gcloud set as cred helper for docker first before running

set -ev

SCRIPT_DIR=$(dirname "$0")

gcloud auth print-access-token | docker login -u oauth2accesstoken --password-stdin https://us-docker.pkg.dev
docker build -f docker/payment/Dockerfile . -t runwhendemo/payment:latest -t us-docker.pkg.dev/runwhen-nonprod-shared/public-images/runwhendemo/payment:latest
docker push us-docker.pkg.dev/runwhen-nonprod-shared/public-images/runwhendemo/payment:latest