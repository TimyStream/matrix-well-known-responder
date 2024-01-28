# Easy to use well-known provider for Matrix

You can use this config for Docker Compose
```yaml
services:
  matrix-well-known-responder:
    container_name: matrix-well-known-responder
    hostname: matrix-well-known-responder
    build:
      context: https://github.com/TimyStream/matrix-well-known-responder.git
    ports:
      - "8080:8080"
    environment:
      - MATRIX_SERVER_URL=https://<your.server.domain>[:YOUR_PORT]
    restart: unless-stopped
```
