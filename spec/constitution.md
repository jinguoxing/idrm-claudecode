# IDRM 项目宪章 (Project Constitution)

> **版本**：v1.0  
> **更新日期**：2025-12-24  
> **状态**：执行中  
> **维护者**：IDRM开发团队

---

## 📖 项目概览

### 项目使命

IDRM (Intelligent Data Resource Management) 是一个智能化的数据资源管理平台，旨在提供高效、可靠、可扩展的数据资源治理能力，帮助企业更好地管理和利用数据资产。

### 核心价值观

1. **质量优先** - 代码质量重于开发速度，可维护性优于便利性
2. **用户至上** - 以终端用户体验为中心，持续优化产品
3. **持续演进** - 拥抱变化，持续改进技术和流程
4. **团队协作** - 知识共享，共同成长，相互赋能
5. **技术卓越** - 追求最佳实践和技术创新

### 技术愿景

构建一个**高性能、高可用、易扩展**的数据资源管理平台，通过现代化的微服务架构、完善的可观测性体系和AI辅助开发，实现快速迭代和高质量交付。

---

## 🏗️ 技术原则

### 1. 架构原则

#### 1.1 分层架构

**严格遵循四层架构**：

```
┌─────────────────────────────────┐
│     API Layer (Go-Zero)         │  HTTP接口、参数验证
│  Handler → Logic → Model        │
└─────────────────────────────────┘
             ↓
┌─────────────────────────────────┐
│      Model Layer (Dual ORM)     │  数据访问抽象
│  Interface → Factory → Impl     │
└─────────────────────────────────┘
             ↓
┌─────────────────────────────────┐
│     Infrastructure Layer        │  基础设施组件
│  MySQL / Redis / Kafka          │
└─────────────────────────────────┘
```

**约束**：
- ✅ 上层可依赖下层，下层不可依赖上层
- ✅ 每层专注自己的职责，不可越界
- ✅ 层间通过明确的接口通信

#### 1.2 接口抽象

**面向接口编程**：
- 所有跨层调用必须通过接口
- 依赖抽象而非具体实现（依赖倒置原则）
- 便于单元测试和Mock

**示例**：
```go
// ✅ 正确：使用接口
type ServiceContext struct {
    CategoryModel category.Model  // 接口类型
}

// ❌ 错误：直接依赖实现
type ServiceContext struct {
    CategoryModel *gorm.CategoryDao  // 具体类型
}
```

#### 1.3 模块化设计

- **高内聚低耦合** - 模块内部紧密相关，模块间松散耦合
- **独立部署** - 每个服务（api/job/consumer）可独立部署和扩展
- **清晰边界** - 模块间通过明确的接口和协议通信

### 2. 数据层原则

#### 2.1 双ORM模式

**设计理念**：同时支持GORM和SQLx，提供灵活性和性能平衡

**选择策略**：
- **GORM优先** - 复杂查询、关联加载、事务管理
- **SQLx降级** - 简单查询、性能敏感场景、批量操作
- **工厂模式** - 自动选择可用的ORM实现
- **接口统一** - 两种ORM实现相同接口，业务层无感知

**实现要求**：
```go
// 1. 统一接口
type Model interface {
    Insert(ctx context.Context, data *T) (*T, error)
    FindOne(ctx context.Context, id int64) (*T, error)
    // ... 其他方法
    WithTx(tx interface{}) Model
    Trans(ctx context.Context, fn func(ctx context.Context, model Model) error) error
}

// 2. 工厂函数
func NewModel(sqlConn *sql.DB, gormDB *gorm.DB) Model {
    if gormDB != nil && gormFactory != nil {
        return gormFactory(gormDB)  // 优先GORM
    }
    if sqlConn != nil && sqlxFactory != nil {
        return sqlxFactory(sqlConn)  // 降级SQLx
    }
    panic("no database connection available")
}
```

#### 2.2 Model层组织

**表级目录模式（Table-Based Pattern）**：

```
model/resource_catalog/category/
├── interface.go    # Model接口定义
├── types.go        # Category数据结构
├── vars.go         # 常量和错误定义
├── factory.go      # ORM工厂函数
├── gorm_dao.go     # GORM实现
└── sqlx_model.go   # SQLx实现
```

**优势**：
- ✅ 每个表是独立单元，职责清晰
- ✅ 新增表不影响现有表
- ✅ 易于定位和维护

#### 2.3 数据访问约束

