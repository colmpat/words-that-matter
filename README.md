# words-that-matter

## Main components

* [`./internal/`](./internal): internal packages
    * _these cannot be imported directly; if you think that an internal package should be made public, open an issue_
* [`./pkg/...`](./pkg): packages especially made to be imported by other projects
* [`./services/...`](./services): services / long-running executables
    * [`./services/ingestor`](./services/ingestor)
    * [`./services/web`](./services/web)
