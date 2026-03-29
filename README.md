# go-mod-version

A simple CLI tool that prints the current Go module path along with the latest Git commit hash.

## Output Format

```
<module-path>@<commit-hash>
```

**Example:**

```
github.com/st-1dev/go-mod-version@a1b2c3d4e5f6...
```

## Installation

```bash
go install github.com/st-1dev/go-mod-version@latest
```

## Usage

Run the tool from the root of a Go project with a Git repository:

```bash
go-mod-version
```

The tool reads `go.mod` to determine the module path and uses the local `.git` history to obtain the latest commit hash.

## Build

This project uses [Mage](https://magefile.org/) as a build tool.

```bash
# Build the binary
mage build

# Format the source code
mage format
```

## Requirements

- Go 1.26+
- Git repository with at least one commit
- `go.mod` file in the current directory

## License

[MIT](LICENSE)
