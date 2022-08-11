#!/bin/bash

gcloud functions deploy get-new-app \
    --region asia-northeast1 \
    --project daily-steam \
    --entry-point GetNewApp \
    --memory 256MB \
    --runtime go116 \
    --trigger-http \
    --allow-unauthenticated \
    --set-env-vars NEW_APP_SCRAPE_URL=$NEW_APP_SCRAPE_URL \