**强制规则**：
- ❌ **禁止** Logic层直接访问database
- ❌ **禁止** 跨层查询（必须通过Model层）
- ❌ **禁止** 在Handler中编写业务逻辑
- ✅ **必须** 使用Model提供的Trans方法处理事务
- ✅ **必须** 统一错误类型和处理方式

### 3. API设计原则

#### 3.1 RESTful API规范

- **资源导向** - URL表示资源，HTTP动词表示操作
```
GET    /api/v1/catalog/categories      # 列表
GET    /api/v1/catalog/categories/:id  # 详情
POST   /api/v1/catalog/categories      # 创建
PUT    /api/v1/catalog/categories/:id  # 更新
DELETE /api/v1/catalog/categories/:id  # 删除
```

- **统一响应格式**：
```go
{
    "code": 0,           // 业务状态码
    "msg": "success",    // 提示信息
    "data": {},          // 响应数据
    "requestId": "..."   // 请求追踪ID
}
```

- **版本管理** - 使用 `/api/v1/` 路径前缀

#### 3.2 Go-Zero框架规范

**必须遵循**：
- 使用 `goctl api` 工具生成代码
- 定义 `.api` 文件描述接口
- Handler只负责参数接收，Logic实现业务逻辑
- 使用 `types` 包定义请求响应结构

**中间件链（固定顺序）**：
```
Recovery → RequestID → Trace → CORS → Logger
```

### 4. 可观测性原则

#### 4.1 日志规范

**结构化日志**：
```go
logx.WithContext(ctx).Infow("operation completed",
    logx.Field("userId", userId),
    logx.Field("operation", "createCategory"),
    logx.Field("duration", duration),
)
```

**分级使用**：
- `Info` - 正常业务流程
- `Warn` - 警告信息（不影响功能）
- `Error` - 错误信息（需要关注）
- `Severe` - 严重错误（需要立即处理）

**上下文传递**：
- 必须携带 `RequestID`
- 必须携带 `TraceID`（分布式追踪）
- 关键业务参数要记录

#### 4.2 链路追踪

**OpenTelemetry集成**：
- 所有HTTP请求自动创建trace
- 关键操作（数据库查询、RPC调用）创建span
- 记录操作耗时和结果

**使用示例**：
```go
ctx, span := trace.StartInternal(ctx, "CreateCategory")
defer span.End()

// 业务逻辑
category, err := l.svcCtx.CategoryModel.Insert(ctx, data)
if err != nil {
    span.RecordError(err)
    return nil, err
}
```

#### 4.3 指标收集

**关键指标**：
- API调用次数、成功率、响应时间
- 数据库连接数、查询时间
- 缓存命中率
- 系统资源（CPU、内存、网络）

---

## 💻 开发规范

### 1. 代码风格

#### 1.1 命名规范

**文件命名**（全小写，下划线分隔）：
```
createcategoryhandler.go   # Handler
createcategorylogic.go     # Logic
category_dao.go            # GORM DAO
category_model.go          # SQLx Model
```

**包命名**（全小写，简短有意义）：
```
category      # 表级包
middleware    # 功能包
validator     # 工具包
```

**类型命名**（大驼峰PascalCase）：
```go
type CategoryModel interface {}  # 接口
type Category struct {}           # 结构体
type UserInfo struct {}          # 数据传输对象
```

**函数命名**：
```go
// 公开函数 - 大驼峰
func NewCategoryModel() Model {}
func CreateCategory() error {}

// 私有函数 - 小驼峰
func validateInput() error {}
func buildQuery() string {}
```

#### 1.2 代码组织

**导入顺序**：
```go
import (
    // 1. 标准库
    "context"
    "fmt"
    
    // 2. 第三方库（字母序）
    "github.com/zeromicro/go-zero/core/logx"
    "gorm.io/gorm"
    
    // 3. 项目内包（依赖关系）
    "idrm/model/resource_catalog/category"
    "idrm/pkg/middleware"
)
```

**文件内结构**：
```go
package xxx

// 1. 类型定义
type Model interface {}

// 2. 构造函数
func NewModel() Model {}

// 3. 公开方法（接口方法优先）
func (m *Model) Insert() {}

// 4. 私有方法
func (m *Model) validate() {}

// 5. init函数
func init() {}
```

#### 1.3 注释规范

**必须注释**：
- 所有 `public` 的 type、function、method
- 复杂的业务逻辑
- 非显而易见的实现细节

**注释语言**：中文

