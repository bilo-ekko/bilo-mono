#!/usr/bin/env bash

# killall.sh - Kill all processes from bilo-mono apps by port
# This script kills processes running on ports used by the monorepo apps
# Works even if the terminal window was closed

set -e

# Color codes for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

echo -e "${BLUE}================================================${NC}"
echo -e "${BLUE}  Bilo Mono - Kill All Running Apps${NC}"
echo -e "${BLUE}================================================${NC}"
echo ""

# Function to kill process on a specific port
kill_port() {
    local app_name=$1
    local port=$2
    
    echo -e "${YELLOW}Checking port ${port} (${app_name})...${NC}"
    
    # Find PIDs listening on the port
    local pids=$(lsof -ti tcp:${port} 2>/dev/null || true)
    
    if [ -z "$pids" ]; then
        echo -e "  ${GREEN}✓${NC} No process found on port ${port}"
        return 0
    fi
    
    # Kill each PID
    for pid in $pids; do
        local process_name=$(ps -p $pid -o comm= 2>/dev/null || echo "unknown")
        echo -e "  ${RED}✗${NC} Found process: ${process_name} (PID: ${pid})"
        
        if kill -9 $pid 2>/dev/null; then
            echo -e "  ${GREEN}✓${NC} Killed PID ${pid}"
        else
            echo -e "  ${RED}✗${NC} Failed to kill PID ${pid} (may require sudo)"
        fi
    done
}

# Kill all app processes by port
kill_port "api-golang" 8080
echo ""
kill_port "api-nestjs" 3000
echo ""
kill_port "web-dashboard" 4000
echo ""
kill_port "web-sdks-apps" 4001
echo ""

echo -e "${BLUE}================================================${NC}"
echo -e "${GREEN}All processes checked and killed!${NC}"
echo -e "${BLUE}================================================${NC}"
echo ""

# Additional cleanup: kill any node/go processes related to the monorepo
echo -e "${YELLOW}Cleaning up additional processes...${NC}"

# Kill any nest processes
pkill -9 -f "nest start" 2>/dev/null && echo -e "  ${GREEN}✓${NC} Killed nest processes" || echo -e "  ${GREEN}✓${NC} No nest processes found"

# Kill any next dev processes
pkill -9 -f "next dev" 2>/dev/null && echo -e "  ${GREEN}✓${NC} Killed next dev processes" || echo -e "  ${GREEN}✓${NC} No next dev processes found"

# Kill any vite dev processes
pkill -9 -f "vite dev" 2>/dev/null && echo -e "  ${GREEN}✓${NC} Killed vite dev processes" || echo -e "  ${GREEN}✓${NC} No vite dev processes found"

# Kill any go run processes from api-golang
pkill -9 -f "go run.*api-golang" 2>/dev/null && echo -e "  ${GREEN}✓${NC} Killed go run processes" || echo -e "  ${GREEN}✓${NC} No go run processes found"

# Kill any api-golang binary processes
pkill -9 -f "api-golang" 2>/dev/null && echo -e "  ${GREEN}✓${NC} Killed api-golang binary" || echo -e "  ${GREEN}✓${NC} No api-golang binary found"

echo ""
echo -e "${GREEN}✓ All cleanup complete!${NC}"
echo ""
