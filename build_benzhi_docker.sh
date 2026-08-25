#!/usr/bin/env sh
set -e
docker build -f benzhi.Dockerfile -t bagsort:local .
