# Moon monorepo

> **Prerequisites**
>
> You should [install proto](https://moonrepo.dev/proto) dependency manager.
> `bash <(curl -fsSL https://moonrepo.dev/install/proto.sh)`
>
> You should install moon with `proto install moon 2.0.0-beta.0`

> **💡 Tip:** Use `./scripts/run.sh --help` to see all commands from this file and execute them easily!

There is a VSCode (and Cursor) extension: [moon console](https://marketplace.visualstudio.com/items?itemName=moonrepo.moon-console)

There is also a [Cheat sheet](https://moonrepo.dev/docs/cheat-sheet) for moon commands.

[Feature comparison table between: moon, Nx & turborepo](https://moonrepo.dev/docs/comparison)

## Quickstart

- `moon run :dev`

### Getting Setup up

To initialise a moon repo requires one command, inside an arbitrary folder which will become the root of the monorepo.

- `moon init` (this creates the [.moon/workspace.yml](.moon/workspace.yml) in your monorepo root)

You can then add a `.moon/toolchains.yml` file, which configures the required versions for your stack / monorepo.

Sample [toolchains.yml](.moon/toolchains.yml)

To setup / teardown the toolchain (required versions):

- `moon setup` (installs all toolchain software specified in `.moon/toolchains.yml`)
- `moon teardown` (uninstalls the toolchain software installed with `moon setup`)


### Workspace commands

#### Toolchain info

- `moon toolchain info <ID>` (e.g. `node` or `go` as the ID)

#### Project details

- `moon projects`

#### Opening graphs in your browser

```sh
moon project-graph
```

```sh
moon task-graph
```

```sh
moon action-graph
```

#### Using tags to run various groups

Running tasks by tag:

```sh
moon run '#backend:dev'
```

Running combined tags

```sh
moon run 'nestjs+#typescript:dev'
```

Exclude a tag

```sh
moon run '#typescript,!nestjs:test'
```

#### Using `layer`, `stack` or `language` for running apps

Layer:
```sh
moon run :dev --query "projectLayer=application"
```

Stack:
```sh
moon run :dev --query "projectStack=backend"
```

Language:
```sh
moon run :dev --query "language=typescript"
```

### Utility Scripts

#### Kill All Running Apps

Kill all running development servers by their ports (works even if terminals were closed):

```sh
./scripts/killall.sh
```

This will:
- Kill processes on port 8080 (api-golang)
- Kill processes on port 3000 (api-nestjs)  
- Kill processes on port 4000 (web-dashboard)
- Kill processes on port 4001 (web-sdks-apps)
- Clean up any remaining nest, next, vite, and go run processes

Useful when:
- Ports are stuck/occupied after closing terminals
- Apps won't restart due to port conflicts
- Need to quickly stop all running development servers
- Cleaning up before starting fresh development session