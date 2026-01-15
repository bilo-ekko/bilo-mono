# bilo-mono

A polyglot monorepo managed with [Moon](https://moonrepo.dev/), featuring frontend and backend applications built with TypeScript, Go, and .NET.

---

## 📁 Directory Structure

```
bilo-mono/
├── .moon/                      # Moon configuration
│   ├── toolchains.yml          # Language/runtime versions
│   └── workspace.yml           # Workspace & project discovery
├── apps/
│   ├── backend/
│   │   ├── api-golang/         # Go API service
│   │   ├── api-nest/           # NestJS (TypeScript) 
│   └── frontend/
│       ├── web-dashboard/          # Next.js dashboard app
│       └── web-sdks/             # SvelteKit UI components
└── packages/                   # Shared packages (future)
```

---

## 🧰 Toolchains

| Tool    | Version   |
|---------|-----------|
| Node.js | 22.20.0   |
| pnpm    | 10.18.0   |
| Go      | 1.25.3    |

---

## 🚀 Getting Started

### Prerequisites

Install Moon globally:

[Install moon](https://moonrepo.dev/docs/install) preferably using [proto's toolchain](https://moonrepo.dev/proto)

> It is recommended to use `proto to manage moon (and other tools), as it allows for multiple versions to be installed and used. The other installation options only allow for a **single version** (typically the last installed).

### Setup

Clone the repo and let Moon handle toolchain installation:

```bash
git clone <repo-url>
cd bilo-mono
moon setup
```

---

## 🌙 Moon Commands

### Run Tasks

```bash
# Run a task for a specific project
moon run <project>:<task>

# Examples:
moon run api-nest:dev          # Start NestJS in dev mode
moon run api-golang:dev        # Start Go API in dev mode
moon run ApiDotNet:dev         # Start .NET API in watch mode
moon run dashboard:build       # Build Next.js dashboard
moon run sdk-ui:dev            # Start SvelteKit dev server
```

### Run Tasks Across All Projects

```bash
# Build all projects
moon run :build

# Test all projects
moon run :test

# Lint all projects
moon run :lint
```

### Quick Command Helper 🎯

For convenience, use the `run.sh` script to easily discover and execute commands from `DOCS.md`:

```bash
# Show all available commands with descriptions
./run.sh --help

# List commands in a simple format
./run.sh --list

# Run a command by its number
./run.sh 7           # Runs: moon projects

# Run any command directly
./run.sh moon run :build
```

### Check & CI

```bash
# Run all checks (affected projects only)
moon check

# Run all tasks in CI mode
moon ci
```

### Useful Commands

```bash
# List all projects
moon query projects

# List all tasks for a project
moon query tasks --project <project>

# View project graph
moon project-graph

# View task dependencies
moon task-graph <project>:<task>
```

---

## 📦 Projects Overview

### Backend

| Project       | Language   | Port | Tasks                                              |
|---------------|------------|------|----------------------------------------------------|
| `api-nest`    | TypeScript | 3000 | `build`, `dev`, `start`, `test`, `lint`, `format`  |
| `api-golang`  | Go         | 8080 | `build`, `dev`, `start`, `test`, `lint`, `format`, `tidy` |
| `ApiDotNet`   | C#/.NET    | 5000 | `build`, `dev`, `start`, `test`, `clean`, `restore`, `format` |

### Frontend

| Project          | Framework  | Port | Tasks                          |
|------------------|------------|------|--------------------------------|
| `web-dashboard`  | Next.js    | 4000 | `build`, `dev`, `start`, `lint`|
| `web-sdks`       | SvelteKit  | 4001 | `build`, `dev`, `check`, `preview` |

---

## 🔧 Common Workflows

### Development

```bash
# Start all backends
moon run api-nest:dev &
moon run api-golang:dev

# Start frontend apps
moon run web-dashboard:dev    # Runs on http://localhost:4000
moon run web-sdks:dev          # Runs on http://localhost:4001
```

### Stopping All Processes

If you need to kill all running development servers (even if terminals were closed):

```bash
# Kill all app processes by port
./scripts/killall.sh

# On Windows
.\scripts\killall.ps1
```

This will:
- Kill processes on port 8080 (api-golang)
- Kill processes on port 3000 (api-nestjs)
- Kill processes on port 4000 (web-dashboard)
- Kill processes on port 4001 (web-sdks-apps)
- Clean up any remaining nest, next, vite, and go run processes

### Build for Production

```bash
# Build everything
moon run :build

# Build specific project
moon run api-nest:build
```

### Testing

```bash
# Run all tests
moon run :test

# Run tests for specific project
moon run api-nest:test
moon run api-golang:test
```

### Code Quality

```bash
# Lint all projects
moon run :lint

# Format all projects
moon run :format
```

---

## Moon Monorepo – Supported Languages

This section is valid at the time of writing.

### ✅ Fully supported (managed toolchains)
Moon can install, pin versions, and fully manage these:

- **JavaScript ecosystem**
  - Node.js
  - Package managers: npm, pnpm, yarn, Bun
  - Runtimes: Bun, Deno
- **Go**
- **Rust**

These languages support `.moon/toolchains.yml` for deterministic installs and caching.

---

### 🧪 Partial / limited support
Recognized by Moon, but toolchain management is incomplete or evolving:

- **Python** (basic recognition, limited ecosystem support)

---

### ❌ Not supported as Moon toolchains
These run via the **system toolchain only** (Moon does *not* install or manage them):

- **.NET / C#**
- **Java / Kotlin**
- **PHP**
- **Ruby**
- **C / C++ / Swift / others**

You can still run tasks for these, but must manage versions yourself.

---

### TL;DR
- Use Moon toolchains for **JS, Go, Rust**
- Use system-installed tooling for **.NET and others**
- Moon still works great as an orchestrator across *all* languages


## 📖 Learn More

- [Moon Documentation](https://moonrepo.dev/docs)
- [Moon Task Configuration](https://moonrepo.dev/docs/config/project#tasks)
- [Moon Workspace Configuration](https://moonrepo.dev/docs/config/workspace)
