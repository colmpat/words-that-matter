# words-that-matter
A tool for english learners that allows users to learn _words that matter_ from media in pop-culture

full-stack webapp that uses Go with HTMX

## Main components

* [`./cmd/`](./internal): commands / short executables
    * [`./cmd/migrate`](./cmd/migrate)
* [`./internal/`](./internal): internal packages
    * _these cannot be imported directly; if you think that an internal package should be made public, open an issue_
* [`./pkg/...`](./pkg): packages especially made to be imported by other projects
* [`./services/...`](./services): services / long-running executables
    * [`./services/ingestor`](./services/ingestor)
    * [`./services/web`](./services/web)
