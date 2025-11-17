#!/bin/bash
# Deployment script for GCP VM

set -e

# Configuration
PROJECT_ID="${GCP_PROJECT_ID:-your-project-id}"
INSTANCE_NAME="${INSTANCE_NAME:-resume-server}"
ZONE="${ZONE:-us-central1-a}"
MACHINE_TYPE="${MACHINE_TYPE:-e2-medium}"

echo "==================================="
echo "Resume Go Server - GCP Deployment"
echo "==================================="
echo "Project: $PROJECT_ID"
echo "Instance: $INSTANCE_NAME"
echo "Zone: $ZONE"
echo "Machine Type: $MACHINE_TYPE"
echo "==================================="

# Check if gcloud is installed
if ! command -v gcloud &> /dev/null; then
    echo "Error: gcloud CLI is not installed"
    echo "Install from: https://cloud.google.com/sdk/docs/install"
    exit 1
fi

# Set project
echo "Setting GCP project..."
gcloud config set project $PROJECT_ID

# Create firewall rules
echo "Creating firewall rules..."
gcloud compute firewall-rules create allow-resume-server \
    --allow=tcp:3000,tcp:80,tcp:443 \
    --target-tags=http-server \
    --description="Allow traffic to Resume Server" \
    --quiet || echo "Firewall rule already exists"

# Create VM instance
echo "Creating VM instance..."
gcloud compute instances create $INSTANCE_NAME \
    --zone=$ZONE \
    --machine-type=$MACHINE_TYPE \
    --image-family=ubuntu-2204-lts \
    --image-project=ubuntu-os-cloud \
    --boot-disk-size=20GB \
    --boot-disk-type=pd-standard \
    --tags=http-server,https-server \
    --metadata-from-file=startup-script=scripts/startup-script.sh \
    --quiet || echo "Instance already exists"

# Wait for instance to be ready
echo "Waiting for instance to be ready..."
sleep 30

# Get external IP
EXTERNAL_IP=$(gcloud compute instances describe $INSTANCE_NAME \
    --zone=$ZONE \
    --format='get(networkInterfaces[0].accessConfigs[0].natIP)')

echo ""
echo "==================================="
echo "Deployment Complete!"
echo "==================================="
echo "Instance Name: $INSTANCE_NAME"
echo "External IP: $EXTERNAL_IP"
echo "Health Check: http://$EXTERNAL_IP:3000/health"
echo ""
echo "Next steps:"
echo "1. SSH into the instance: gcloud compute ssh $INSTANCE_NAME --zone=$ZONE"
echo "2. Configure .env file: sudo nano /opt/resume-server/.env"
echo "3. Restart services: cd /opt/resume-server && sudo docker-compose restart"
echo "==================================="

