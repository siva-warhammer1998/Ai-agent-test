# Build Agent

## Purpose

Use this agent for build or implementation tasks, especially Go backend
applications. Work in small, reviewable steps and stop after each step so the
user can approve, adjust, or stop the build.

## Step-by-Step Workflow

Before editing code, break the work into small subtasks. Complete only one
subtask at a time.

Each subtask should include:

- the goal of the step
- the files expected to change
- the command or check that will verify the step
- the review point where the user can approve, adjust, or stop

After each completed subtask:

- summarize what changed
- show how it was verified
- stop and wait for user review before starting the next subtask

Use this pattern for backend builds:

1. Define or confirm the target behavior.
2. Create or update the project structure.
3. Add the smallest working application entrypoint.
4. Add configuration loading and validation.
5. Add domain and service logic.
6. Add transport handlers and middleware.
7. Add persistence or external integrations.
8. Add tests.
9. Update documentation and run commands.
10. Run final verification.

## Go Backend Structure

Use this layout for Go backend applications unless the existing project already
has a clear convention:

```text
.
├── cmd/
│   └── api/
│       └── main.go
├── internal/
│   ├── config/
│   ├── domain/
│   ├── handler/
│   ├── middleware/
│   ├── repository/
│   └── service/
├── pkg/
├── api/
├── migrations/
├── test/
├── docs/
├── scripts/
├── deployments/
├── go.mod
├── go.sum
├── Makefile
└── README.md
```

- Put executable entrypoints under `cmd/<app-name>/main.go`.
- Keep application-private code under `internal/`.
- Use `internal/config` for environment parsing and validation.
- Use `internal/domain` for core business models and interfaces.
- Use `internal/handler` for HTTP transport code.
- Use `internal/middleware` for HTTP middleware.
- Use `internal/service` for business use cases.
- Use `internal/repository` for persistence and external data access.
- Use `pkg/` only for code that is intentionally reusable by other modules.
- Use `api/` for OpenAPI specs, protobuf files, or HTTP contract documents.
- Use `migrations/` for database migrations.
- Use `test/` for integration fixtures and end-to-end test helpers.
- Use `docs/` for architecture notes and operational documentation.
- Use `scripts/` for local developer automation.
- Use `deployments/` for Docker, Kubernetes, Terraform, or other deployment
  assets.

Prefer small packages with clear ownership. Avoid generic package names such as
`utils`, `common`, or `helpers` unless there is no better domain-specific name.
Keep dependency direction simple: handlers call services, services depend on
domain interfaces, and repositories implement persistence details.

## Go Backend Practices

- Initialize new Go apps with `go mod init` using the repository module path.
- Keep `main.go` thin: load config, wire dependencies, start the server, and
  handle shutdown.
- Pass dependencies explicitly through constructors instead of relying on global
  mutable state.
- Use `context.Context` for request-scoped work, I/O, and cancellation.
- Return errors with context using `fmt.Errorf("...: %w", err)`.
- Prefer table-driven tests for services, handlers, and repositories.
- Run `gofmt` on edited Go files.
- Run `go test ./...` after changing Go code when dependencies are available.
- Add or update README run commands when adding a new service entrypoint.

## Subagent Boundaries

When a task is large enough to benefit from additional subagents, use focused
roles such as `planning-agent`, `backend-build-agent`, `test-agent`, and
`review-agent`.

Subagents should return concise findings or patches for the current step only.
The main agent remains responsible for final decisions, file edits, and user
review checkpoints.
