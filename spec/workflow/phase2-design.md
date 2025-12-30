# Phase 2: Design (技术方案)

> **Level**: MUST  
> **Version**: 2.1.0  
> **Updated**: Includes Technical Constraints and Data Model

## 目标

基于Phase 1的业务需求，创建详细的技术实现方案。

## 输入
- Phase 1的requirements.md（业务需求）

## 输出模板

创建 `specs/features/{name}/design.md`：

```markdown
# Design: {功能名}

## Architecture Overview
遵循IDRM分层架构：Handler → Logic → Model

## ORM Strategy
GORM vs SQLx选择及理由

## File Structure
详细的文件清单和目录结构

## Interface Definitions
\`\`\`go
type Model interface {
    // 方法定义
}
\`\`\`

## Sequence Diagrams
主要流程的序列图

## Technical Constraints

### IDRM通用约束
- 架构约束
- 编码约束  
- 质量约束

### 功能特定约束
- ORM选择理由
- 性能要求
- 安全考虑

## Data Model

### 表结构
\`\`\`sql
CREATE TABLE ...
\`\`\`

### 字段说明
详细的字段定义和约束

### Go结构体
\`\`\`go
type Entity struct { ... }
\`\`\`

## Error Handling
错误定义和处理策略

## Testing Strategy
测试方法和Mock方案
```

## 各部分说明

### Architecture Overview
- 分层架构设计
- 组件职责划分
- 依赖关系

### Sequence Diagrams（在约束之前）
- 主要业务流程图
- 异常处理流程
- 系统交互图

### Technical Constraints

**IDRM通用约束**（必须包含）：
```markdown
### IDRM通用约束
- MUST follow layered architecture (Handler → Logic → Model)
- MUST implement dual ORM support
- Functions MUST be < 50 lines
- MUST use Chinese comments for all public items
- Error wrapping MUST use %w
- Test coverage MUST be > 80%
- MUST pass golangci-lint check
```

**功能特定约束**：
- ORM选择（GORM/SQLx）及理由
- 性能要求（响应时间、并发）
- 特殊的技术限制

### Data Model

**DDL 文件位置**: `migrations/{module}/{table}.sql`

每个功能需要输出独立的 DDL 文件，用于 `goctl model` 代码生成。

**DDL 格式要求**:
```sql
-- migrations/resource_catalog/tags.sql
CREATE TABLE `tags` (
    `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '标签ID',
    `name` varchar(50) NOT NULL COMMENT '标签名称',
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted_at` datetime DEFAULT NULL COMMENT '删除时间(软删除)',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='标签表';
```

**关键检查点**:
- 完整的表结构定义
- 字段详细说明（COMMENT）
- Go结构体定义
- 索引设计

### API Contract (go-zero)

每个功能需要定义 go-zero 格式的 API 文件，作为前后端接口契约。

**文件位置**: `specs/features/{name}/{name}.api`

**格式要求**:
```api
syntax = "v1"

info(
    title: "{Feature} API"
    desc: "功能描述"
    version: "v1"
)

type (
    CreateXxxReq {
        Name string `json:"name"`
    }
    CreateXxxResp {
        Id int64 `json:"id"`
    }
)

@server(
    prefix: /api/v1
    group: {feature}
)
service idrm-api {
    @handler CreateXxx
    post /xxx (CreateXxxReq) returns (CreateXxxResp)
}
```

**关键检查点**:
- 遵循 RESTful 规范
- 请求/响应类型定义完整
- 分组和前缀符合项目约定

## AI工具使用

**Kiro.dev**：
```
基于requirements.md
自动生成design.md框架
```

**Cursor**：
```
@requirements.md
@spec/architecture/layered-architecture.md
@spec/architecture/dual-orm-pattern.md

Phase 2: 生成Design

要求：
1. 完整的架构设计
2. 序列图
3. 技术约束
4. 数据模型
```

## 质量门禁 (Gate 2)

- [ ] 符合分层架构
- [ ] ORM选择合理
- [ ] 文件清单完整
- [ ] 接口定义清晰
- [ ] 序列图完整
- [ ] 技术约束明确
- [ ] 数据模型完整
- [ ] **DDL 文件 (.sql) 定义完整**
- [ ] **API Contract (.api) 定义完整**

参考：`../quality/quality-gates.md#gate-2`

## 下一步
→ Phase 3: Tasks