**注释格式**：
```go
// CategoryModel 定义类别仓储接口
// 支持GORM和SQLx双ORM实现，通过工厂模式自动选择
type CategoryModel interface {
    // Insert 插入新的类别记录
    // 参数data中的Id会被自动生成
    // 返回插入后的完整记录（包含Id和时间戳）
    Insert(ctx context.Context, data *Category) (*Category, error)
}
```

### 2. 错误处理

#### 2.1 错误定义

**在 vars.go 中定义**：
```go
var (
    ErrNotFound          = errors.New("category not found")
    ErrCodeAlreadyExists = errors.New("category code already exists")
    ErrInvalidStatus     = errors.New("invalid status")
)
```

**错误封装**：
```go
// 使用 %w 保留错误链
return fmt.Errorf("failed to insert category: %w", err)
```

#### 2.2 错误返回

**显式返回**：
```go
// ✅ 正确
func CreateCategory() (*Category, error) {
    if err := validate(); err != nil {
        return nil, err
    }
    return category, nil
}

// ❌ 错误：忽略error
func CreateCategory() *Category {
    validate() // 忽略了error
    return category
}
```

**上下文信息**：
```go
// ✅ 包含足够上下文
return nil, fmt.Errorf("failed to create category %s: %w", name, err)

// ❌ 信息不足
return nil, err
```

### 3. 测试要求

#### 3.1 单元测试

**覆盖率要求**：
- 核心业务逻辑 >80%
- 工具函数 >90%

**命名规范**：
```go
func TestCreateCategory(t *testing.T) {}
func TestCategoryModel_Insert(t *testing.T) {}
```

**表驱动测试**：
```go
func TestValidate(t *testing.T) {
    tests := []struct {
        name    string
        input   *Category
        wantErr bool
    }{
        {"valid", &Category{Name: "test"}, false},
        {"empty name", &Category{}, true},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            err := validate(tt.input)
            if (err != nil) != tt.wantErr {
                t.Errorf("unexpected error: %v", err)
            }
        })
    }
}
```

#### 3.2 集成测试

- 使用测试数据库（独立于开发/生产）
- 测试完整的请求-响应流程
- 外部依赖使用Mock

### 4. 文档要求

#### 4.1 代码文档

**必须文档**：
- 每个主要模块的 `README.md`
- API定义文件（`.api`）
- 复杂算法的说明文档

**README.md 模板**：
```markdown
# 模块名称

## 功能说明
简要描述模块功能

## 使用方法
提供使用示例

## 注意事项
列出关键约束和注意点
```

#### 4.2 变更文档

- **CHANGELOG.md** - 记录重要变更
- **架构文档** - `docs/architecture/`
- **API文档** - `api/doc/`

---

## 🤖 AI 协作规范

### 1. AI的角色定位

**AI是辅助工具，不是替代品**：

- ✅ **AI负责**：
  - 代码生成（基于明确的规范）
  - 重复性工作（样板代码）
  - 模式应用（按既定模板）
  - 文档生成

- ✅ **开发者负责**：
  - 架构决策
  - 业务理解
  - Code Review
  - 质量把控

- ✅ **共同负责**：
  - 质量保证
  - 测试验证
  - 持续改进

### 2. 工作流约束

#### 2.1 四阶段流程（严格遵守）

```
┌─────────────┐    ┌─────────────┐    ┌─────────────┐    ┌─────────────┐
│   Specify   │ => │    Plan     │ => │    Tasks    │ => │  Implement  │
│  功能规范    │    │  技术方案    │    │  任务拆分    │    │   代码实现   │
└─────────────┘    └─────────────┘    └─────────────┘    └─────────────┘
      ↓                  ↓                  ↓                  ↓
   需求澄清          架构设计          任务分解          增量交付
```

**不允许跳过任何阶段！**

#### 2.2 各阶段要求

**Specify 阶段**：
```markdown
必须包含：
- 功能描述（做什么）
- 用户故事（为谁、为什么）
- 技术约束（边界条件）
- 不确定项（需澄清的问题）

禁止包含：
- 具体技术实现
- 代码片段
```

**Plan 阶段**：
```markdown
必须包含：
- 技术方案（架构设计）
- 文件清单（新增/修改/删除）
- 数据模型（表结构、字段）
- 依赖关系（前置条件）

必须遵循：
- 现有架构原则
- 已定技术栈
- 编码规范
```

**Tasks 阶段**：
```markdown
任务要求：
- 每个任务 <50行代码
- 任务可独立测试
- 明确验收标准
- 按依赖排序

任务格式：
Task 1: 描述
- 具体步骤
- 验收条件
```

