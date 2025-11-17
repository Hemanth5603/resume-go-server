#!/bin/bash

# Start Local MongoDB for Development
echo "🚀 Starting Local MongoDB with Docker..."

# Check if Docker is running
if ! docker info > /dev/null 2>&1; then
    echo "❌ Docker is not running. Please start Docker first."
    exit 1
fi

# Start MongoDB container
docker run -d \
  --name resume-mongo-local \
  -p 27017:27017 \
  -v resume-mongo-data:/data/db \
  mongo:7.0

# Wait for MongoDB to be ready
echo "⏳ Waiting for MongoDB to be ready..."
sleep 5

# Check if MongoDB is running
if docker ps | grep -q resume-mongo-local; then
    echo "✅ MongoDB is running on mongodb://localhost:27017"
    echo ""
    echo "📝 Next steps:"
    echo "   1. Copy .env.local to .env:"
    echo "      cp .env.local .env"
    echo ""
    echo "   2. Run the server:"
    echo "      go run cmd/server/main.go"
    echo ""
    echo "   3. Test the API:"
    echo "      curl http://localhost:3000/health"
else
    echo "❌ Failed to start MongoDB"
    exit 1
fi

