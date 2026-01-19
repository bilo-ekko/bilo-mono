#!/usr/bin/env bash

# docker-run.sh - Run all Docker containers for bilo-mono apps
# This script runs all Docker containers in the correct order

set -e

# Color codes for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

echo -e "${BLUE}================================================${NC}"
echo -e "${BLUE}  Bilo Mono - Docker Container Runner${NC}"
echo -e "${BLUE}================================================${NC}"
echo ""

# Get version from argument or use 'latest'
VERSION="${1:-latest}"
echo -e "${YELLOW}Running images with tag: ${VERSION}${NC}"
echo ""

# Function to check if container is running
is_container_running() {
    local container_name=$1
    docker ps --format '{{.Names}}' | grep -q "^${container_name}$"
}

# Function to stop and remove existing container
cleanup_container() {
    local container_name=$1
    if docker ps -a --format '{{.Names}}' | grep -q "^${container_name}$"; then
        echo -e "${YELLOW}  Stopping and removing existing container: ${container_name}${NC}"
        docker rm -f "${container_name}" > /dev/null 2>&1 || true
    fi
}

# Function to run a Docker container
run_container() {
    local app_name=$1
    local container_name=$2
    local image_name=$3
    local port_mapping=$4
    shift 4
    local env_vars=("$@")
    
    echo -e "${BLUE}Starting ${app_name}...${NC}"
    echo -e "  Container: ${container_name}"
    echo -e "  Image: ${image_name}:${VERSION}"
    echo -e "  Port: ${port_mapping}"
    
    # Clean up existing container
    cleanup_container "${container_name}"
    
    # Build docker run command
    local docker_cmd="docker run -d --name ${container_name} -p ${port_mapping}"
    
    # Add environment variables
    for env in "${env_vars[@]}"; do
        docker_cmd="${docker_cmd} -e ${env}"
    done
    
    # Add network
    docker_cmd="${docker_cmd} --network bilo-network"
    
    # Add restart policy
    docker_cmd="${docker_cmd} --restart unless-stopped"
    
    # Add image
    docker_cmd="${docker_cmd} ${image_name}:${VERSION}"
    
    # Run the container
    if eval "${docker_cmd}"; then
        echo -e "${GREEN}✓ Successfully started ${container_name}${NC}"
        echo ""
        return 0
    else
        echo -e "${RED}✗ Failed to start ${container_name}${NC}"
        echo ""
        return 1
    fi
}

# Function to wait for service to be healthy
wait_for_service() {
    local service_name=$1
    local url=$2
    local max_attempts=30
    local attempt=0
    
    echo -e "${YELLOW}  Waiting for ${service_name} to be ready...${NC}"
    
    while [ $attempt -lt $max_attempts ]; do
        if curl -s -f "${url}" > /dev/null 2>&1; then
            echo -e "${GREEN}  ✓ ${service_name} is ready!${NC}"
            return 0
        fi
        attempt=$((attempt + 1))
        sleep 2
    done
    
    echo -e "${RED}  ✗ ${service_name} failed to become ready${NC}"
    return 1
}

# Create network if it doesn't exist
if ! docker network ls | grep -q bilo-network; then
    echo -e "${YELLOW}Creating bilo-network...${NC}"
    docker network create bilo-network
    echo -e "${GREEN}✓ Network created${NC}"
    echo ""
fi

# Track running status
failed_containers=()
successful_containers=()

# Start Backend Services
echo -e "${YELLOW}━━━ Starting Backend Services ━━━${NC}"
echo ""

if run_container "Go API" \
    "bilo-api-golang" \
    "bilo-api-golang" \
    "8080:8080" \
    "PORT=8080" \
    "GO_ENV=production"; then
    successful_containers+=("bilo-api-golang")
    wait_for_service "Go API" "http://localhost:8080/api/health" || true
else
    failed_containers+=("bilo-api-golang")
fi

if run_container "NestJS API" \
    "bilo-api-nestjs" \
    "bilo-api-nestjs" \
    "3000:3000" \
    "PORT=3000" \
    "NODE_ENV=production"; then
    successful_containers+=("bilo-api-nestjs")
    wait_for_service "NestJS API" "http://localhost:3000/" || true
else
    failed_containers+=("bilo-api-nestjs")
fi

# Start Frontend Services
echo ""
echo -e "${YELLOW}━━━ Starting Frontend Services ━━━${NC}"
echo ""

if run_container "Next.js Dashboard" \
    "bilo-web-dashboard" \
    "bilo-web-dashboard" \
    "4000:3000" \
    "PORT=3000" \
    "NODE_ENV=production" \
    "NEXT_PUBLIC_API_GOLANG_URL=http://localhost:8080" \
    "NEXT_PUBLIC_API_NESTJS_URL=http://localhost:3000"; then
    successful_containers+=("bilo-web-dashboard")
    wait_for_service "Next.js Dashboard" "http://localhost:4000/" || true
else
    failed_containers+=("bilo-web-dashboard")
fi

if run_container "SvelteKit SDKs" \
    "bilo-web-sdks-apps" \
    "bilo-web-sdks-apps" \
    "4001:4001" \
    "PORT=4001" \
    "NODE_ENV=production" \
    "PUBLIC_API_GOLANG_URL=http://localhost:8080" \
    "PUBLIC_API_NESTJS_URL=http://localhost:3000"; then
    successful_containers+=("bilo-web-sdks-apps")
    wait_for_service "SvelteKit SDKs" "http://localhost:4001/" || true
else
    failed_containers+=("bilo-web-sdks-apps")
fi

# Summary
echo ""
echo -e "${BLUE}================================================${NC}"
echo -e "${BLUE}  Runtime Summary${NC}"
echo -e "${BLUE}================================================${NC}"
echo ""

if [ ${#successful_containers[@]} -gt 0 ]; then
    echo -e "${GREEN}✓ Successfully started ${#successful_containers[@]} container(s):${NC}"
    for container in "${successful_containers[@]}"; do
        echo -e "  - ${container}"
    done
    echo ""
fi

if [ ${#failed_containers[@]} -gt 0 ]; then
    echo -e "${RED}✗ Failed to start ${#failed_containers[@]} container(s):${NC}"
    for container in "${failed_containers[@]}"; do
        echo -e "  - ${container}"
    done
    echo ""
fi

# Show running containers
echo -e "${YELLOW}Running containers:${NC}"
docker ps --format "table {{.Names}}\t{{.Status}}\t{{.Ports}}" | grep "bilo-" || echo "  No bilo containers running"
echo ""

# Show service URLs
echo -e "${YELLOW}Service URLs:${NC}"
echo -e "  ${BLUE}Go API:${NC}           http://localhost:8080"
echo -e "  ${BLUE}NestJS API:${NC}       http://localhost:3000"
echo -e "  ${BLUE}Next.js Dashboard:${NC} http://localhost:4000"
echo -e "  ${BLUE}SvelteKit SDKs:${NC}   http://localhost:4001"
echo ""

echo -e "${YELLOW}Useful commands:${NC}"
echo -e "  View logs:           ${BLUE}docker logs -f <container-name>${NC}"
echo -e "  Stop all:            ${BLUE}docker stop \$(docker ps -q --filter name=bilo-)${NC}"
echo -e "  Remove all:          ${BLUE}docker rm \$(docker ps -aq --filter name=bilo-)${NC}"
echo -e "  Or use docker-compose: ${BLUE}docker-compose up -d${NC}"
echo ""

if [ ${#failed_containers[@]} -gt 0 ]; then
    exit 1
fi

echo -e "${GREEN}✓ All containers started successfully!${NC}"
echo ""
