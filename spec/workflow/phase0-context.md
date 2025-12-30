# Phase 0: Context (上下文准备)

> **Level**: SHOULD  
> **Version**: 2.0.0

## 目标

准备开发上下文，理解项目规范。

## 活动清单

### 1. 阅读核心规范
- [ ] `../core/project-charter.md`
- [ ] `../core/tech-stack.md`
- [ ] `../core/workflow.md`

### 2. 了解相关规范
- [ ] `../architecture/layered-architecture.md`
- [ ] `../architecture/dual-orm-pattern.md`
- [ ] `../coding-standards/go-style-guide.md`

### 3. 准备环境
- [ ] 代码已clone
- [ ] 依赖已安装
- [ ] 数据库已启动
- [ ] AI工具已配置

## AI工具使用

**Cursor**:
```
@private_doc/spec/core/project-charter.md
请总结IDRM项目的核心规范
```

**Claude CLI**:
```bash
claude --files "private_doc/spec/core/*.md" \
  "总结项目规范要点"
```

**Kiro.dev**: Steering自动加载规范

## 检查清单
- [ ] 理解分层架构
- [ ] 理解双ORM模式
- [ ] 熟悉编码规范
- [ ] 开发环境就绪

## 下一步
→ Phase 1: Specify
