# Claude Code CLI 使用指南

> **文档版本**: v1.0  
> **最后更新**: 2025-12-24  
> **适用于**: IDRM项目

---

## 📖 概述

本文档介绍如何使用Claude Code CLI配合IDRM项目的spec规范进行命令行AI辅助开发。

---

## ⚠️ 重要说明

### 关于规范文档路径

本指南中所有涉及的规范文档路径（如 `private_doc/spec/`）都是基于**IDRM项目的实际目录结构**。

**如果您在其他项目中使用本指南**，需要根据您项目的实际情况调整：

- **规范文档目录**：可能是 `docs/specs/`、`.github/specs/` 或其他路径
- **文档文件名**：如 `constitution.md`、`architecture.md` 等需要对应您项目的实际文件名
- **目录结构**：根据您项目的组织方式调整

**示例调整**：

```bash
# IDRM项目示例
claude --files "private_doc/spec/constitution.md" "..."

# 您的项目可能是
claude --files "docs/project-guidelines.md" "..."
# 或
claude --files ".github/specs/development-standards.md" "..."
```

**建议**：
1. 确认您项目中规范文档的实际位置
2. 将本指南中的所有 `private_doc/spec/` 替换为您的实际路径
3. 在 `.clinerules` 文件中配置正确的路径

---

## 🚀 快速开始

### 1. 安装

```bash
# macOS (推荐使用Homebrew)
brew install anthropics/claude/claude

# 验证安装
claude --version
```

### 2. 配置API Key

```bash
# 设置环境变量
export ANTHROPIC_API_KEY="your-api-key-here"

# 永久保存（添加到shell配置）
echo 'export ANTHROPIC_API_KEY="your-api-key"' >> ~/.zshrc
source ~/.zshrc

# 验证
claude auth status
```

### 3. 进入项目目录

```bash
cd /Users/kingnet/code_workspace/go_workspace/src/idrm
```

### 4. 验证配置

```bash
# Claude会自动读取.clinerules
claude "读取项目规范，你了解这个项目吗？"
```

如果AI能回答出项目的架构规范，说明`.clinerules`已生效。

---

## 💡 三种使用方式

### 方式1：交互式对话

```bash
# 启动交互式会话
claude

# 然后输入提示词
> 读取private_doc/spec/constitution.md，理解项目规范
> 帮我实现directory表的CRUD功能
```

### 方式2：直接命令

```bash
# 一次性执行单个命令
claude "读取private_doc/spec/constitution.md，总结项目规范"

# 带文件上下文
claude --files "private_doc/spec/**/*.md" "实现新功能"
```

### 方式3：使用Slash Commands（推荐⭐）

```bash
# Phase 1: Specify
claude /speckit.specify

# Phase 2: Plan  
claude /speckit.plan

# Phase 3: Tasks
claude /speckit.tasks

# Phase 4: Implement
claude /speckit.implement
```

---

## 📋 4阶段工作流

### Phase 1: Specify（需求规范）

**方式A：交互式**

```bash
claude
```

然后输入：
```
读取 private_doc/spec/constitution.md

Phase 1: Specify

我需要添加directory表，用于管理资源目录。

请创建功能规范，包括：
- 用户故事
- 功能需求  
- 技术约束
- 不确定项
```

**方式B：直接命令**

```bash
claude --files "private_doc/spec/constitution.md" \
  "创建directory表的功能规范（Specification）"
```

**输出**：AI生成规范文档，可保存为 `specs/features/directory.md`

### Phase 2: Plan（技术方案）

```bash
claude --files "private_doc/spec/architecture/layered-architecture.md" \
       --files "private_doc/spec/architecture/dual-orm-pattern.md" \
  "基于前面的规范，创建技术实施方案"
```

**输出**：
- 架构设计
- 文件清单
- 数据模型
- 依赖关系

### Phase 3: Tasks（任务拆分）

