# Kiro.dev 使用指南

> **文档版本**: v1.0  
> **最后更新**: 2025-12-26  
> **适用于**: IDRM项目  
> **官方文档**: https://kiro.dev/docs/specs/

---

## 📖 概述

本文档介绍如何使用Kiro.dev IDE配合IDRM项目的spec规范进行AI辅助开发。

Kiro.dev是一个智能IDE，提供**结构化规范(Specs)**功能，通过Requirements → Design → Implementation三阶段工作流，将高层想法转化为详细的实施计划。

---

## ⚠️ 重要说明

### 关于规范文档路径

本指南中所有涉及的规范文档路径（如 `private_doc/spec/`）都是基于**IDRM项目的实际目录结构**。

**如果您在其他项目中使用本指南**，需要根据您项目的实际情况调整：

- **规范文档目录**：可能是 `docs/specs/`、`.github/specs/` 或其他路径
- **文档文件名**：对应您项目的实际文件名
- **Steering配置**：更新引用路径

---

## 🎯 Kiro.dev Specs核心概念

### 三阶段工作流

```
Phase 1: Requirements
  → 用户故事 + 验收标准 (EARS notation)
  → 文件：requirements.md

Phase 2: Design  
  → 技术架构 + 序列图 + 实现考虑
  → 文件：design.md

Phase 3: Implementation
  → 离散任务 + 实时进度追踪
  → 文件：tasks.md
```

### EARS Notation

Kiro使用EARS（Easy Approach to Requirements Syntax）规范编写需求：

**格式**：
```
WHEN [condition/event] THE SYSTEM SHALL [expected behavior]
```

**示例**：
```
WHEN a user submits a form with invalid data
THE SYSTEM SHALL display validation errors next to the relevant fields
```

**优势**：
- ✅ 清晰明确
- ✅ 可测试
- ✅ 可追踪
- ✅ 完整性

---

## 🚀 快速开始

### 1. 安装Kiro.dev

```bash
# 访问官网下载
https://kiro.dev/downloads/

# macOS安装后
# 应用程序 → Kiro.app
```

### 2. 打开IDRM项目

```bash
# 方式1：命令行
kiro /Users/kingnet/code_workspace/go_workspace/src/idrm

# 方式2：IDE界面
File → Open → 选择idrm目录
```

### 3. 配置Steering（项目规则）

**Steering是Kiro的规则引擎，类似于Cursor的.cursorrules**

创建 `.kiro/steering/idrm-rules.md`:

```bash
mkdir -p .kiro/steering
```

内容：

