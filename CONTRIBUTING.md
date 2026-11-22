# Contributing

Thanks for your interest in contributing!

## How to Contribute

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Submit a pull request

## Development Setup

Go 1.27+ is the only requirement:

    go build ./...
    go vet ./...
    go test ./... -race

CI runs the same checks on every pull request. Formatting and lint policy
is enforced by the maintainer's BuildFlow gate at merge time; match the
style of the surrounding code, and extend the existing test tables rather
than adding parallel suites.

## Reporting Issues

Please use GitHub Issues to report bugs or request features.
