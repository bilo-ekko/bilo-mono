# Moon monorepo

> **Prerequisites**
>
> You should [install proto](https://moonrepo.dev/proto) dependency manager.
> `bash <(curl -fsSL https://moonrepo.dev/install/proto.sh)`
>
> You should install moon with `proto install moon 2.0.0-beta.0`

> **💡 Tip:** Use `.scripts/run.sh --help` to see all commands from this file and execute them easily!

There is a VSCode (and Cursor) extension: [moon console](https://marketplace.visualstudio.com/items?itemName=moonrepo.moon-console)

There is also a [Cheat sheet](https://moonrepo.dev/docs/cheat-sheet) for moon commands.

[Feature comparison table between: moon, Nx & turborepo](https://moonrepo.dev/docs/comparison)

## Quickstart

- `moon run :dev`

### Getting Setup up

> You can clone said monorepo (in this case [https://github.com/bilo-ekko/bilo-mono](https://github.com/bilo-ekko/bilo-mono)).

Alternatively, if you want to add this to your own monorepo, navigate to it and run:

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