```markdown
# IDRM Project Steering Rules

## 项目上下文

**项目名称**: IDRM (Intelligent Data Resource Management)  
**技术栈**: Go 1.21+, Go-Zero v1.9+, MySQL 8.0, Redis, Kafka  
**规范位置**: private_doc/spec/

## 核心原则

参考：`private_doc/spec/constitution.md`

1. **4阶段工作流** (适配Kiro的3阶段)
   - Specify → Kiro Requirements
   - Plan → Kiro Design
   - Tasks → Kiro Implementation (前半)
   - Implement → Kiro Implementation (后半)

2. **分层架构** (严格执行)
   - Handler → Logic → Model
   - 禁止跨层调用
   - 参考：`private_doc/spec/architecture/layered-architecture.md`

3. **双ORM模式**
   - 同时支持GORM和SQLx
   - 通过Factory模式选择
   - 参考：`private_doc/spec/architecture/dual-orm-pattern.md`

## Requirements阶段规则

生成requirements.md时：

1. **使用EARS notation** 编写验收标准
2. **包含技术约束** 来自IDRM constitution
3. **明确数据模型** 遵循table-based结构

示例格式：
\`\`\`markdown
## User Stories
AS a [role]
I WANT [feature]
SO THAT [benefit]

## Acceptance Criteria (EARS)
WHEN [condition] THE SYSTEM SHALL [behavior]

## Technical Constraints (IDRM)
- MUST follow layered architecture
- MUST implement dual ORM
- MUST use table-based directory structure
\`\`\`

## Design阶段规则

生成design.md时：

1. **严格分层**
   - Handler: `api/internal/handler/{module}/{table}/`
   - Logic: `api/internal/logic/{module}/{table}/`
   - Model: `model/{module}/{table}/`

2. **Model结构** (Table-based Pattern)
   \`\`\`
   model/{module}/{table}/
   ├── interface.go    # Model接口
   ├── types.go        # 数据结构
   ├── vars.go         # 常量错误
   ├── factory.go      # ORM工厂
   ├── gorm_dao.go     # GORM实现
   └── sqlx_model.go   # SQLx实现
   \`\`\`

3. **接口设计**
   - 使用统一的Model接口
   - 支持事务（WithTx, Trans）
   - 完整的CRUD方法

## Implementation阶段规则

生成tasks.md和执行时：

1. **编码规范** (参考：`private_doc/spec/coding-standards/`)
   - 所有公开项必须有中文注释
   - 函数长度 < 50行
   - 错误封装使用 %w
   - 导入分组：stdlib → third-party → internal

2. **命名规范**
   - 文件: 全小写下划线 (create_category_logic.go)
   - 包: 全小写简短 (category, middleware)
   - 类型: 大驼峰 (CategoryModel, UserInfo)
   - 函数: 公开大驼峰，私有小驼峰

3. **任务拆分**
   - 每个任务 < 50行代码
   - 明确依赖关系
   - 清晰的验收标准

## 质量标准

所有生成的代码必须：

- ✅ 编译成功 (`go build ./...`)
- ✅ 测试通过 (`go test ./...`)
- ✅ Lint无错误 (`golangci-lint run`)
- ✅ 符合分层架构
- ✅ 包含完整注释

## 参考文档

在生成代码时，优先参考以下文档：

- **总体规范**: `private_doc/spec/constitution.md`
- **分层架构**: `private_doc/spec/architecture/layered-architecture.md`
- **双ORM模式**: `private_doc/spec/architecture/dual-orm-pattern.md`
- **API设计**: `private_doc/spec/architecture/api-design-guide.md`
- **Go风格**: `private_doc/spec/coding-standards/go-style-guide.md`
- **命名规范**: `private_doc/spec/coding-standards/naming-conventions.md`
- **错误处理**: `private_doc/spec/coding-standards/error-handling.md`
- **测试规范**: `private_doc/spec/coding-standards/testing-standards.md`

## 禁止操作

❌ **绝对禁止**：
1. 在Handler中实现业务逻辑
2. 在Logic中直接访问数据库
3. 在Model中包含业务逻辑
4. 跨层调用
5. 忽略错误处理
6. 生成没有注释的代码
```

### 4. 验证配置

在Kiro中打开Chat，询问：

```
请总结IDRM项目的架构规范和编码要求
```

如果Kiro能准确回答分层架构、双ORM等信息，说明steering配置成功。

---

## 💡 创建第一个Spec

### 方式1：通过UI创建

1. 打开Kiro侧边栏
2. 点击"Specs"下的 `+` 按钮
3. 或在Chat中输入"创建新Spec"

### 方式2：通过Chat创建

在Kiro Chat中输入：

```
创建新的Spec: Directory Management

功能描述：
添加directory表，用于管理资源目录，支持树形结构。

主要功能：
- 创建目录
- 查询目录（支持分页）
- 更新目录信息
- 删除目录（需检查是否有子目录）
- 支持父子关系（level字段）

请开始Phase 1: Requirements
```

---

## 📋 三阶段详细流程

### Phase 1: Requirements（需求定义）

**目标**：创建清晰的需求文档，使用EARS notation

**Kiro操作**：

```
在Spec中，Kiro会引导你完成：

1. 描述功能概述
2. 定义用户故事
3. 编写验收标准（EARS格式）
4. 明确技术约束（来自IDRM steering）
```

**生成的requirements.md示例**：

