# Scripts

Utility scripts for managing the bilo-mono monorepo.

## Available Scripts

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
