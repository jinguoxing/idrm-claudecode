# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

---

## Project Overview

**IDRM** (Intelligent Data Resource Management) is a Go-based microservices platform for data resource management. This project follows **Spec-Driven Development** with a strict 5-phase workflow.

**Tech Stack**: Go 1.21+, Go-Zero v1.9+, GORM, SQLx, MySQL 8.0, Redis 7.0, Kafka 3.0

---

## Development Workflow: 5-Phase Standard

All feature development MUST follow this workflow (defined in [spec/core/workflow.md](spec/core/workflow.md)):

```
Phase 0: Context    - Read project specs, understand architecture
Phase 1: Specify    - Write requirements.md with user stories & EARS
Phase 2: Design     - Create design.md with architecture & interfaces
Phase 3: Tasks      - Break into tasks <50 lines each
Phase 4: Implement  - Code, test, verify
```

**Key Principle**: Specs before code. Do not implement without completing phases 0-3.

---

## Architecture: 4-Layer Pattern

Strict layering (see [spec/architecture/layered-architecture.md](spec/architecture/layered-architecture.md)):

```
Handler (API Layer) → Logic (Business Layer) → Model (Persistence) → Database
```

- **Handler** (`api/internal/handler/`): HTTP only, parameter parsing, calls Logic
- **Logic** (`api/internal/logic/`): Business rules, validation, calls Model
- **Model** (`model/`): Data access via interface, GORM/SQLx implementations

**Dependency Rule**: Upper layers can depend on lower layers. Lower layers MUST NEVER depend on upper layers.

---

## Dual ORM Pattern

Model layer supports both GORM and SQLx via factory pattern (see [spec/architecture/dual-orm-pattern.md](spec/architecture/dual-orm-pattern.md)):

- **GORM**: Complex queries, relationships
- **SQLx**: Simple CRUD, performance-critical paths

Model directory structure:
```
model/{domain}/{entity}/
├── interface.go      # Unified interface
├── types.go          # Data structures
├── factory.go        # ORM factory
├── gorm_dao.go       # GORM implementation
└── sqlx_model.go     # SQLx implementation
```

---

## Common Commands

### Build & Test
```bash
# Build all
go build ./...

# Run tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run specific test
go test -v -run TestCreateCategory ./model/resource_catalog/category/
```

### Code Quality
```bash
# Lint
golangci-lint run

# Format
go fmt ./...
```

### Go-Zero Code Generation
```bash
# Generate API from .api file
goctl api go -api api/user.api -dir .

# Generate model from SQL
goctl model mysql datasource -url="user:pass@tcp(127.0.0.1:3306)/db" -table="*" -dir="./model"
```

---

## Development Standards

### Code Constraints
- Functions MUST be < 50 lines
- Comments MUST be in Chinese
- Error wrapping MUST use `%w`
- Test coverage MUST be > 80% for business logic

### Git Workflow
- `main` - production (stable)
- `develop` - development
- `feature/*` - feature branches
- Conventional Commits: `feat:`, `fix:`, `docs:`, `refactor:`

### Quality Gates
- ✅ Compile passes
- ✅ Tests pass (>80% coverage)
- ✅ Lint clean (golangci-lint)
- ✅ Peer review approved

---

## Feature Development Template

When creating a new feature, follow this structure:

```
specs/features/{feature-name}/
├── requirements.md   # Phase 1: User stories, EARS, business rules
├── design.md         # Phase 2: Architecture, interfaces, data model
└── tasks.md          # Phase 3: Implementation tasks <50 lines
```

Use templates from [spec/templates/](spec/templates/):
- `requirements.md` - Requirements template
- `design.md` - Design template
- `tasks.md` - Task breakdown template

---

## Key Packages

| Package | Purpose |
|---------|---------|
| `pkg/config/` | Multi-database config, Redis, Kafka, JWT |
| `pkg/errorx/` | Structured error codes (10000-49999) |
| `pkg/response/` | Standardized HTTP responses |
| `pkg/middleware/` | CORS, Auth, Recovery, RequestID, Tracing |
| `pkg/telemetry/` | Audit logging, distributed tracing |
| `pkg/validator/` | Struct validation with custom rules |

---

## Important Notes

1. **Specs are sacred** - All architecture patterns in `spec/` directory must be followed
2. **Interface-based design** - All cross-layer calls use interfaces for testability
3. **No business logic in Model** - Validation and rules go in Logic layer
4. **No HTTP in Logic** - Handler is the only layer that knows about HTTP
5. **Error handling** - Use `pkg/errorx` for structured errors, wrap with `%w`

---

## Reference Documentation

- [Project Charter](spec/core/project-charter.md) - Project mission and values
- [Workflow](spec/core/workflow.md) - Complete 5-phase workflow definition
- [Layered Architecture](spec/architecture/layered-architecture.md) - Detailed layering rules
- [Dual ORM Pattern](spec/architecture/dual-orm-pattern.md) - GORM/SQLx strategy
- [API Design Guide](spec/architecture/api-design-guide.md) - RESTful API standards

---

**This project prioritizes quality over speed. When in doubt, refer to the specs.**