```markdown
# Feature: Directory Management

## Overview
实现资源目录管理功能，支持树形结构组织资源。

## User Stories

### Story 1: Create Directory
AS a resource manager
I WANT to create a new directory
SO THAT I can organize resources hierarchically

### Story 2: Query Directory
AS a user
I WANT to query directories with pagination
SO THAT I can efficiently browse large directory trees

### Story 3: Update Directory
AS a resource manager
I WANT to update directory information
SO THAT I can maintain accurate metadata

### Story 4: Delete Directory
AS a resource manager  
I WANT to delete a directory safely
SO THAT I don't accidentally remove directories with content

## Acceptance Criteria (EARS Notation)

### Create Directory
WHEN a user submits a create directory request
THE SYSTEM SHALL validate the name is unique within the parent

WHEN a user creates a root directory (parent_id is null)
THE SYSTEM SHALL set level to 0

WHEN a user creates a child directory
THE SYSTEM SHALL set level to parent.level + 1

WHEN the directory name is empty
THE SYSTEM SHALL return a validation error

### Query Directory
WHEN a user requests directory list
THE SYSTEM SHALL return paginated results

WHEN a user filters by parent_id
THE SYSTEM SHALL return only direct children

WHEN a user requests a specific directory by ID
THE SYSTEM SHALL return full directory details or 404

### Update Directory
WHEN a user updates a directory
THE SYSTEM SHALL validate the new data

WHEN a user moves a directory (changes parent_id)
THE SYSTEM SHALL check for circular references

WHEN a user renames a directory
THE SYSTEM SHALL check name uniqueness within new parent

### Delete Directory
WHEN a user deletes a directory
THE SYSTEM SHALL check if it has children

WHEN a directory has children
THE SYSTEM SHALL return an error preventing deletion

WHEN a directory is empty
THE SYSTEM SHALL soft-delete (set deleted_at)

## Technical Constraints (IDRM)

### Architecture
- MUST follow layered architecture (Handler → Logic → Model)
- MUST NOT have cross-layer dependencies

### Data Layer
- MUST use table-based directory structure
- MUST implement both GORM and SQLx
- MUST use unified Model interface

### Code Standards
- ALL public items MUST have Chinese comments
- Functions MUST be < 50 lines
- Error wrapping MUST use %w

## Data Model

### Directory Table
\`\`\`sql
CREATE TABLE directories (
    id BIGINT PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    code VARCHAR(50) NOT NULL,
    parent_id BIGINT REFERENCES directories(id),
    level INT NOT NULL DEFAULT 0,
    path VARCHAR(500),
    status TINYINT DEFAULT 1,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    UNIQUE KEY uk_parent_name (parent_id, name)
);
\`\`\`

### Fields
- id: 唯一标识（雪花算法）
- name: 目录名称
- code: 目录编码
- parent_id: 父目录ID（NULL表示根目录）
- level: 层级（0表示根）
- path: 完整路径（如：/root/child1/child2）
- status: 状态（0-禁用，1-启用）

## Open Questions

1. level最大深度限制？建议：10层
2. 删除级联策略？建议：禁止删除有子节点的目录
3. 权限控制？待确认

## Next Steps

→ Phase 2: Design
```

**你的操作**：
1. Review requirements.md
2. 补充未确定项
3. 确认后进入Design阶段

### Phase 2: Design（技术设计）

**目标**：创建技术架构文档

**Kiro操作**：

基于requirements.md，Kiro会生成：
1. 系统架构设计
2. 文件结构
3. 接口定义
4. 序列图
5. 数据流

**生成的design.md示例**：

