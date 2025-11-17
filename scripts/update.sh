#!/bin/bash
# Update script for deployed application

set -e

INSTANCE_NAME="${INSTANCE_NAME:-resume-server}"
ZONE="${ZONE:-us-central1-a}"

echo "Updating Resume Go Server on $INSTANCE_NAME..."

# SSH and update
gcloud compute ssh $INSTANCE_NAME --zone=$ZONE --command="
    set -e
    cd /opt/resume-server
    echo 'Pulling latest changes...'
    git pull
    echo 'Rebuilding containers...'
    docker-compose down
    docker-compose up -d --build
    echo 'Update complete!'
    docker-compose ps
"

echo "Update completed successfully!"

