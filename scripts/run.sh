#!/usr/bin/env bash

set -e

# Colors for output
GREEN='\033[0;32m'
BLUE='\033[0;34m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m' # No Color

DOCS_FILE="$(dirname "$0")/../DOCS.md"

# Function to display help
show_help() {
    echo -e "${BLUE}Available commands from DOCS.md:${NC}\n"
    
    # First command: setup proto (hardcoded)
    local counter=1
    echo -e "${GREEN}[$counter]${NC} ${YELLOW}Install Proto dependency manager${NC}"
    echo -e "    Command: ${BLUE}bash <(curl -fsSL https://moonrepo.dev/install/proto.sh)${NC}"
    echo ""
    ((counter++))
    
    # Parse DOCS.md and extract commands
    local current_section=""
    
    while IFS= read -r line; do
        # Detect section headers
        if [[ $line =~ ^##[[:space:]]+(.*) ]]; then
            current_section="${BASH_REMATCH[1]}"
        elif [[ $line =~ ^###[[:space:]]+(.*) ]]; then
            current_section="${BASH_REMATCH[1]}"
        fi
        
        # Extract commands (lines with backticks containing commands)
        if [[ $line =~ \`([^\`]+)\` ]]; then
            command="${BASH_REMATCH[1]}"
            # Skip if it's just a reference to a file or URL
            # Also skip the proto install command since we hardcode it as #1
            if [[ ! $command =~ ^http && ! $command =~ ^\. && $command =~ [[:space:]] && ! $command =~ "moonrepo.dev/install/proto.sh" ]]; then
                # Special case: custom titles for specific commands
                local display_section="$current_section"
                if [[ $command == "proto install moon 2.0.0-beta.0" ]]; then
                    display_section="Install moon v2"
                elif [[ $command == "moon init" ]]; then
                    display_section="Initialising moon monorepo"
                fi
                
                echo -e "${GREEN}[$counter]${NC} ${YELLOW}$display_section${NC}"
                echo -e "    Command: ${BLUE}$command${NC}"
                
                # Extract description if it follows the command
                if [[ $line =~ \`[^\`]+\`[[:space:]]*\((.+)\) ]]; then
                    desc="${BASH_REMATCH[1]}"
                    echo -e "    Description: $desc"
                fi
                echo ""
                ((counter++))
            fi
        fi
    done < "$DOCS_FILE"
    
    echo -e "${BLUE}Usage:${NC}"
    echo "  ./run.sh --help              Show this help message"
    echo "  ./run.sh --list              List all commands"
    echo "  ./run.sh <number>            Run command by number"
    echo "  ./run.sh <command>           Run command directly (e.g., './run.sh moon projects')"
    echo ""
}

# Function to list commands in a simple format
list_commands() {
    local counter=1
    
    # First command: setup proto (hardcoded)
    echo "[$counter] bash <(curl -fsSL https://moonrepo.dev/install/proto.sh)"
    ((counter++))
    
    while IFS= read -r line; do
        if [[ $line =~ \`([^\`]+)\` ]]; then
            command="${BASH_REMATCH[1]}"
            # Skip if it's just a reference to a file or URL
            # Also skip the proto install command since we hardcode it as #1
            if [[ ! $command =~ ^http && ! $command =~ ^\. && $command =~ [[:space:]] && ! $command =~ "moonrepo.dev/install/proto.sh" ]]; then
                echo "[$counter] $command"
                ((counter++))
            fi
        fi
    done < "$DOCS_FILE"
}

# Function to get command by number
get_command_by_number() {
    local target_number=$1
    local counter=1
    
    # First command: setup proto (hardcoded)
    if [[ $counter -eq $target_number ]]; then
        echo 'bash <(curl -fsSL https://moonrepo.dev/install/proto.sh)'
        return 0
    fi
    ((counter++))
    
    while IFS= read -r line; do
        if [[ $line =~ \`([^\`]+)\` ]]; then
            command="${BASH_REMATCH[1]}"
            # Skip if it's just a reference to a file or URL
            # Also skip the proto install command since we hardcode it as #1
            if [[ ! $command =~ ^http && ! $command =~ ^\. && $command =~ [[:space:]] && ! $command =~ "moonrepo.dev/install/proto.sh" ]]; then
                if [[ $counter -eq $target_number ]]; then
                    echo "$command"
                    return 0
                fi
                ((counter++))
            fi
        fi
    done < "$DOCS_FILE"
    
    return 1
}

# Function to execute a command
execute_command() {
    local cmd="$1"
    
    echo -e "${GREEN}Executing:${NC} ${BLUE}$cmd${NC}\n"
    eval "$cmd"
}

# Main script logic
main() {
    # Check if DOCS.md exists
    if [[ ! -f "$DOCS_FILE" ]]; then
        echo -e "${RED}Error: DOCS.md not found at $DOCS_FILE${NC}"
        exit 1
    fi
    
    # No arguments - show help
    if [[ $# -eq 0 ]]; then
        show_help
        exit 0
    fi
    
    # Handle flags
    case "$1" in
        --help|-h)
            show_help
            exit 0
            ;;
        --list|-l)
            list_commands
            exit 0
            ;;
        *)
            # Check if argument is a number
            if [[ "$1" =~ ^[0-9]+$ ]]; then
                cmd=$(get_command_by_number "$1")
                if [[ $? -eq 0 ]]; then
                    execute_command "$cmd"
                else
                    echo -e "${RED}Error: Command number $1 not found${NC}"
                    echo -e "Run '${BLUE}./run.sh --list${NC}' to see available commands"
                    exit 1
                fi
            else
                # Treat all arguments as a command to execute
                execute_command "$*"
            fi
            ;;
    esac
}

main "$@"
