# Moon monorepo

> **Prerequisites**
>
> You should install [proto]() dependency manager.
>
> You should install moon with `proto install moon 2.0.0-beta.0`

## Quickstart

- `moon run :dev`

### Getting Setup up

> You can clone said monorepo (in this case [https://github.com/bilo-ekko/bilo-mono](https://github.com/bilo-ekko/bilo-mono)).

Alternatively, if you want to add this to your own monorepo, navigate to it and run:

- `moon init` (this creates the `.moon/workspace.yml` in your monorepo root)

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

- `moon project-graph`
- `moon task-graph`
- `moon action-graph`



