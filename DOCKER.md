# Docker Guide for bilo-mono

Complete guide for building, running, and managing Docker containers in this monorepo.

## Table of Contents

- [Quick Start](#quick-start)
- [Available Services](#available-services)
- [Docker Compose](#docker-compose)
- [Individual Container Management](#individual-container-management)
- [Development vs Production](#development-vs-production)
- [Troubleshooting](#troubleshooting)
- [Best Practices](#best-practices)

## Quick Start

### Option 1: Using Helper Scripts (Easiest)

```bash
# Build all Docker images
./scripts/docker-build.sh

# Run all containers
./scripts/docker-run.sh

# Stop all containers
docker stop $(docker ps -q --filter name=bilo-)
```

### Option 2: Using Docker Compose (Recommended)

```bash
# Build and start all services
docker-compose up --build

# Run in detached mode (background)
docker-compose up -d --build

# View logs
docker-compose logs -f

# Stop all services
docker-compose down
```

## Available Services

| Service | Port | Container Name | Health Check Endpoint |
|---------|------|----------------|----------------------|
| api-golang | 8080 | bilo-api-golang | http://localhost:8080/api/health |
| api-nestjs | 3000 | bilo-api-nestjs | http://localhost:3000/ |
| web-dashboard | 4000 | bilo-web-dashboard | http://localhost:4000/ |
| web-sdks-apps | 4001 | bilo-web-sdks-apps | http://localhost:4001/ |

## Docker Compose

### Production

```bash
# Start all services in production mode
docker-compose up -d

# Build specific service
docker-compose build api-golang

# Rebuild without cache
docker-compose build --no-cache api-golang

# View logs for specific service
docker-compose logs -f api-golang

# Restart specific service
docker-compose restart api-nestjs

# Stop all services
docker-compose down

# Stop and remove volumes
docker-compose down -v
```

### Development

```bash
# Start with development overrides (hot reload)
docker-compose -f docker-compose.yml -f docker-compose.dev.yml up

# Build and start in dev mode
docker-compose -f docker-compose.yml -f docker-compose.dev.yml up --build
```

### Selective Services

```bash
# Start only backend services
docker-compose up api-golang api-nestjs

# Start only frontend
docker-compose up web-dashboard web-sdks-apps
```

## Helper Scripts

The monorepo includes two helper scripts for Docker management:

### docker-build.sh

Builds all Docker images with colored output and progress tracking.

```bash
# Build with 'latest' tag
./scripts/docker-build.sh

# Build with specific version
./scripts/docker-build.sh v1.0.0
```

**Features:**
- Builds all 4 application images in sequence
- Colored output for easy progress tracking
- Summary of successful and failed builds
- Lists all built images at the end
- Exit code 1 if any build fails

### docker-run.sh

Runs all Docker containers with proper networking and health checks.

```bash
# Run with 'latest' tag
./scripts/docker-run.sh

# Run with specific version
./scripts/docker-run.sh v1.0.0
```

**Features:**
- Creates `bilo-network` if it doesn't exist
- Starts containers in the correct order (backends first)
- Waits for services to become healthy
- Shows running containers and service URLs
- Provides useful Docker commands for management
- Exit code 1 if any container fails to start

**Service URLs after running:**
- Go API: http://localhost:8080
- NestJS API: http://localhost:3000
- Next.js Dashboard: http://localhost:4000
- SvelteKit SDKs: http://localhost:4001

## Individual Container Management

### api-golang (Go API)

```bash
# Build
docker build -f apps/backend/api-golang/Dockerfile -t bilo-api-golang:latest .

# Run
docker run -d -p 8080:8080 --name api-golang bilo-api-golang:latest

# View logs
docker logs -f api-golang

# Stop and remove
docker rm -f api-golang
```

### api-nestjs (NestJS API)

```bash
# Build
docker build -f apps/backend/api-nestjs/Dockerfile -t bilo-api-nestjs:latest .

# Run
docker run -d -p 3000:3000 --name api-nestjs bilo-api-nestjs:latest

# With environment variables
docker run -d -p 3000:3000 \
  -e NODE_ENV=production \
  -e PORT=3000 \
  --name api-nestjs bilo-api-nestjs:latest
```

### web-dashboard (Next.js)

```bash
# Build
docker build -f apps/frontend/web-dashboard/Dockerfile -t bilo-web-dashboard:latest .

# Run
docker run -d -p 4000:3000 \
  -e NEXT_PUBLIC_API_GOLANG_URL=http://localhost:8080 \
  -e NEXT_PUBLIC_API_NESTJS_URL=http://localhost:3000 \
  --name web-dashboard bilo-web-dashboard:latest
```

### web-sdks-apps (SvelteKit)

```bash
# Build
docker build -f apps/frontend/web-sdks-apps/Dockerfile -t bilo-web-sdks-apps:latest .

# Run
docker run -d -p 4001:4001 --name web-sdks-apps bilo-web-sdks-apps:latest
```

## Development vs Production

### Production Build

- Optimized multi-stage builds
- Minimal final image size
- Production dependencies only
- Runs built/compiled code
- Uses `docker-compose.yml`

### Development Build

- Includes dev dependencies
- Volume mounts for hot reload
- Source maps enabled
- Debug mode
- Uses `docker-compose.yml` + `docker-compose.dev.yml`

## Troubleshooting

### Container won't start

```bash
# Check logs
docker-compose logs [service-name]

# Check container status
docker ps -a

# Inspect container
docker inspect [container-name]

# Check if port is already in use
lsof -i :[port]
# Or use the killall script
./scripts/killall.sh
```

### Build fails

```bash
# Clean build (no cache)
docker-compose build --no-cache [service-name]

# Remove all unused images
docker image prune -a

# Remove all containers and rebuild
docker-compose down
docker-compose up --build
```

### Network issues between containers

```bash
# Ensure containers are on the same network
docker network ls
docker network inspect bilo-mono_bilo-network

# Restart with network recreation
docker-compose down
docker-compose up
```

### High memory/CPU usage

```bash
# View resource usage
docker stats

# Limit resources in docker-compose.yml
services:
  api-golang:
    deploy:
      resources:
        limits:
          cpus: '0.5'
          memory: 512M
```

### Clean up everything

```bash
# Stop all containers
docker-compose down

# Remove all bilo-mono images
docker images | grep bilo | awk '{print $3}' | xargs docker rmi -f

# Remove all unused Docker resources
docker system prune -a
```

## Best Practices

### 1. Use .dockerignore

The `.dockerignore` file is configured to exclude unnecessary files from the build context, speeding up builds.

### 2. Layer Caching

Moon's multi-stage builds are optimized for layer caching. Make code changes frequently but dependency changes rarely for fastest builds.

### 3. Health Checks

All services include health checks. Use `docker-compose ps` to see health status.

```bash
# View health status
docker-compose ps
```

### 4. Logs Management

```bash
# View logs with timestamps
docker-compose logs -f -t

# Limit log lines
docker-compose logs -f --tail=100

# Save logs to file
docker-compose logs > logs.txt
```

### 5. Environment Variables

Use `.env` file at the root for environment variables:

```bash
# Create .env file
cat > .env <<EOF
NODE_ENV=production
GO_ENV=production
PORT=3000
EOF

# docker-compose will automatically load it
docker-compose up
```

### 6. Production Deployment

```bash
# Tag images with versions
docker build -f apps/backend/api-golang/Dockerfile -t bilo-api-golang:1.0.0 .

# Push to registry (if using one)
docker tag bilo-api-golang:1.0.0 your-registry/bilo-api-golang:1.0.0
docker push your-registry/bilo-api-golang:1.0.0
```

### 7. Security

- Never commit sensitive data to Dockerfiles
- Use environment variables for secrets
- Scan images for vulnerabilities:

```bash
docker scan bilo-api-golang:latest
```

## Useful Commands

```bash
# Remove stopped containers
docker container prune

# Remove unused images
docker image prune

# Remove unused volumes
docker volume prune

# Remove everything unused
docker system prune -a --volumes

# View container disk usage
docker system df

# Execute command in running container
docker-compose exec api-golang sh
docker-compose exec api-nestjs sh

# Copy files from container
docker cp api-golang:/app/logs ./logs
```

## Advanced Usage

### Custom Network Configuration

```bash
# Create custom network
docker network create --driver bridge bilo-custom-network

# Use in docker-compose.yml
networks:
  default:
    external:
      name: bilo-custom-network
```

### Volume Mounts for Development

```bash
# Mount specific directories for hot reload
docker run -v $(pwd)/apps/backend/api-golang:/app/apps/backend/api-golang \
  bilo-api-golang:latest
```

### Multi-Architecture Builds

```bash
# Build for multiple architectures (ARM64, AMD64)
docker buildx build --platform linux/amd64,linux/arm64 \
  -f apps/backend/api-golang/Dockerfile \
  -t bilo-api-golang:latest .
```

## Integration with CI/CD

Example GitHub Actions workflow:

```yaml
name: Docker Build and Push

on: [push]

jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v2
      
      - name: Build images
        run: docker-compose build
      
      - name: Push to registry
        run: |
          echo "${{ secrets.REGISTRY_PASSWORD }}" | docker login -u "${{ secrets.REGISTRY_USERNAME }}" --password-stdin
          docker-compose push
```

## Related Documentation

- [Main README](./README.md)
- [Scripts README](./scripts/README.md)
- [DOCS.md](./DOCS.md)
- [Moon Docker Documentation](https://moonrepo.dev/docs/guides/docker)
