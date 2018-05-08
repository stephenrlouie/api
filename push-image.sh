#!/bin/bash

REPO = optikon/api

# If the tag is undefined
if [-z $TRAVIS_TAG ]; then
  echo "NOT A TAG"
  # docker push $REPO:$SHORT_SHA
  # docker tag $REPO:$SHORT_SHA $REPO:latest
  # docker push $REPO:latest
# If the tag is set
else
  echo "THIS IS A TAG"
  docker tag $REPO:$SHORT_SHA $REPO:$TRAVIS_TAG
  # docker push $REPO:$TRAVIS_TAG
fi
