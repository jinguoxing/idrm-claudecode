# IDRM Steering Rules for Kiro.dev

## Project
IDRM - Intelligent Data Resource Management

## Tech Stack
- Go 1.21+, Go-Zero v1.9+
- MySQL 8.0, Redis, Kafka
- GORM + SQLx (Dual ORM)

## Workflow
5-Phase: Context → Specify (EARS) → Design → Tasks → Implement

参考：`private_doc/spec/core/workflow.md`

## Architecture
Layered: Handler → Logic → Model

参考：`private_doc/spec/architecture/layered-architecture.md`

## Requirements Phase
- Use EARS notation
- Include technical constraints from specs
- Reference: `private_doc/spec/workflow/phase1-specify.md`

## Design Phase
- Follow layered architecture
- Dual ORM support
- Reference: `private_doc/spec/workflow/phase2- design.md`

## Implementation Phase
- Functions < 50 lines
- Chinese comments
- Error wrapping with %w
- Reference: `private_doc/spec/workflow/phase4-implement.md`

## Quality Standards
- Build pass: `go build ./...`
- Test >80%: `go test -cover ./...`
- Lint clean: `golangci-lint run`

参考：`private_doc/spec/quality/quality-gates.md`

## Documentation
All specs in: `private_doc/spec/`