```markdown
# Design: Directory Model

## Architecture Overview

遵循IDRM分层架构规范：
- 参考：`private_doc/spec/architecture/layered-architecture.md`

### Layer Separation

\`\`\`
┌─────────────────────────────────┐
│   Handler Layer                 │
│   api/internal/handler/         │
│   resource_catalog/directory/   │
└─────────────────────────────────┘
            ↓
┌─────────────────────────────────┐
│   Logic Layer                   │
│   api/internal/logic/           │
│   resource_catalog/directory/   │
└─────────────────────────────────┘
            ↓
┌─────────────────────────────────┐
│   Model Layer                   │
│   model/resource_catalog/       │
│   directory/                    │
└─────────────────────────────────┘
\`\`\`

## File Structure

### Model Layer (Table-based Pattern)

参考：`private_doc/spec/architecture/dual-orm-pattern.md`

\`\`\`
model/resource_catalog/directory/
├── interface.go       # Model接口定义
├── types.go          # Directory数据结构
├── vars.go           # 常量和错误定义
├── factory.go        # ORM工厂函数
├── gorm_dao.go       # GORM实现
└── sqlx_model.go     # SQLx实现
\`\`\`

### Logic Layer

\`\`\`
api/internal/logic/resource_catalog/directory/
├── createdirectorylogic.go
├── getdirectorylogic.go
├── listdirectorylogic.go
├── updatedirectorylogic.go
└── deletedirectorylogic.go
\`\`\`

### Handler Layer

\`\`\`
api/internal/handler/resource_catalog/directory/
├── createdirectoryhandler.go
├── getdirectoryhandler.go
├── listdirectoryhandler.go
├── updatedirectoryhandler.go
└── deletedirectoryhandler.go
\`\`\`

## Interface Definitions

### Model Interface

\`\`\`go
package directory

import "context"

// Model 目录仓储接口
// 支持GORM和SQLx双ORM实现
type Model interface {
    // Insert 插入新目录
    Insert(ctx context.Context, data *Directory) (*Directory, error)
    
    // FindOne 根据ID查询目录
    FindOne(ctx context.Context, id int64) (*Directory, error)
    
    // FindByCode 根据code查询
    FindByCode(ctx context.Context, code string) (*Directory, error)
    
    // FindByParent 查询父目录下的所有子目录
    FindByParent(ctx context.Context, parentId int64) ([]*Directory, error)
    
    // List 分页查询
    List(ctx context.Context, page, pageSize int) ([]*Directory, int64, error)
    
    // Update 更新目录
    Update(ctx context.Context, data *Directory) error
    
    // SoftDelete 软删除
    SoftDelete(ctx context.Context, id int64) error
    
    // CheckCircular 检查循环引用
    CheckCircular(ctx context.Context, id, newParentId int64) (bool, error)
    
    // 事务支持
    WithTx(tx interface{}) Model
    Trans(ctx context.Context, fn func(ctx context.Context, model Model) error) error
}
\`\`\`

### Data Types

\`\`\`go
// Directory 目录数据结构
type Directory struct {
    Id        int64      `gorm:"column:id;primaryKey" db:"id"`
    Name      string     `gorm:"column:name" db:"name"`
    Code      string     `gorm:"column:code" db:"code"`
    ParentId  *int64     `gorm:"column:parent_id" db:"parent_id"`
    Level     int        `gorm:"column:level" db:"level"`
    Path      string     `gorm:"column:path" db:"path"`
    Status    int8       `gorm:"column:status" db:"status"`
    CreatedAt time.Time  `gorm:"column:created_at" db:"created_at"`
    UpdatedAt time.Time  `gorm:"column:updated_at" db:"updated_at"`
    DeletedAt *time.Time `gorm:"column:deleted_at" db:"deleted_at"`
}

func (Directory) TableName() string {
    return "directories"
}
\`\`\`

## Sequence Diagrams

### Create Directory Flow

\`\`\`
User → Handler: POST /api/v1/catalog/directories
Handler → Logic: CreateDirectory(req)
Logic → Logic: validateInput()
Logic → Model: FindByCode(code)
Model → DB: SELECT
DB → Model: result
Logic → Model: Insert(directory)
Model → DB: INSERT
DB → Model: success
Model → Logic: directory
Logic → Handler: resp
Handler → User: 200 OK
\`\`\`

### Delete Directory with Children Check

\`\`\`
User → Handler: DELETE /api/v1/catalog/directories/:id
Handler → Logic: DeleteDirectory(id)
Logic → Model: FindByParent(id)
Model → DB: SELECT WHERE parent_id = ?
DB → Model: children
alt has children
    Logic → Handler: error
    Handler → User: 400 Bad Request
else no children
    Logic → Model: SoftDelete(id)
    Model → DB: UPDATE SET deleted_at = NOW()
    DB → Model: success
    Logic → Handler: success
    Handler → User: 200 OK
end
\`\`\`

## Implementation Considerations

### 循环引用检查

移动目录时需要检查：
1. 新父目录不能是自己
2. 新父目录不能是自己的子孙节点

算法：
\`\`\`go
func (d *DirectoryDao) CheckCircular(ctx context.Context, id, newParentId int64) (bool, error) {
    // 检查newParentId的path是否包含id
    parent, err := d.FindOne(ctx, newParentId)
    if err != nil {
        return false, err
    }
    
    selfDir, err := d.FindOne(ctx, id)
    if err != nil {
        return false, err
    }
    
    // 如果parent的path包含selfDir的path，说明有循环
    return strings.Contains(parent.Path, selfDir.Path), nil
}
\`\`\`

### Path维护

创建/移动目录时自动更新path：
\`\`\`go
func buildPath(parentPath string, name string) string {
    if parentPath == "" {
        return "/" + name
    }
    return parentPath + "/" + name
}
\`\`\`

### 性能优化

1. 索引：`(parent_id, name)` 唯一索引
2. Level字段：快速过滤深度
3. Path字段：快速查询祖先路径

## Dependencies

- 依赖：雪花算法ID生成 (`pkg/utils/sonyflake.go`)
- 依赖：统一响应格式 (`pkg/response/`)
- 依赖：中间件（RequestID, Trace, Logger）

## Next Steps

→ Phase 3: Implementation
```

