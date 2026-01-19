#!/usr/bin/env bash

# docker-build.sh - Build all Docker images for bilo-mono apps
# This script builds Docker images for all applications in the monorepo

set -e

# Color codes for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

echo -e "${BLUE}================================================${NC}"
echo -e "${BLUE}  Bilo Mono - Docker Image Builder${NC}"
echo -e "${BLUE}================================================${NC}"
echo ""

# Get version from argument or use 'latest'
VERSION="${1:-latest}"
echo -e "${YELLOW}Building images with tag: ${VERSION}${NC}"
echo ""

# Function to build a Docker image
build_image() {
    local app_name=$1
    local dockerfile_path=$2
    local image_name=$3
    
    echo -e "${BLUE}Building ${app_name}...${NC}"
    echo -e "  Dockerfile: ${dockerfile_path}"
    echo -e "  Image: ${image_name}:${VERSION}"
    
    if docker build -f "${dockerfile_path}" -t "${image_name}:${VERSION}" .; then
        echo -e "${GREEN}✓ Successfully built ${image_name}:${VERSION}${NC}"
        echo ""
        return 0
    else
        echo -e "${RED}✗ Failed to build ${image_name}:${VERSION}${NC}"
        echo ""
        return 1
    fi
}

# Track build status
failed_builds=()
successful_builds=()

# Build Backend Services
echo -e "${YELLOW}━━━ Building Backend Services ━━━${NC}"
echo ""

if build_image "Go API" \
    "apps/backend/api-golang/Dockerfile" \
    "bilo-api-golang"; then
    successful_builds+=("bilo-api-golang")
else
    failed_builds+=("bilo-api-golang")
fi

if build_image "NestJS API" \
    "apps/backend/api-nestjs/Dockerfile" \
    "bilo-api-nestjs"; then
    successful_builds+=("bilo-api-nestjs")
else
    failed_builds+=("bilo-api-nestjs")
fi

# Build Frontend Services
echo -e "${YELLOW}━━━ Building Frontend Services ━━━${NC}"
echo ""

if build_image "Next.js Dashboard" \
    "apps/frontend/web-dashboard/Dockerfile" \
    "bilo-web-dashboard"; then
    successful_builds+=("bilo-web-dashboard")
else
    failed_builds+=("bilo-web-dashboard")
fi

if build_image "SvelteKit SDKs" \
    "apps/frontend/web-sdks-apps/Dockerfile" \
    "bilo-web-sdks-apps"; then
    successful_builds+=("bilo-web-sdks-apps")
else
    failed_builds+=("bilo-web-sdks-apps")
fi

# Summary
echo -e "${BLUE}================================================${NC}"
echo -e "${BLUE}  Build Summary${NC}"
echo -e "${BLUE}================================================${NC}"
echo ""

if [ ${#successful_builds[@]} -gt 0 ]; then
    echo -e "${GREEN}✓ Successfully built ${#successful_builds[@]} image(s):${NC}"
    for image in "${successful_builds[@]}"; do
        echo -e "  - ${image}:${VERSION}"
    done
    echo ""
fi

if [ ${#failed_builds[@]} -gt 0 ]; then
    echo -e "${RED}✗ Failed to build ${#failed_builds[@]} image(s):${NC}"
    for image in "${failed_builds[@]}"; do
        echo -e "  - ${image}:${VERSION}"
    done
    echo ""
    exit 1
fi

echo -e "${GREEN}✓ All images built successfully!${NC}"
echo ""
echo -e "${YELLOW}Next steps:${NC}"
echo -e "  Run images individually: ${BLUE}./docker-run.sh${NC}"
echo -e "  Or use docker-compose:   ${BLUE}docker-compose up${NC}"
echo ""

# List all built images
echo -e "${YELLOW}Built images:${NC}"
docker images | grep "bilo-" | grep "${VERSION}"
echo ""