**Implement 阶段**：
```markdown
实施规则：
- 一次执行一个任务
- 生成完整可运行代码
- 包含必要的测试
- 符合所有规范
```

#### 2.3 检查点

每个阶段结束必须：
1. ✅ Review输出内容
2. ✅ 确认符合规范
3. ✅ 明确批准后进入下一阶段

### 3. 代码生成规则

#### 3.1 必须遵守的模式

**Model层模板**：
```go
// 目录结构
model/{module}/{table}/
  ├── interface.go    // Model接口
  ├── types.go        // 数据结构
  ├── vars.go         // 常量错误
  ├── factory.go      // ORM工厂
  ├── gorm_dao.go     // GORM实现
  └── sqlx_model.go   // SQLx实现

// 接口定义
type Model interface {
    // CRUD方法
    Insert(ctx context.Context, data *T) (*T, error)
    FindOne(ctx context.Context, id int64) (*T, error)
    Update(ctx context.Context, data *T) error
    Delete(ctx context.Context, id int64) error
    
    // 事务方法
    WithTx(tx interface{}) Model
    Trans(ctx context.Context, fn func(ctx context.Context, model Model) error) error
}
```

**API层模板**：
```go
// Handler
type {Action}{Resource}Handler struct {
    svcCtx *svc.ServiceContext
}

func New{Action}{Resource}Handler(svcCtx *svc.ServiceContext) *{Action}{Resource}Handler {
    return &{Action}{Resource}Handler{svcCtx: svcCtx}
}

func (h *{Action}{Resource}Handler) {Action}{Resource}(w http.ResponseWriter, r *http.Request) {
    var req types.{Action}{Resource}Req
    if err := httpx.Parse(r, &req); err != nil {
        response.Error(w, err)
        return
    }
    
    l := logic.New{Action}{Resource}Logic(r.Context(), h.svcCtx)
    resp, err := l.{Action}{Resource}(&req)
    if err != nil {
        response.Error(w, err)
    } else {
        response.Success(w, resp)
    }
}

// Logic
type {Action}{Resource}Logic struct {
    logx.Logger
    ctx    context.Context
    svcCtx *svc.ServiceContext
}

func New{Action}{Resource}Logic(ctx context.Context, svcCtx *svc.ServiceContext) *{Action}{Resource}Logic {
    return &{Action}{Resource}Logic{
        Logger: logx.WithContext(ctx),
        ctx:    ctx,
        svcCtx: svcCtx,
    }
}

func (l *{Action}{Resource}Logic) {Action}{Resource}(req *types.{Action}{Resource}Req) (resp *types.{Action}{Resource}Resp, err error) {
    // 实现业务逻辑
    return
}
```

#### 3.2 强制约束

**❌ 严禁操作**：
- 跳过接口，直接使用具体实现
- 在Logic层直接访问数据库
- 硬编码配置信息
- 忽略错误处理
- 生成没有注释的代码
- 修改已有的架构模式
- 改变既定的目录结构

**✅ 必须遵守**：
- 所有生成代码都有完整注释
- 遵循命名规范
- 遵循文件组织规则
- 包含必要的错误处理
- 符合代码风格指南

### 4. 质量标准

#### 4.1 代码质量检查清单

- ✅ 编译通过 (`go build ./...`)
- ✅ 单元测试通过 (`go test ./...`)
- ✅ Lint无错误 (`golangci-lint run`)
- ✅ 符合代码规范
- ✅ 所有公开接口都有注释
- ✅ 错误处理完整

#### 4.2 提交前验证

```bash
# 1. 编译检查
go build ./...

# 2. 运行测试
go test ./... -v

# 3. 代码检查
golangci-lint run

# 4. 格式化
gofmt -s -w .

# 5. 导入整理
goimports -w .
```

### 5. Review流程

#### 5.1 AI生成代码的Review重点

1. **架构一致性** - 是否符合分层架构？
2. **模式正确性** - 是否使用了正确的模式？
3. **安全性** - 是否有SQL注入、XSS等风险？
4. **性能** - 是否有明显的性能问题？
5. **可维护性** - 代码是否清晰易懂？

#### 5.2 发现问题的处理

- **小问题** - 人工修正，记录反馈
- **大问题** - 要求AI重新生成
- **模式问题** - 更新规范，避免再犯

---

## 🔄 版本管理

### 1. Git工作流

#### 1.1 分支策略

