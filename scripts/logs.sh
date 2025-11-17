#!/bin/bash
# View logs from deployed application

INSTANCE_NAME="${INSTANCE_NAME:-resume-server}"
ZONE="${ZONE:-us-central1-a}"

echo "Fetching logs from $INSTANCE_NAME..."

gcloud compute ssh $INSTANCE_NAME --zone=$ZONE --command="
    cd /opt/resume-server
    docker-compose logs -f --tail=100
"

