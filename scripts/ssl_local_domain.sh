#!/bin/bash
DOMAIN="my.domain"

# 1. Generate self-signed TLS certs for the domain
openssl req -x509 -nodes -days 365 -newkey rsa:2048 \
  -keyout tls.key -out tls.crt \
  -subj "/CN=$DOMAIN" -addext "subjectAltName=DNS:$DOMAIN"

# 2. Add domain to /etc/hosts
echo "127.0.0.1 $DOMAIN" | sudo tee -a /etc/hosts

# 3. Install server with TLS and custom domain
leap server install -c tls.crt -k tls.key --domain $DOMAIN