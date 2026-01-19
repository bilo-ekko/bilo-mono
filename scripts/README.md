# Scripts

Utility scripts for managing the bilo-mono monorepo.

## Available Scripts

### docker-build.sh

Build all Docker images for the monorepo applications.

**Usage:**
```bash
# Build all images with 'latest' tag
./scripts/docker-build.sh

# Build with specific version tag
./scripts/docker-build.sh v1.0.0
```

**What it does:**
- Builds Docker images for all 4 applications:
  - `bilo-api-golang` (Go API)
  - `bilo-api-nestjs` (NestJS API)
  - `bilo-web-dashboard` (Next.js Dashboard)
  - `bilo-web-sdks-apps` (SvelteKit SDKs)
- Uses multi-stage builds for optimized images
- Provides colored output for build status
- Shows summary of successful and failed builds
- Lists all built images at the end

**When to use:**
- Before deploying to production
- After making changes to Dockerfiles
- When you want to test the full containerized stack
- For CI/CD pipelines

### docker-run.sh

Run all Docker containers for the monorepo applications.

**Usage:**
```bash
# Run all containers with 'latest' tag
./scripts/docker-run.sh

# Run with specific version tag
./scripts/docker-run.sh v1.0.0
```

**What it does:**
- Creates a `bilo-network` Docker network if it doesn't exist
- Starts all 4 application containers in the correct order:
  1. Backend services (Go API, NestJS API)
  2. Frontend services (Next.js Dashboard, SvelteKit SDKs)
- Configures port mappings:
  - Go API: `localhost:8080`
  - NestJS API: `localhost:3000`
  - Next.js Dashboard: `localhost:4000`
  - SvelteKit SDKs: `localhost:4001`
- Waits for services to become healthy
- Shows running containers and service URLs
- Provides useful Docker commands for logs and cleanup

**When to use:**
- Testing the full containerized stack locally
- Verifying Docker builds work correctly
- Simulating production environment
- Alternative to docker-compose for more control

### killall.sh / killall.ps1

Kill all running processes from the monorepo apps by their ports.

**Usage (macOS/Linux):**
```bash
# From the monorepo root
./scripts/killall.sh

# Or from anywhere
bash /path/to/bilo-mono/scripts/killall.sh
```

**Usage (Windows PowerShell):**
```powershell
# From the monorepo root
.\scripts\killall.ps1

# Or from anywhere
& "C:\path\to\bilo-mono\scripts\killall.ps1"
```

**What it does:**
- Kills processes on port 8080 (api-golang)
- Kills processes on port 3000 (api-nestjs)
- Kills processes on port 4000 (web-dashboard)
- Kills processes on port 4001 (web-sdks-apps)
- Cleans up any remaining nest, next, vite, and go run processes
- Works even if terminal windows were closed

**When to use:**
- Ports are stuck/occupied after closing terminals
- Apps won't restart due to port conflicts
- Need to quickly stop all running development servers
- Cleaning up before starting fresh development session

### run.sh

Development server launcher script.

## Adding New Scripts

When adding new scripts to this directory:

1. Add a descriptive comment at the top explaining what it does
2. Make bash scripts executable: `chmod +x script-name.sh`
3. Use consistent naming: kebab-case for filenames
4. Document the script in this README
5. Add error handling with `set -e` for bash scripts
6. Use colored output for better UX

## Script Conventions

- **Bash scripts** (`.sh`): For macOS/Linux
- **PowerShell scripts** (`.ps1`): For Windows
- **Cross-platform**: Provide both when possible
- **Exit codes**: 0 for success, non-zero for errors
- **Output**: Use colors for better readability
  - 🔵 Blue: Headers/sections
  - 🟢 Green: Success messages
  - 🟡 Yellow: Warnings/info
  - 🔴 Red: Errors