```bash
claude "将方案拆分为可执行任务：
- 每个任务<50行代码
- 明确依赖关系
- 标注验收标准"
```

**输出**：任务清单（可保存为 `specs/tasks/directory.md`）

### Phase 4: Implement（实施）

```bash
# 实施Task 1
claude --files "private_doc/spec/coding-standards/go-style-guide.md" \
       --files "private_doc/spec/coding-standards/naming-conventions.md" \
  "实施Task 1: 创建directory基础结构"

# 实施Task 2
claude "继续实施Task 2: 实现factory"

# 实施Task 3
claude "继续实施Task 3: 实现GORM DAO"
```

---

## 🎯 常用场景

### 场景1：新增功能（完整流程）

```bash
# Step 1: 进入项目目录
cd /path/to/idrm

# Step 2: 启动交互式会话
claude

# Step 3: Phase 1 - Specify
> 读取 private_doc/spec/constitution.md
> 
> 我要添加资源目录管理功能，支持树形结构。
> Phase 1: 请创建Specification。

# Step 4: Review spec，确认后继续
> Phase 2: 基于spec创建技术方案
> 读取 private_doc/spec/architecture/layered-architecture.md
> 读取 private_doc/spec/architecture/dual-orm-pattern.md

# Step 5: Review plan，确认后继续
> Phase 3: 将方案拆分为可执行任务

# Step 6: Review tasks，确认后继续
> Phase 4: 实施Task 1
> 读取 private_doc/spec/coding-standards/go-style-guide.md

# Step 7: 逐个实施剩余任务
> 继续实施Task 2
> 继续实施Task 3
```

### 场景2：代码Review

```bash
claude --files "private_doc/spec/coding-standards/code-review-checklist.md" \
       --files "api/internal/logic/category/createcategorylogic.go" \
  "使用checklist检查这个文件的代码质量"
```

### 场景3：重构代码

```bash
# 分析现有代码
claude --files "api/internal/handler/category/*.go" \
  "分析这些Handler是否符合架构规范"

# 生成重构方案
claude --files "private_doc/spec/architecture/layered-architecture.md" \
  "给出重构方案"

# 执行重构（分任务）
claude "拆分重构任务"
claude "实施重构Task 1"
```

### 场景4：批量生成测试

```bash
claude --files "private_doc/spec/coding-standards/testing-standards.md" \
       --files "model/resource_catalog/category/*.go" \
  "为category model的所有方法生成单元测试"
```

### 场景5：学习规范

```bash
claude --files "private_doc/spec/**/*.md" \
  "总结IDRM项目的所有规范，以表格形式呈现"
```

---

## 🔧 高级用法

### 1. 精确控制文件上下文

```bash
# 只读取需要的文档
claude --files "private_doc/spec/architecture/dual-orm-pattern.md" \
  "实现新的model"

# 读取多个相关文档
claude --files "private_doc/spec/architecture/*.md" \
       --files "model/resource_catalog/category/*.go" \
  "参考category实现directory"
```

### 2. 使用glob模式

```bash
# 读取所有Go文件
claude --files "api/**/*.go" \
  "分析API层的代码结构"

# 读取所有spec文档
claude --files "private_doc/spec/**/*.md" \
  "总结项目规范"

# 读取specific文件  
claude --files "model/resource_catalog/*/gorm_dao.go" \
  "分析所有GORM实现的共同模式"
```

### 3. 批量处理

```bash
# 批量review所有Logic文件
for file in api/internal/logic/**/*.go; do
  claude --files "private_doc/spec/coding-standards/code-review-checklist.md" \
         --files "$file" \
    "review这个文件" >> review_report.txt
done
```

### 4. 管道组合

```bash
# 生成代码后直接保存
claude "实现directory的interface.go" > model/resource_catalog/directory/interface.go

# 组合多个命令
claude "生成任务清单" | tee tasks.md
```

