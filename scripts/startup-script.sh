#!/bin/bash
# GCP VM Startup Script for Resume Go Server

set -e

echo "Starting Resume Go Server setup..."

# Update system
apt-get update
apt-get install -y \
    apt-transport-https \
    ca-certificates \
    curl \
    gnupg \
    lsb-release \
    git

# Install Docker
if ! command -v docker &> /dev/null; then
    echo "Installing Docker..."
    curl -fsSL https://get.docker.com -o get-docker.sh
    sh get-docker.sh
    systemctl start docker
    systemctl enable docker
    rm get-docker.sh
fi

# Install Docker Compose
if ! command -v docker-compose &> /dev/null; then
    echo "Installing Docker Compose..."
    curl -L "https://github.com/docker/compose/releases/latest/download/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
    chmod +x /usr/local/bin/docker-compose
fi

# Create application directory
mkdir -p /opt/resume-server
cd /opt/resume-server

# Clone repository (replace with your repo URL)
if [ ! -d ".git" ]; then
    echo "Cloning repository..."
    git clone https://github.com/Hemanth5603/resume-go-server.git .
fi

# Create .env file from metadata (if provided)
if [ -f "/tmp/app.env" ]; then
    cp /tmp/app.env .env
else
    echo "Warning: No .env file found. Please create one manually."
    cat > .env << 'EOF'
PORT=3000
# MongoDB Atlas - Update with your connection string
DATABASE_URL=mongodb+srv://username:password@cluster.mongodb.net/resumedb?retryWrites=true&w=majority
# Or use local MongoDB from docker-compose
# DATABASE_URL=mongodb://mongo:27017/resumedb
JWKS_URL=https://your-clerk-domain.clerk.accounts.dev/.well-known/jwks.json
EOF
fi

# Build and start the application
echo "Building and starting application..."
docker-compose up -d --build

# Setup log rotation
cat > /etc/logrotate.d/resume-server << 'EOF'
/var/log/resume-server/*.log {
    daily
    rotate 7
    compress
    delaycompress
    missingok
    notifempty
    create 0640 root root
}
EOF

echo "Setup complete! Application should be running on port 3000"
echo "Check status with: docker-compose ps"
echo "View logs with: docker-compose logs -f"

