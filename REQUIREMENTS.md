## T1 priorities
documentation:
- [X] [DOCS.md](./DOCS.md)

scripts to run from root:
- [X] install
- [X] build
- [X] dev
- [X] test

- [ ] changing common packages:
    - [ ] manual rebuild required?
    - [ ] sequential, targeted builds (build packages before apps)
    - [ ] deploying only apps depending on built packages
- [ ] running :dev in just a single project's root
- [ ] environment variables required at root, like in `turbo.json`
- [ ] test from a performance and size perspective
    - [ ] measure core commands (`install`, `build`, `test`, `dev`)
    - [ ] measure size of outputs, caches, etc.


## T2 priorities
- [ ] docker & moon