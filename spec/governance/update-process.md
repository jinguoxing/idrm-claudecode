# 规范更新流程

> **Version**: 2.0.0

## 提议变更

### 1. 创建Issue
- 标题：`[Spec] 更新xxx规范`
- 说明变更原因和影响

### 2. 讨论
- 团队Review
- 评估影响范围

### 3. 批准
- Tech Lead批准
- Breaking change需全员同意

## 实施变更

### 1. 创建PR
- 分支：`spec/update-xxx`
- 修改相关文档
- 更新CHANGELOG

### 2. Review
- 至少2人Review
- 自动检查通过

### 3. 合并
- 合并到main
- 打tag（如v2.1.0）

### 4. 通知
- 发送变更通知
- 更新文档wiki

### 5. 同步工具配置
```bash
# 重新生成工具配置
make update-tool-configs
```

## 定期Review

### 季度Review
- 检查规范适用性
- 收集改进建议
- 更新最佳实践

### 年度Review
- 全面评估
- 重大版本规划
