# Design: {{feature_name}}

## Architecture Overview
<!-- Follows IDRM Layered Architecture: Handler -> Logic -> Model -->

## ORM Strategy
<!-- GORM vs SQLx selection and reasoning -->

## File Structure
<!-- Detailed file list and directory structure -->

## Interface Definitions
```go
type Model interface {
    // Method definitions
}
```

## Sequence Diagrams
<!-- Mermaid sequence diagrams for main flows -->

## Technical Constraints
### IDRM General Constraints
- MUST follow layered architecture (Handler → Logic → Model)
- MUST implement dual ORM support
- Functions MUST be < 50 lines
- MUST use Chinese comments
- Error wrapping MUST use %w
- Test coverage MUST be > 80%

### Feature Specific Constraints
<!-- Performance, security, specific limitations -->

## Data Model
### Table Schema
```sql
CREATE TABLE ...
```

### Go Structs
```go
type Entity struct { ... }
```

## Error Handling
<!-- Error definitions and strategies -->

## Testing Strategy
<!-- Test approach and Mock strategy -->

## API Contract (go-zero)
见 `{{feature_name}}.api` 文件

<!-- 生成命令: speckit generate api {{feature_name}} -->