**你的操作**：
1. Review设计文档
2. 确认架构符合IDRM规范
3. 确认后进入Implementation阶段

### Phase 3: Implementation（任务执行）

**目标**：拆分任务并逐个实施

**Kiro操作**：

基于design.md生成可执行的任务清单

**生成的tasks.md示例**：

```markdown
# Implementation Tasks: Directory Model

## Task 1: 创建Model基础结构
**Status**: ⏸️ Not Started

### Files
- `model/resource_catalog/directory/interface.go`
- `model/resource_catalog/directory/types.go`
- `model/resource_catalog/directory/vars.go`

### Description
创建Model层的基础接口、数据类型和常量定义

### Acceptance Criteria
- Model接口包含所有CRUD方法
- Directory结构体完整定义
- 错误常量定义完整
- 所有公开项有中文注释

### Estimated Lines: 150

---

## Task 2: 实现Factory模式
**Status**: ⏸️ Not Started
**Depends on**: Task 1

### Files
- `model/resource_catalog/directory/factory.go`

### Description
实现ORM工厂函数，支持GORM和SQLx切换

### Acceptance Criteria
- NewModel函数优先选择GORM
- GORM不可用时降级到SQLx
- 两者都不可用时panic with clear message
- 工厂注册机制完整

### Estimated Lines: 50

---

## Task 3: 实现GORM DAO
**Status**: ⏸️ Not Started
**Depends on**: Task 1, Task 2

### Files
- `model/resource_catalog/directory/gorm_dao.go`

### Description
实现GORM版本的Model接口

### Acceptance Criteria
- 实现所有Model接口方法
- Insert包含ID生成
- FindByParent查询正确
- CheckCircular逻辑正确
- Trans事务支持完整
- 错误处理使用%w封装

### Estimated Lines: 250

---

## Task 4: 实现SQLx Model
**Status**: ⏸️ Not Started
**Depends on**: Task 1, Task 2

### Files
- `model/resource_catalog/directory/sqlx_model.go`

### Description
实现SQLx版本的Model接口

### Acceptance Criteria
- 实现所有Model接口方法
- SQL语句正确且安全
- 参数绑定使用占位符
- 事务支持完整

### Estimated Lines: 280

---

## Task 5: 更新ServiceContext
**Status**: ⏸️ Not Started
**Depends on**: Task 1, Task 2

### Files
- `api/internal/svc/servicecontext.go`

### Description
在ServiceContext中添加DirectoryModel

### Acceptance Criteria
- 添加DirectoryModel字段
- 在NewServiceContext中初始化
- 使用factory.NewModel创建
- 添加blank imports触发工厂注册

### Estimated Lines: 10

---

## Task 6: 创建Logic层
**Status**: ⏸️ Not Started
**Depends on**: Task 5

### Files
- `api/internal/logic/resource_catalog/directory/createdirectorylogic.go`
- `api/internal/logic/resource_catalog/directory/getdirectorylogic.go`
- `api/internal/logic/resource_catalog/directory/listdirectorylogic.go`
- `api/internal/logic/resource_catalog/directory/updatedirectorylogic.go`
- `api/internal/logic/resource_catalog/directory/deletedirectorylogic.go`

### Description
实现所有业务逻辑

### Acceptance Criteria
- CreateDirectoryLogic: 验证、构建path、调用Model
- GetDirectoryLogic: 调用Model.FindOne
- ListDirectoryLogic: 分页参数处理、调用Model.List
- UpdateDirectoryLogic: 验证、检查循环引用、调用Model.Update
- DeleteDirectoryLogic: 检查子节点、调用Model.SoftDelete
- 所有Logic包含完整的error handling和logging

### Estimated Lines: 400

---

## Task 7: 创建Handler层
**Status**: ⏸️ Not Started
**Depends on**: Task 6

### Files
- `api/internal/handler/resource_catalog/directory/*.go`

### Description
创建所有HTTP Handlers

### Acceptance Criteria
- Handler只负责参数解析和响应格式化
- 调用对应的Logic
- 使用统一response格式
- HTTP状态码正确

### Estimated Lines: 200

---

## Task 8: 定义API接口
**Status**: ⏸️ Not Started
**Depends on**: Task 7

### Files
- `api/doc/api/resource_catalog/directory.api`

### Description
定义RESTful API接口

### Acceptance Criteria
- 定义所有CRUD接口
- Request/Response types定义完整
- 路由路径符合REST规范
- 包含参数验证规则

### Estimated Lines: 150

---

## Task 9: 生成代码并集成
**Status**: ⏸️ Not Started
**Depends on**: Task 8

### Commands
\`\`\`bash
# 生成API代码
goctl api go -api api/doc/api/resource_catalog/directory.api -dir api/

# 验证编译
go build ./...

# 运行测试
go test ./model/resource_catalog/directory/...
\`\`\`

### Acceptance Criteria
- goctl生成成功
- 编译无错误
- lint通过

---

## Task 10: 编写单元测试
**Status**: ⏸️ Not Started
**Depends on**: Task 3, Task 4

### Files
- `model/resource_catalog/directory/gorm_dao_test.go`
- `model/resource_catalog/directory/sqlx_model_test.go`

### Description
为Model层编写单元测试

### Acceptance Criteria
- 使用表驱动测试
- 覆盖正常和异常情况
- Mock数据库连接
- 测试覆盖率 >80%

### Estimated Lines: 300

---

## Summary
- Total Tasks: 10
- Total Estimated Lines: ~2000
- Dependencies: 有依赖关系，需按顺序执行
```

**执行任务**：

在Kiro中，每个Task都是可点击的：

1. 点击"Task 1"
2. Kiro生成对应的代码
3. Review代码
4. 应用到项目
5. Task状态更新为"✅ Completed"
6. 继续下一个Task

---

## 🎯 实际操作演示

### 完整流程：添加Directory功能

**Step 1: 创建Spec**

在Kiro Chat中：
```
创建新的Spec: Directory Management

参考IDRM steering规则，实现资源目录管理功能。
```

**Step 2: Requirements Phase**

Kiro引导你定义：
- 用户故事
- EARS notation验收标准
- 技术约束（自动引用IDRM规范）
- 数据模型

Review后确认。

**Step 3: Design Phase**

Kiro生成：
- 分层架构设计
- 文件清单
- 接口定义
- 序列图

Review后确认。

**Step 4: Implementation Phase**

Kiro生成tasks.md，逐个执行：

```
点击Task 1 → Kiro生成interface.go, types.go, vars.go
Review → Apply → ✅
点击Task 2 → Kiro生成factory.go  
Review → Apply → ✅
...依次执行
```

**Step 5: Verification**

```bash
# 编译检查
go build ./...

# 运行测试
go test ./...

# Lint检查
golangci-lint run
```

---

## 🔧 高级功能

### 1. 引用现有代码

在Chat中：
```
参考已有实现：
- model/resource_catalog/category/

为directory创建相同的结构
```

Kiro会分析category的实现并复制模式。

### 2. 更新现有Spec

```
打开现有Spec: Directory Management
修改requirements.md
Kiro自动更新design.md和tasks.md
```

### 3. 追踪进度

在Kiro Specs面板查看：
- 所有specs列表
- 每个spec的阶段状态
- tasks完成进度

### 4. 协作和分享

```
Export Spec → 生成markdown文件
Share → 与团队成员共享
```

---

## 📊 工具对比

| 功能 | Kiro.dev | Cursor | Claude CLI |
|------|----------|--------|-----------|
| **核心优势** | 结构化Specs | 实时补全 | 批处理 |
| **工作流** | 3阶段 | 自由对话 | 命令行 |
| **任务追踪** | ⭐⭐⭐ 内置 | ⭐ 无 | ⭐ 无 |
| **EARS规范** | ⭐⭐⭐ 原生 | ❌ 不支持 | ❌ 不支持 |
| **可视化** | ⭐⭐⭐ 强大 | ⭐⭐ 好 | ❌ 无 |
| **IDE集成** | ⭐⭐⭐ 独立IDE | ⭐⭐⭐ Fork VSCode | ❌ CLI |
| **规范集成** | ⭐⭐⭐ Steering | ⭐⭐ .cursorrules | ⭐⭐ .clinerules |

### 推荐使用场景

**Kiro.dev**：
- ✅ 结构化大型功能开发
- ✅ 需要清晰的需求文档
- ✅ 团队协作和追踪
- ✅ 喜欢3阶段工作流

**Cursor**：
- ✅ 日常快速编码
- ✅ 代码补全和重构
- ✅ 快速原型开发

**Claude CLI**：
- ✅ 批量处理
- ✅ CI/CD集成
- ✅ 脚本化任务

### 组合使用

```
Kiro.dev (规划和结构)
  → Requirements, Design, Tasks
  
Cursor (快速实施)
  → 基于tasks快速编码
  
Claude CLI (验证和批处理)
  → Code review
  → 批量测试生成
```

---

## 🎯 最佳实践

### DO ✅

1. **完整配置Steering**
   - 引用所有IDRM规范
   - 明确架构约束
   - 定义编码标准

2. **遵循3阶段流程**
   - 不跳过Requirements
   - Design要详细
   - Tasks要细分

3. **使用EARS notation**
   - 清晰的验收标准
   - 便于测试

4. **充分利用可视化**
   - 查看spec进度
   - 追踪task状态

5. **Review每个Task**
   - 不盲目应用代码
   - 确保符合规范

### DON'T ❌

1. ❌ **不要跳过Steering配置**
   - 没有规则Kiro无法遵循IDRM规范

2. ❌ **不要跳过阶段**
   - 3阶段是有逻辑的

3. ❌ **不要盲目应用代码**
   - 必须Review

4. ❌ **不要忽略依赖关系**
   - Tasks有顺序

5. ❌ **不要重复定义规范**
   - Steering引用即可

---

## 🔍 常见问题

### Q: Kiro如何读取IDRM规范？

**A**: 通过Steering配置引用。Kiro会读取指定的规范文档并在生成时遵循。

### Q: 可以修改生成的requirements.md吗？

**A**: 可以。手动编辑后，Kiro会基于新的requirements更新design和tasks。

### Q: Task的顺序可以调整吗？

**A**: 可以手动重排tasks.md，但要注意依赖关系。

### Q: 如何与团队共享Specs？

**A**: Specs文件在`.kiro/specs/`目录，可以提交到Git。

### Q: Kiro支持中文吗？

**A**: 支持。可以用中文编写requirements和任务描述。

---

## 📚 参考资源

### IDRM规范

- `private_doc/spec/constitution.md` - 项目宪章
- `private_doc/spec/architecture/` - 架构规范
- `private_doc/spec/coding-standards/` - 编码规范

### Kiro.dev官方

- [Specs文档](https://kiro.dev/docs/specs/)
- [Specs概念](https://kiro.dev/docs/specs/concepts/)
- [最佳实践](https://kiro.dev/docs/specs/best-practices/)
- [下载](https://kiro.dev/downloads/)

### 配置文件

- `.kiro/steering/idrm-rules.md` - 项目规则
- `.kiro/specs/` - 功能specs

---

## ✨ 总结

### Kiro.dev核心优势

1. **结构化工作流** - Requirements → Design → Implementation
2. **EARS规范** - 清晰的验收标准
3. **任务追踪** - 实时进度可视化
4. **与IDRM集成** - 通过Steering引用规范

### 推荐工作流

```
早上：
→ 在Kiro中review昨天的specs
→ 更新task状态

开发新功能：
→ 创建新Spec
→ Phase 1: Requirements (EARS)
→ Phase 2: Design (引用IDRM architecture)
→ Phase 3: Implementation (逐个执行tasks)

Code Review：
→ 使用Cursor或Claude CLI
→ 参考code-review-checklist.md

提交：
→ Specs文件随代码一起提交
→ 保持docs同步
```

---

**Kiro.dev + IDRM Spec = 结构化的规范驱动开发！** 🚀
