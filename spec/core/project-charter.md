# IDRM Project Charter

> **Version**: 2.0.0  
> **Last Updated**: 2025-12-26  
> **Status**: Draft  
> **Owner**: Tech Lead

---

## 概述

IDRM (Intelligent Data Resource Management) 是一个智能数据资源管理平台，采用Go-Zero微服务架构，支持数据资源的全生命周期管理。

---

## 项目使命

为企业提供统一、高效、智能的数据资源管理平台，实现数据资产的标准化、可视化和智能化管理。

---

## 核心价值观

### 1. 规范先行 (Spec-Driven)
- 所有功能开发前先定义规范
- AI辅助开发严格遵循规范
- 规范持续演进和优化

### 2. 质量第一 (Quality First)
- 代码质量高于开发速度
- 架构规范严格执行
- 测试覆盖充分完整

### 3. 团队协作 (Collaboration)
- 清晰的工作流程
- 充分的文档和沟通
- 知识共享和传承

### 4. 持续改进 (Continuous Improvement)
- 定期Review和优化
- 快速响应需求变化
- 技术栈适时升级

---

## 技术栈

参考：`tech-stack.md`

**核心技术**：
- **语言**: Go 1.21+
- **框架**: Go-Zero v1.9+
- **数据库**: MySQL 8.0
- **缓存**: Redis 7.0
- **消息队列**: Kafka 3.0

---

## 开发原则

### 架构原则

参考：`../architecture/`

1. **分层架构** - Handler → Logic → Model 严格分离
2. **接口驱动** - 面向接口编程，便于测试和替换
3. **双ORM支持** - GORM + SQLx，灵活选择
4. **微服务化** - 服务边界清晰，独立部署

### 编码原则

参考：`../coding-standards/`

1. **可读性优先** - 清晰的命名和完整的注释
2. **简单优于复杂** - KISS原则
3. **DRY原则** - 避免重复代码
4. **错误处理完整** - 所有error必须处理

---

## AI协作规范

参考：`workflow.md`

### 标准工作流

采用**5阶段工作流**：

```
Phase 0: Context    - 准备上下文，理解规范
Phase 1: Specify    - 定义需求和验收标准
Phase 2: Design     - 技术方案和架构设计
Phase 3: Tasks      - 任务拆分和规划
Phase 4: Implement  - 实施、测试和验收
```

### AI工具使用

- **Kiro.dev** - 大型功能规划和团队协作
- **Cursor** - 日常快速开发和重构
- **Claude CLI** - 批量processing和CI/CD

### AI角色定位

- AI是**辅助工具**，不是决策者
- 开发者保留**最终责任**
- 所有AI生成代码必须**Review**

---

## 质量标准

参考：`../quality/`

### 代码质量

- ✅ 编译通过 (`go build ./...`)
- ✅ 测试通过 (`go test ./...`)
- ✅ 覆盖率 >80% (核心业务逻辑)
- ✅ Lint无错误 (`golangci-lint run`)

### 架构质量

- ✅ 符合分层架构
- ✅ 接口设计合理
- ✅ 依赖关系清晰
- ✅ 可测试性强

### 文档质量

- ✅ API文档完整
- ✅ 公开接口有注释
- ✅ README及时更新
- ✅ 规范文档准确

---

## 版本管理

### Git工作流

- **main** - 生产分支，稳定代码
- **develop** - 开发分支，日常开发
- **feature/** - 功能分支
- **hotfix/** - 紧急修复分支

### Commit规范

遵循Conventional Commits：

```
feat: 添加directory管理功能
fix: 修复category查询bug
docs: 更新API文档
refactor: 重构model层
test: 添加unit tests
```

### PR流程

1. Self Review
2. 自动化检查 (CI/CD)
3. Peer Review (至少1人)
4. Tech Lead批准
5. 合并到develop

---

## 参考文档

### 核心规范
- `tech-stack.md` - 技术栈详细说明
- `values.md` - 价值观详解
- `workflow.md` - 完整工作流定义

### 其他规范
- `../architecture/` - 架构规范
- `../coding-standards/` - 编码规范
- `../workflow/` - 工作流指南
- `../quality/` - 质量保证
- `../governance/` - 治理机制

---

## 变更历史

| Version | Date | Changes | Author |
|---------|------|---------|--------|
| 2.0.0 | 2025-12-26 | 优化为v2规范体系 | Tech Lead |
| 1.0.0 | 2025-12-24 | 初始版本 | Tech Lead |

---

**本文档是IDRM项目的根本法，所有开发活动必须遵循！**