### 5. 使用配置文件

创建 `.claude/config.yaml`:

```yaml
default_files:
  - "private_doc/spec/constitution.md"
  - "private_doc/spec/architecture/layered-architecture.md"

aliases:
  review: "--files 'private_doc/spec/coding-standards/code-review-checklist.md'"
  spec: "--files 'private_doc/spec/**/*.md'"
```

使用：
```bash
claude review --files "api/logic/*.go" "review代码"
claude spec "总结规范"
```

---

## 📝 创建辅助脚本

### 脚本1：spec.sh（规范管理）

创建 `scripts/spec.sh`:

```bash
#!/bin/bash

case $1 in
  "read")
    claude --files "private_doc/spec/**/*.md" "总结项目规范"
    ;;
  "check")
    claude --files "private_doc/spec/coding-standards/code-review-checklist.md" \
           --files "$2" \
      "检查代码是否符合规范"
    ;;
  "new")
    echo "启动新功能开发流程..."
    claude --files "private_doc/spec/constitution.md" \
      "Phase 1: Specify - 请引导我完成功能规范"
    ;;
  *)
    echo "Usage: $0 {read|check|new}"
    echo "  read       - 读取并总结规范"
    echo "  check FILE - 检查文件是否符合规范"
    echo "  new        - 开始新功能开发"
    ;;
esac
```

```bash
chmod +x scripts/spec.sh

# 使用
./scripts/spec.sh read
./scripts/spec.sh check "api/internal/logic/category/createcategorylogic.go"
./scripts/spec.sh new
```

### 脚本2：ai-dev.sh（开发流程）

创建 `scripts/ai-dev.sh`:

```bash
#!/bin/bash

SPEC_DIR="private_doc/spec"

# Phase 1: Specify
specify() {
  echo "=== Phase 1: Specify ==="
  claude --files "$SPEC_DIR/constitution.md" \
    "Phase 1: 创建功能规范。功能描述: $1"
}

# Phase 2: Plan
plan() {
  echo "=== Phase 2: Plan ==="
  claude --files "$SPEC_DIR/architecture/*.md" \
    "Phase 2: 基于spec创建技术方案"
}

# Phase 3: Tasks
tasks() {
  echo "=== Phase 3: Tasks ==="
  claude "Phase 3: 拆分为可执行任务（每个<50行）"
}

# Phase 4: Implement
implement() {
  echo "=== Phase 4: Implement ==="
  claude --files "$SPEC_DIR/coding-standards/*.md" \
    "Phase 4: 实施Task $1"
}

# 主流程
case $1 in
  "specify") specify "$2" ;;
  "plan") plan ;;
  "tasks") tasks ;;
  "implement") implement "$2" ;;
  "full")
    specify "$2"
    read -p "按Enter继续Plan阶段..."
    plan
    read -p "按Enter继续Tasks阶段..."
    tasks
    read -p "按Enter继续Implement阶段..."
    implement "1"
    ;;
  *)
    echo "Usage: $0 {specify|plan|tasks|implement|full} [args]"
    ;;
esac
```

使用：
```bash
chmod +x scripts/ai-dev.sh

# 完整流程
./scripts/ai-dev.sh full "添加directory表"

# 单独阶段
./scripts/ai-dev.sh specify "添加directory表"
./scripts/ai-dev.sh plan
./scripts/ai-dev.sh tasks
./scripts/ai-dev.sh implement 1
```

---

## ⚙️ Claude Code CLI配置

### 项目配置文件

`.clinerules` - 已存在于项目根目录

```.clinerules
# IDRM Project Rules for Claude Code
...（内容已包含）
```

### 可选配置

创建 `.claude/context.md`:

```markdown
# IDRM Project Context

## 当前状态
- 分层架构已建立
- 双ORM模式已实现
- category模型已完成

## 正在进行
- directory模型开发

## 下一步
- API接口完善
- 测试覆盖提升
```

