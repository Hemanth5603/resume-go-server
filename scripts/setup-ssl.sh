#!/bin/bash
# Setup SSL with Let's Encrypt

set -e

DOMAIN="${1:-your-domain.com}"
EMAIL="${2:-your-email@example.com}"

if [ "$DOMAIN" = "your-domain.com" ]; then
    echo "Usage: ./setup-ssl.sh your-domain.com your-email@example.com"
    exit 1
fi

echo "Setting up SSL for $DOMAIN..."

# Install Certbot
sudo apt-get update
sudo apt-get install -y certbot python3-certbot-nginx

# Install Nginx if not installed
if ! command -v nginx &> /dev/null; then
    echo "Installing Nginx..."
    sudo apt-get install -y nginx
fi

# Stop Nginx temporarily
sudo systemctl stop nginx

# Obtain certificate
sudo certbot certonly --standalone \
    -d $DOMAIN \
    -d www.$DOMAIN \
    --non-interactive \
    --agree-tos \
    --email $EMAIL \
    --preferred-challenges http

# Copy nginx config
sudo cp scripts/nginx.conf /etc/nginx/sites-available/resume-server

# Update domain in config
sudo sed -i "s/your-domain.com/$DOMAIN/g" /etc/nginx/sites-available/resume-server

# Enable site
sudo ln -sf /etc/nginx/sites-available/resume-server /etc/nginx/sites-enabled/
sudo rm -f /etc/nginx/sites-enabled/default

# Test nginx config
sudo nginx -t

# Start Nginx
sudo systemctl start nginx
sudo systemctl enable nginx

# Setup auto-renewal
sudo systemctl enable certbot.timer
sudo systemctl start certbot.timer

echo "SSL setup complete!"
echo "Your site should now be accessible at https://$DOMAIN"

