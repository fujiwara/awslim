# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

awslim is a lightweight, fast alternative to the AWS CLI, built as a thin wrapper around AWS SDK Go v2. Service method wrappers are generated at build time, so a binary contains only the services listed in `gen.yaml` (or `AWSLIM_GEN`).

## Build & Development Commands

```bash
make                 # code generation + build binary (./awslim)
make gen             # code generation only
make awslim          # build binary only (after gen)
make test            # go test -v .
go test -v -run TestRun .
make clean           # remove generated files and run go mod tidy
go fmt ./... && go fix ./...   # always before committing
```

- `AWSLIM_GEN=ecs,s3,lambda make` builds only the listed services (overrides `gen.yaml`).
- `AWSLIM_OS` / `AWSLIM_ARCH`: cross-compilation targets.
- Run `make clean` before `make test` if generated files exist: tests assert the exact service list (`bar`, `baz`, `foo`) and fail when real services are registered.
- CI (`.github/workflows/test.yml`) runs `go test -v .` and then `go generate ./cmd/awslim-gen . && go build ./cmd/awslim/main.go`.

## Architecture

### Package Structure

The root package is `package sdkclient`. `cmd/awslim/main.go` only calls `sdkclient.Run()`.

| File | Purpose |
|---|---|
| `main.go` | `CLI` struct (kong flags), `Run`/`NewCLI`/`Dispatch`/`CallMethod`, Jsonnet evaluation (`evaluateJsonnet`, native funcs `_`, `env`, `must_env`), client options loading |
| `param.go` | `clientMethodParam` (input bytes, streams, dry-run, client options, next-token handling), `UnmarshalJSON` |
| `config.go` | `RuntimeConfig` loaded from `~/.config/awslim/config.(jsonnet\|json\|yaml\|yml)`: `open`, `aliases`, `client_options` |
| `flags.go` | Dynamic flags: unknown `--flag value` args are injected into the input JSON (kebab-case → PascalCase). Must recognize kong negated names (`--no-strict`) |
| `json.go` | Output marshaling, `--camel` key conversion |

### Three-Stage Code Generation

```
gen.yaml (or AWSLIM_GEN env var)
  → cmd/awslim-gen-gen/main.go    generates → cmd/awslim-gen/gen.go
  → cmd/awslim-gen/main.go+gen.go generates → {service}_gen.go + main_gen.go (root dir)
  → go build cmd/awslim/main.go   produces  → awslim binary
```

`*_gen.go` and `cmd/awslim-gen/gen.go` are build artifacts and are not committed (`make clean` removes them). `go generate` directives chain the stages:
- `cmd/awslim-gen/main.go`: `//go:generate go run ../awslim-gen-gen/main.go`
- `main.go` (root): `//go:generate go run cmd/awslim-gen/main.go cmd/awslim-gen/gen.go`

The generator template lives in `cmd/awslim-gen/main.go` (`serviceTemplateStr`). Any change to how a client is created or called goes there; verify with `AWSLIM_GEN=s3,sts make` since the generated code is not compiled by `go test`.

### Generated Code Pattern

Each wrapper (e.g. `s3_ListObjectsV2`) creates the client with `NewFromConfig(p.awsCfg, func(o *s3.Options) { p.ApplyClientOptions(o) })`, unmarshals the input JSON (strict by default), validates streams, honors dry-run, and calls the SDK method. The generator detects `io.Reader` input fields (`-i`), `*Length` fields (auto content-length), and `io.ReadCloser` output fields (`-o`) by reflection.

### Dispatch Flow

`Run()` → `NewCLI()` (load runtime config → resolve aliases → extract dynamic flags → kong parse, with `AWSLIM_*` env vars via `kong.DefaultEnvars`) → `Dispatch()`:
- No service → list services
- No method → list methods
- Otherwise → `CallMethod()`, looping on `FollowNext` when `--follow-next OutputField=InputField` is given

Client options (`--client-option`/`-C`, `AWSLIM_CLIENT_OPTION`, or `client_options.<service>` in the config) are merged (flag wins) and applied to the SDK `Options` struct via JSON unmarshal. Only JSON-decodable fields are supported.

## Testing

Tests live in the root package as `package sdkclient_test`. `export_test.go` exposes internals (`SetClientMethod`, `ClientMethods`, `ClientMethodParam`). `cli_test.go` registers mock services `foo`, `bar`, `baz` in `init()`, sets `XDG_CONFIG_HOME=testdata` (config in `testdata/awslim/config.yaml`), and uses table-driven cases (`TestCases`) with `Args`, `Env`, `Expect`, `IsError`. Add new CLI behavior as cases there.

## Binary Size

- Size is dominated by per-API serializer/deserializer code in the SDK service packages, so it scales with the number of methods included, not with anything in awslim itself. Measured (linux/amd64, `-s -w`): `sts` only 16MB (baseline), `ec2` all methods +30MB, ec2 limited to 2 methods ≈ baseline, all services (423) ≈ 640MB.
- The only effective reduction is restricting methods in `gen.yaml`: unreferenced methods are dead-code-eliminated by the linker. `-trimpath` has no effect; `-s -w` is already applied.
- UPX is not an option for the all-services binary: decompression time scales linearly with size (≈80ms per 46MB with default settings, ≈330ms with `--best --lzma`), so it would add over a second to startup, plus the whole image is held in memory. It also has issues on macOS.
- A large binary does not start slower by itself (pages are loaded lazily via mmap); the cost is download/disk size only.
- To inspect what contributes to size: `go build -o awslim ./cmd/awslim` (without `-s -w`) then `go tool nm -size -sort size awslim`. Ignore `B` (BSS) symbols such as `crypto/internal/fips140/drbg.memory`; they do not occupy file space.

## Release & Maintenance

- `.goreleaser.yml`: release configuration (linux/darwin × amd64/arm64), `Dockerfile` / `build-in-docker.sh` for custom builds.
- `.tagpr`: releases are cut by tagpr on merge to `main`; `CHANGELOG.md` is generated, do not edit by hand.
- `all-services.yaml` is regenerated from the SDK repo by `make all-services.yaml` (daily `check-sdk-updates.yml` workflow opens PRs); don't edit it manually.