---

## 📊 Claude Code vs Cursor对比

| 功能 | Claude Code CLI | Cursor |
|------|----------------|--------|
| **使用环境** | 终端/脚本 | IDE内 |
| **适合场景** | 结构化开发、批处理 | 快速编码 |
| **上下文控制** | 精确（--files指定） | 自动推断 |
| **批量处理** | ⭐⭐⭐ 强大 | ⭐ 一般 |
| **实时补全** | ❌ 不支持 | ⭐⭐⭐ 强大 |
| **CI/CD集成** | ⭐⭐⭐ 完美 | ❌ 不支持 |
| **学习曲线** | 中等 | 简单 |
| **Spec-Kit支持** | ⭐⭐⭐ 原生 | ⭐⭐ 需配置 |

**推荐组合**：
- 结构化开发 → Claude Code CLI
- 日常编码 → Cursor
- Code Review → 两者都可

---

## 🎯 最佳实践

### DO ✅

1. **明确指定文件**
   ```bash
   claude --files "private_doc/spec/constitution.md" "..."
   ```

2. **遵循4阶段流程**
   - 不跳过任何阶段
   - 每阶段都Review

3. **使用脚本自动化**
   - 创建常用流程脚本
   - 减少重复命令

4. **保存中间产物**
   ```bash
   claude "生成spec" > specs/features/directory.md
   claude "生成任务" > specs/tasks/directory.md
   ```

5. **利用管道**
   ```bash
   claude "生成代码" | tee output.go
   ```

### DON'T ❌

1. ❌ 不要忽略规范
   - 总是用--files引用spec文档

2. ❌ 不要一次生成太多
   - 遵循增量开发

3. ❌ 不要盲目信任输出
   - 必须Review

4. ❌ 不要忽略错误
   - 检查命令执行结果

5. ❌ 不要跳过测试
   - 及时验证生成的代码

---

## 🔍 常见问题

### Q: 如何查看可用命令？

```bash
claude --help
claude --list-commands
```

### Q: 如何中断长时间运行的命令？

`Ctrl + C`

### Q: 如何查看对话历史？

```bash
claude history
claude history --last 10
```

### Q: 如何清除对话历史？

```bash
claude history clear
```

### Q: 文件太大怎么办？

只读取需要的部分：
```bash
head -100 large_file.go | claude "分析这段代码"
```

### Q: 如何在CI/CD中使用？

```bash
# 在GitHub Actions中
- name: Code Review
  run: |
    claude --files "private_doc/spec/coding-standards/*.md" \
           --files "**/*.go" \
      "生成代码质量报告" > report.md
```

---

## 📚 参考资源

### IDRM规范

- `private_doc/spec/constitution.md`
- `private_doc/spec/architecture/`
- `private_doc/spec/coding-standards/`

### Claude Code CLI

- [官方文档](https://docs.anthropic.com/claude/docs/claude-code)
- [GitHub](https://github.com/anthropics/claude-code)

### 配置文件

- `.clinerules` - 项目规则
- `.claude/context.md` - 项目上下文（可选）
- `scripts/spec.sh` - 辅助脚本

---

## ✨ 总结

### 核心优势

1. **精确控制** - 通过--files精确指定上下文
2. **脚本化** - 可集成到开发工具链
3. **批量处理** - 一次处理多个文件/任务
4. **原生Spec-Kit** - 为spec-driven设计

### 推荐工作流

```bash
# 早上：复习规范
claude --files "private_doc/spec/constitution.md"

# 开发：4阶段流程
./scripts/ai-dev.sh full "新功能"

# Review：质量检查
./scripts/spec.sh check "path/to/file.go"

# 提交前：最终验证
claude --files "private_doc/spec/**/*.md" \
  "检查今天的代码变更"
```

---

**Claude Code CLI + IDRM Spec = 专业的spec-driven development！** 🚀
