# IDRM Spec规范文档 v2.0

> **Version**: 2.0.0  
> **Last Updated**: 2025-12-26

## 概述

本目录包含IDRM项目的所有规范文档，采用spec-driven development理念。

## 目录结构

```
spec/
├── core/              # 核心规范（MUST READ）
├── architecture/      # 架构规范
├── coding-standards/  # 编码规范
├── workflow/          # 工作流指南
├── quality/           # 质量保证
├── tools/             # 工具配置
└── governance/        # 治理机制
```

## 快速开始

### 新人入职
1. 阅读 `core/project-charter.md`
2. 了解 `core/tech-stack.md`
3. 学习 `core/workflow.md`

### 开发新功能
1. 遵循 `workflow/` 的5阶段流程
2. 引用 `architecture/` 和 `coding-standards/`
3. 通过 `quality/` 的质量门禁

### 使用AI工具
- Cursor: 参考 `tools/cursor/`
- Claude CLI: 参考 `tools/claude-cli/`
- Kiro.dev: 参考 `tools/kiro/`

## 核心原则

1. **规范先行** - Spec-Driven Development
2. **严格执行** - 规范必须遵守  
3. **持续演进** - 定期Review和更新

## 版本历史

详见 [CHANGELOG.md](CHANGELOG.md)
