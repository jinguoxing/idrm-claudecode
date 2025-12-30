# Cursor IDE 使用指南

> **文档版本**: v1.0  
> **最后更新**: 2025-12-24  
> **适用于**: IDRM项目

---

## 📖 概述

本文档介绍如何在Cursor IDE中使用IDRM项目的spec规范进行AI辅助开发。

---

## ⚠️ 重要说明

### 关于规范文档路径

本指南中所有涉及的规范文档路径（如 `private_doc/spec/`）都是基于**IDRM项目的实际目录结构**。

**如果您在其他项目中使用本指南**，需要根据您项目的实际情况调整：

- **规范文档目录**：可能是 `docs/specs/`、`.github/specs/` 或其他路径
- **文档文件名**：如 `constitution.md`、`architecture.md` 等需要对应您项目的实际文件名
- **@-mentions路径**：在Cursor中引用文档时使用您的实际路径

**示例调整**：

```
# IDRM项目示例
@private_doc/spec/constitution.md 帮我实现功能

# 您的项目可能是
@docs/project-guidelines.md 帮我实现功能
# 或
@.github/specs/development-standards.md 帮我实现功能
```

**建议**：
1. 确认您项目中规范文档的实际位置
2. 在使用@-mentions时使用正确的相对路径
3. 在 `.cursorrules` 文件中更新规范文档的路径引用

---

## 🚀 快速开始

### 1. 前置条件

- ✅ 已安装Cursor IDE
- ✅ 已打开IDRM项目
- ✅ `.cursorrules`文件已在项目根目录

### 2. 验证配置

在Cursor的AI聊天窗口询问：

```
你了解IDRM项目的架构规范吗？
```

如果AI能回答出分层架构、双ORM等信息，说明`.cursorrules`已生效。

---

## 💡 三种使用方式

### 方式1：自动加载规则

Cursor会自动读取`.cursorrules`文件，AI已经了解：
- 项目的技术栈
- 分层架构规范
- 4阶段工作流
- 编码标准

**使用**：
```
帮我实现directory表的CRUD功能
```

AI会自动遵循规范。

### 方式2：@-mentions引用文档

使用`@`符号引用特定规范文档：

```
@private_doc/spec/constitution.md 
按照这个规范帮我实现新功能
```

```
@private_doc/spec/architecture/layered-architecture.md
检查这段代码是否符合分层架构
```

### 方式3：组合使用（推荐⭐）

结合自动规则和@-mentions：

```
@private_doc/spec/architecture/dual-orm-pattern.md
@private_doc/spec/coding-standards/go-style-guide.md

帮我实现directory模型，要求：
1. 遵循table-based目录结构
2. 实现GORM和SQLx双ORM
3. 添加完整的中文注释
```

---

## 📋 4阶段工作流

### Phase 1: Specify（需求规范）

**提示词模板**：
```
@private_doc/spec/constitution.md

我需要添加{功能描述}。

Phase 1: 请先帮我创建功能规范(Specification)：
- 用户故事
- 功能需求
- 技术约束
- 不确定项
```

**AI输出**：详细的功能规范文档

**你的操作**：Review并确认

### Phase 2: Plan（技术方案）

**提示词模板**：
```
规范已确认。

Phase 2: 请给出技术实施方案：
- 架构设计
- 文件清单（新增/修改/删除）
- 数据模型
- 依赖关系

@private_doc/spec/architecture/layered-architecture.md
@private_doc/spec/architecture/dual-orm-pattern.md
```

**AI输出**：详细的技术方案

**你的操作**：Review并确认

### Phase 3: Tasks（任务拆分）

**提示词模板**：
```
方案已确认。

Phase 3: 请将方案拆分为可执行任务：
- 每个任务<50行代码
- 明确依赖关系
- 标注验收标准
```

**AI输出**：任务清单

**你的操作**：Review并确认

### Phase 4: Implement（实施）

**提示词模板**：
```
请实施Task 1: {任务描述}

要求：
- 遵循命名规范
- 添加中文注释
- 完整的错误处理

@private_doc/spec/coding-standards/go-style-guide.md
@private_doc/spec/coding-standards/naming-conventions.md
@private_doc/spec/coding-standards/error-handling.md
```

**AI输出**：完整的代码实现

**你的操作**：Review、测试、提交

---

## 🎯 常用场景

### 场景1：新增功能

**完整流程**：
```
# Step 1: Specify
@private_doc/spec/constitution.md
我要添加资源目录管理功能，支持树形结构。
请按4阶段工作流，先给我Specification。

# Step 2: Plan
[确认spec后]
给出技术方案
@private_doc/spec/architecture/layered-architecture.md

# Step 3: Tasks
[确认plan后]
拆分为可执行任务

# Step 4: Implement
[确认tasks后]
实施Task 1
@private_doc/spec/coding-standards/
```

### 场景2：代码Review

**提示词**：
```
@private_doc/spec/coding-standards/code-review-checklist.md

请Review以下代码：
{粘贴代码或文件路径}

检查：
- 是否符合分层架构
- 错误处理是否正确
- 命名是否规范
- 是否需要改进
```

### 场景3：重构代码

**提示词**：
```
@private_doc/spec/architecture/layered-architecture.md
@private_doc/spec/architecture/dual-orm-pattern.md

以下代码违反了架构规范，请重构：
{粘贴代码}

要求：
- 分离Handler/Logic/Model
- 使用接口抽象
- 保持功能不变
```