```
main (生产)
  ↑
develop (开发)
  ↑
feature/* (功能)
bugfix/*  (修复)
refactor/* (重构)
```

**规则**：
- `main` - 受保护，只能通过PR合并
- `develop` - 开发主分支
- `feature/*` - 新功能开发
- `bugfix/*` - Bug修复
- `refactor/*` - 代码重构

#### 1.2 Commit规范

**格式** ([Conventional Commits](https://www.conventionalcommits.org/)):
```
<type>(<scope>): <subject>

<body>

<footer>
```

**Type类型**：
- `feat` - 新功能
- `fix` - Bug修复
- `refactor` - 重构（不改变功能）
- `docs` - 文档更新
- `style` - 格式调整（不影响代码）
- `test` - 测试相关
- `chore` - 构建/工具相关
- `perf` - 性能优化

**示例**：
```
feat(model): add directory model with dual ORM support

- Create model/resource_catalog/directory/
- Implement GORM and SQLx implementations
- Add factory pattern for ORM selection
- Update ServiceContext with DirectoryModel

Closes #123
```

#### 1.3 PR规范

**PR标题**：与Commit message一致

**PR描述**：
```markdown
## 变更说明
简要描述本次PR的目的和内容

## 变更类型
- [ ] 新功能
- [ ] Bug修复
- [ ] 重构
- [ ] 文档更新

## 影响范围
- [ ] API变更
- [ ] 数据库变更
- [ ] 配置变更

## 测试情况
- [ ] 单元测试通过
- [ ] 集成测试通过
- [ ] 手动测试通过

## 检查清单
- [ ] 代码已经过自我Review
- [ ] 添加了必要的注释
- [ ] 更新了相关文档
- [ ] 无breaking changes（或已在描述中说明）
```

### 2. 变更管理

#### 2.1 影响评估

**任何变更前必须评估**：
1. 是否影响现有API？ → 需要版本号升级
2. 是否影响数据结构？ → 需要数据迁移脚本
3. 是否需要配置调整？ → 需要更新配置文档
4. 是否影响性能？ → 需要性能测试

#### 2.2 向后兼容原则

- **API变更** - 必须保持向后兼容，或提供明确的迁移指南
- **数据库Schema** - 变更必须可回滚
- **配置项** - 新增配置必须有合理默认值
- **依赖升级** - 必须经过充分测试

#### 2.3 Breaking Changes处理

如果必须引入Breaking Changes：
1. 提前通知（至少一个版本周期）
2. 提供迁移指南
3. 保留兼容层（至少一个版本）
4. 在CHANGELOG中明确标注

---

## 📚 附录

### A. 技术栈清单

**后端框架**：
- Go 1.21+
- Go-Zero v1.9+

**ORM**：
- GORM v1.31+
- go-zero/core/stores/sqlx

**数据库**：
- MySQL 8.0+
- Redis 7.0+

**消息队列**：
- Kafka 3.0+

**可观测性**：
- OpenTelemetry
- Jaeger
- ELK Stack
  - Elasticsearch 8.0+
  - Logstash 8.0+
  - Kibana 8.0+
  - Filebeat 8.0+

**容器化**：
- Docker
- Docker Compose

### B. 常用命令

```bash
# 开发命令
make dev          # 本地开发
make build        # 构建
make test         # 测试
make lint         # 代码检查

# Docker命令
make docker-build # 构建镜像
make docker-up    # 启动服务
make docker-down  # 停止服务
make docker-logs  # 查看日志

# 代码生成
make api-gen      # 生成API代码
make model-gen    # 生成Model代码

# 代码质量
make fmt          # 格式化
make imports      # 整理导入
```

### C. 参考资源

**项目文档**：
- `/docs/architecture/` - 架构文档
- `/model/README.md` - Model层说明
- `/api/doc/STRUCTURE.md` - API结构说明
- `/pkg/middleware/README.md` - 中间件文档

**外部资源**：
- [Go-Zero文档](https://go-zero.dev/)
- [GORM文档](https://gorm.io/)
- [GitHub Spec-Kit](https://github.com/github/spec-kit)

---

## 📝 宪章维护

### 更新流程

1. 提出修改建议（通过Issue或PR）
2. 团队讨论和Review
3. 达成共识后更新
4. 更新版本号和日期
5. 通知所有成员

### 版本历史

- **v1.0** (2024-12-24) - 初始版本，确立基本原则和规范

---

**本宪章是IDRM项目的根本指导文件，所有开发活动都必须遵循本文档的原则和规范。**
