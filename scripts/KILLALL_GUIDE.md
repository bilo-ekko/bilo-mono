# Killall Script Guide

## Overview

The `killall.sh` and `killall.ps1` scripts provide a reliable way to kill all running processes from the bilo-mono monorepo by their ports. This works even if terminal windows were closed without properly stopping the processes.

## Problem Solved

When developing with multiple apps, you may encounter:
- Port conflicts preventing apps from starting
- Processes running in the background after closing terminals
- "Address already in use" errors
- Need to restart all services cleanly

## Usage

### macOS/Linux

```bash
# From the monorepo root
./scripts/killall.sh

# From any directory
bash /path/to/bilo-mono/scripts/killall.sh
```

### Windows (PowerShell)

```powershell
# From the monorepo root
.\scripts\killall.ps1

# From any directory
& "C:\path\to\bilo-mono\scripts\killall.ps1"
```

## What It Does

The script performs the following actions:

### 1. Port-Based Process Killing

Kills processes listening on these ports:

| Port | App             | Technology |
|------|-----------------|------------|
| 8080 | api-golang      | Go         |
| 3000 | api-nestjs      | NestJS     |
| 4000 | web-dashboard   | Next.js    |
| 4001 | web-sdks-apps   | SvelteKit  |

### 2. Additional Cleanup

After port-based killing, the script also cleans up:
- `nest start` processes
- `next dev` processes  
- `vite dev` processes
- `go run` processes related to api-golang
- api-golang binary processes

## Output Example

```
================================================
  Bilo Mono - Kill All Running Apps
================================================

Checking port 8080 (api-golang)...
  ✗ Found process: main (PID: 12345)
  ✓ Killed PID 12345

Checking port 3000 (api-nestjs)...
  ✗ Found process: node (PID: 12346)
  ✓ Killed PID 12346

Checking port 4000 (web-dashboard)...
  ✓ No process found on port 4000

Checking port 4001 (web-sdks-apps)...
  ✓ No process found on port 4001

================================================
All processes checked and killed!
================================================

Cleaning up additional processes...
  ✓ No nest processes found
  ✓ Killed next dev processes
  ✓ No vite dev processes found
  ✓ No go run processes found
  ✓ No api-golang binary found

✓ All cleanup complete!
```

## Features

### Color-Coded Output

- 🔵 **Blue**: Headers and sections
- 🟢 **Green**: Success messages (process killed, nothing to kill)
- 🟡 **Yellow**: Information (checking ports)
- 🔴 **Red**: Found processes (before killing)

### Safe Execution

- Uses `lsof` to find processes by port (macOS/Linux)
- Uses `Get-NetTCPConnection` for Windows
- Handles cases where no processes are found
- Reports failures if unable to kill a process

### Works Anywhere

- Can be run from any directory
- No need to navigate to monorepo root
- Works even if terminals were force-closed

## Common Scenarios

### Scenario 1: Port Conflict on Startup

```bash
# Error when starting app
Error: listen EADDRINUSE: address already in use :::3000

# Solution
./scripts/killall.sh
moon run api-nestjs:dev
```

### Scenario 2: Force-Closed Terminal

```bash
# You closed terminal with Cmd+Q or closed window
# Processes are still running in background

# Solution - open new terminal
./scripts/killall.sh
```

### Scenario 3: Clean Restart

```bash
# Stop everything and start fresh
./scripts/killall.sh
sleep 1
moon run :dev
```

### Scenario 4: Switching Branches

```bash
# Before switching branches with running services
./scripts/killall.sh
git checkout feature-branch
moon run :dev
```

## Troubleshooting

### "Permission denied" on macOS/Linux

Make sure the script is executable:

```bash
chmod +x scripts/killall.sh
```

### "Failed to kill PID" message

Some processes may require elevated privileges:

```bash
sudo ./scripts/killall.sh
```

### Script not found

Ensure you're using the correct path:

```bash
# Check if script exists
ls -la scripts/killall.sh

# Run with full path
bash /full/path/to/bilo-mono/scripts/killall.sh
```

## Adding New Apps

If you add a new app to the monorepo with a different port:

1. Open `scripts/killall.sh`
2. Add a new `kill_port` call:
   ```bash
   kill_port "your-app-name" 5000
   echo ""
   ```
3. Update this documentation
4. Update the main README.md

## Related Documentation

- [Main README](../README.md) - General monorepo documentation
- [Scripts README](./README.md) - All available scripts
- [DOCS.md](../DOCS.md) - Moon commands and workflows
