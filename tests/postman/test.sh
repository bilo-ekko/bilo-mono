#!/usr/bin/env bash

set -e

# Colors for output
GREEN='\033[0;32m'
BLUE='\033[0;34m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m' # No Color

# Get the directory where this script is located
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
COLLECTION_PATH="$SCRIPT_DIR/collection.json"

echo -e "${BLUE}🧪 Running API Tests with Newman${NC}\n"

# Check if Newman is installed
if ! command -v newman &> /dev/null; then
    echo -e "${RED}❌ Newman is not installed${NC}"
    echo -e "${YELLOW}Installing Newman locally...${NC}\n"
    cd "$SCRIPT_DIR"
    npm install
    echo -e "\n${GREEN}✅ Newman installed${NC}\n"
fi

# Check if both servers are running
echo -e "${YELLOW}Checking if APIs are running...${NC}"

NEST_RUNNING=false
GO_RUNNING=false

if curl -s http://localhost:3000/ > /dev/null 2>&1; then
    echo -e "${GREEN}✅ NestJS API is running on port 3000${NC}"
    NEST_RUNNING=true
else
    echo -e "${RED}❌ NestJS API is NOT running on port 3000${NC}"
fi

if curl -s http://localhost:8080/ > /dev/null 2>&1; then
    echo -e "${GREEN}✅ Go API is running on port 8080${NC}"
    GO_RUNNING=true
else
    echo -e "${RED}❌ Go API is NOT running on port 8080${NC}"
fi

if [ "$NEST_RUNNING" = false ] || [ "$GO_RUNNING" = false ]; then
    echo -e "\n${YELLOW}⚠️  Please start the APIs first:${NC}"
    echo -e "   ${BLUE}moon run api-nest:dev${NC}"
    echo -e "   ${BLUE}moon run api-golang:dev${NC}\n"
    exit 1
fi

echo -e "\n${BLUE}Running tests...${NC}\n"

# Run Newman with the collection
if [ -x "$SCRIPT_DIR/node_modules/.bin/newman" ]; then
    # Use local newman
    "$SCRIPT_DIR/node_modules/.bin/newman" run "$COLLECTION_PATH" "$@"
else
    # Use global newman
    newman run "$COLLECTION_PATH" "$@"
fi

EXIT_CODE=$?

if [ $EXIT_CODE -eq 0 ]; then
    echo -e "\n${GREEN}✅ All tests passed!${NC}\n"
else
    echo -e "\n${RED}❌ Some tests failed${NC}\n"
fi

exit $EXIT_CODE
