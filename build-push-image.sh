#!/bin/bash
SHORT_SHA=${TRAVIS_COMMIT} | awk '{print substr($0,0,7)}'
REPO=optikon/api
echo ${TRAVIS_COMMIT}
echo "${TRAVIS_COMMIT}"

# If the tag is undefined
if [ -z $TRAVIS_TAG ]; then
  echo "NOT A TAG"
  echo ${SHORT_SHA}
  echo ${REPO}
  # docker push $REPO:$SHORT_SHA
  # docker tag $REPO:$SHORT_SHA $REPO:latest
  # docker push $REPO:latest
# If the tag is set
else
  echo "THIS IS A TAG"
  echo ${SHORT_SHA}
  echo ${REPO}
  docker tag $REPO:$SHORT_SHA $REPO:$TRAVIS_TAG
  # docker push $REPO:$TRAVIS_TAG
fi
