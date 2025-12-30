# 规范版本管理

> **Version**: 2.0.0

## 语义化版本号

遵循 SemVer：`MAJOR.MINOR.PATCH`

- **MAJOR**: 破坏性变更（如架构调整）
- **MINOR**: 新增功能（向后兼容）
- **PATCH**: Bug修复

## 版本追踪

每个spec文档头部包含：

```markdown
> **Version**: 2.0.0  
> **Last Updated**: 2025-12-26  
> **Status**: Draft | Review | Approved
```

## 变更日志

`private_doc/spec/CHANGELOG.md`：

```markdown
# Spec变更日志

## [2.0.0] - 2025-12-26

### Breaking Changes
- 统一为5阶段工作流
- 重构目录结构

### Added
- workflow/ 工作流指南
- quality/ 质量保证
- governance/ 治理机制

## [1.0.0] - 2025-12-24

### Added
- 初始spec框架
```

## 更新流程

1. 创建Issue提议变更
2. 团队讨论Review
3. 创建PR实施
4. 合并并打tag
5. 通知团队