### 场景4：生成测试

**提示词**：
```
@private_doc/spec/coding-standards/testing-standards.md

为以下代码生成单元测试：
{粘贴代码或文件路径}

要求：
- 使用表驱动测试
- 覆盖边界情况
- 使用Mock（如需要）
```

### 场景5：学习规范

**提示词**：
```
@private_doc/spec/architecture/layered-architecture.md

请总结IDRM项目的分层架构规范，包括：
1. 各层职责
2. 依赖规则
3. 常见错误
```

---

## 🔧 高级技巧

### 技巧1：使用Composer（多文件编辑）

**快捷键**：`Cmd/Ctrl + I`

**使用场景**：一次性创建完整的model

**示例**：
```
@private_doc/spec/architecture/dual-orm-pattern.md

请创建完整的directory model：
1. model/resource_catalog/directory/interface.go
2. model/resource_catalog/directory/types.go
3. model/resource_catalog/directory/vars.go
4. model/resource_catalog/directory/factory.go
5. model/resource_catalog/directory/gorm_dao.go
6. model/resource_catalog/directory/sqlx_model.go

参考category model的实现。
```

Composer会同时生成所有文件。

### 技巧2：引用现有代码

```
@model/resource_catalog/category/gorm_dao.go
@private_doc/spec/architecture/dual-orm-pattern.md

参考category的实现，为directory创建相同的结构
```

### 技巧3：批量Review

```
@private_doc/spec/coding-standards/code-review-checklist.md

Review api/internal/logic/目录下的所有文件
```

### 技巧4：持续对话

Cursor会记住对话历史：

```
# 第1轮
实现Task 1

# 第2轮
继续实现Task 2

# 第3轮
现在基于前面的Model，创建Logic层
```

---

## ⚙️ Cursor配置建议

### AI Model选择

**Settings → Features → AI**

- **Primary Model**: Claude 3.5 Sonnet（推荐）
  - 更好理解复杂规范
  - 代码质量更高
  
- **Long Context**: 启用
  - 支持读取大量spec文档

### 启用的功能

- ✅ **Cursor Tab** - 智能代码补全
- ✅ **Copilot++** - 多行补全
- ✅ **Chat** - AI对话
- ✅ **Composer** - 多文件编辑

### 快捷键设置

- `Cmd/Ctrl + K` - 快速AI编辑
- `Cmd/Ctrl + I` - Composer
- `Cmd/Ctrl + L` - AI聊天
- `Cmd/Ctrl + Shift + L` - 新建对话

---

## 📝 最佳实践

### DO ✅

1. **明确引用规范**
   ```
   @private_doc/spec/constitution.md
   帮我实现xxx
   ```

2. **遵循4阶段流程**
   - 不要跳过Specify和Plan阶段
   - 确保每个阶段都经过Review

3. **提供上下文**
   ```
   @existing/file.go
   @spec/document.md
   基于这些实现新功能
   ```

4. **分步实施**
   - 一次一个任务
   - 及时测试验证

5. **及时Review**
   - 每个任务完成后Review
   - 发现问题立即修正

### DON'T ❌

1. ❌ **不要只依赖AI**
   - AI是辅助工具，最终决策权在你

2. ❌ **不要跳过规范**
   - 没有引用spec就让AI生成代码

3. ❌ **不要直接采用全部代码**
   - 必须Review后再使用

4. ❌ **不要忽略错误**
   - AI生成的代码可能有误

5. ❌ **不要一次生成大量代码**
   - 遵循增量开发

---

## 🔍 常见问题

### Q: AI没有遵循规范怎么办？

**A**: 
1. 确认`.cursorrules`文件存在
2. 重启Cursor
3. 明确引用规范文档：`@private_doc/spec/xxx.md`

### Q: @-mentions不生效？

**A**: 
1. 确认文件路径正确
2. 使用相对于项目根的路径
3. 文件不要太大（<100KB）

### Q: AI生成的代码质量不高？

**A**: 
1. 提供更详细的要求
2. 引用多个相关规范文档
3. 给出具体示例代码
4. 使用Claude 3.5 Sonnet模型

### Q: 如何重新开始对话？

**A**: 
- 快捷键：`Cmd/Ctrl + Shift + L`
- 或点击聊天窗口的"New Chat"

---

## 📚 参考资源

### IDRM规范文档

- `private_doc/spec/constitution.md` - 项目宪章
- `private_doc/spec/architecture/` - 架构规范
- `private_doc/spec/coding-standards/` - 编码规范

### Cursor官方文档

- [Cursor文档](https://cursor.sh/docs)
- [.cursorrules示例](https://github.com/PatrickJS/awesome-cursorrules)

### 工具配置

- `.cursorrules` - 项目AI规则
- 项目根目录

---

## ✨ 总结

### 核心要点

1. **自动+手动结合**：`.cursorrules` + `@-mentions`
2. **严格遵循4阶段**：Specify → Plan → Tasks → Implement
3. **充分引用规范**：让AI了解完整上下文
4. **增量开发**：一次一个任务
5. **持续Review**：不盲目信任AI

### 预期效果

- ✅ 代码自动符合规范
- ✅ 开发效率提升30%+
- ✅ 代码质量一致性
- ✅ 减少Code Review时间

---

**使用Cursor + IDRM Spec = 高效规范的AI辅助开发！** 🚀
