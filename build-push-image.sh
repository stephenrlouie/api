#!/bin/bash
SHORT_SHA=`echo ${TRAVIS_COMMIT::7}`
REPO=optikon/api

echo "${DOCKER_PASS}" | docker login -u "${DOCKER_USER}" --password-stdin

# If the tag is undefined
if [ ! -z $TRAVIS_TAG ]; then
  docker tag $REPO:$SHORT_SHA $REPO:$TRAVIS_TAG
  docker push $REPO:$TRAVIS_TAG
# If the tag is set
elif [ "$TRAVIS_BRANCH" == "master" ]; then
  docker push $REPO:$SHORT_SHA
  docker tag $REPO:$SHORT_SHA $REPO:latest
  docker push $REPO:latest
fi
