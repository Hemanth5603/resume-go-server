#!/bin/bash
# Deploy to Google Cloud Run

set -e

PROJECT_ID="${GCP_PROJECT_ID:-your-project-id}"
SERVICE_NAME="${SERVICE_NAME:-resume-server}"
REGION="${REGION:-us-central1}"

echo "==================================="
echo "Deploying to Google Cloud Run"
echo "==================================="
echo "Project: $PROJECT_ID"
echo "Service: $SERVICE_NAME"
echo "Region: $REGION"
echo "==================================="

# Check if gcloud is installed
if ! command -v gcloud &> /dev/null; then
    echo "Error: gcloud CLI is not installed"
    exit 1
fi

# Set project
gcloud config set project $PROJECT_ID

# Check if .env file exists
if [ ! -f ".env" ]; then
    echo "Error: .env file not found"
    echo "Please create .env file with required environment variables"
    exit 1
fi

# Load environment variables
export $(cat .env | grep -v '^#' | xargs)

# Deploy to Cloud Run
echo "Deploying to Cloud Run..."
gcloud run deploy $SERVICE_NAME \
    --source . \
    --platform managed \
    --region $REGION \
    --allow-unauthenticated \
    --set-env-vars PORT=3000,DATABASE_URL="$DATABASE_URL",JWKS_URL="$JWKS_URL" \
    --memory 512Mi \
    --cpu 1 \
    --min-instances 0 \
    --max-instances 10 \
    --timeout 300 \
    --quiet

# Get service URL
SERVICE_URL=$(gcloud run services describe $SERVICE_NAME \
    --region $REGION \
    --format 'value(status.url)')

echo ""
echo "==================================="
echo "Deployment Complete!"
echo "==================================="
echo "Service URL: $SERVICE_URL"
echo "Health Check: $SERVICE_URL/health"
echo ""
echo "View logs: gcloud run logs tail $SERVICE_NAME --region $REGION"
echo "==================================="

