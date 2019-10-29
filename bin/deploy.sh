#!/bin/sh

set -o errexit
set -o xtrace

if [ "$GITHUB_EVENT_NAME" = "release" ] ; then
  export PROJECT_ID=dappface-prd-v2
elif [ "$GITHUB_EVENT_NAME" = "push" ] || [ $(basename  "$GITHUB_REF") = 'master' ] ; then
    export PROJECT_ID=dappface-stg-v2
else
	export PROJECT_ID=dappface-dev
fi

APP_NAME=news-loader
IMAGE_SRC_PATH=gcr.io/"$PROJECT_ID"/"$APP_NAME"

gcloud auth configure-docker
docker build -t "$IMAGE_SRC_PATH":latest -t "$IMAGE_SRC_PATH":"$GITHUB_SHA" .
docker push "$IMAGE_SRC_PATH":latest
docker push "$IMAGE_SRC_PATH":"$GITHUB_SHA"

BERGLAS_PATH=berglas://"$PROJECT_ID"-berglas-secrets

gcloud beta run deploy "$APP_NAME" \
	--project "$PROJECT_ID" \
	--image "$IMAGE_SRC_PATH":latest \
	--platform managed \
	--allow-unauthenticated \
	--region us-east1 \
	--set-env-vars "\
PROJECT_ID="$PROJECT_ID",\
TWITTER_ACCESS_TOKEN="$BERGLAS_PATH"/twitter-access-token,\
TWITTER_ACCESS_TOKEN_SECRET="$BERGLAS_PATH"/twitter-access-token-secret,\
TWITTER_API_KEY="$BERGLAS_PATH"/twitter-api-key,\
TWITTER_API_SECRET="$BERGLAS_PATH"/twitter-api-secret
"
