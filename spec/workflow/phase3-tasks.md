# Phase 3: Tasks (任务拆分)

> **Level**: MUST  
> **Version**: 2.0.0

## 目标

将设计拆分为可执行的小任务（<50行代码）。

## 输入
- Phase 2的design.md

## 输出模板

创建 `specs/features/{name}/tasks.md`

## 任务拆分原则

1. **小而完整** - 每个task <50行
2. **可独立验证** - 有明确验收标准
3. **依赖清晰** - 按依赖关系排序
4. **职责单一** - 每个task做一件事

## 状态标识
- ⏸️ Not Started
- 🚧 In Progress
- ✅ Completed
- ❌ Blocked

## 质量门禁

- [ ] 每个task <50行
- [ ] 依赖关系清晰
- [ ] 验收标准明确

参考：`../quality/quality-gates.md#gate-3`

## 下一步
→ Phase 4: Implement
