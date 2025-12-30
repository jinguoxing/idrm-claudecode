# IDRM技术栈

> **Version**: 2.0.0  
> **Last Updated**: 2025-12-26  
> **Status**: Approved

---

## 编程语言

### Go 1.21+

**选择理由**：
- 高性能、高并发
- 简洁的语法
- 强大的标准库
- 优秀的工具链

**编码规范**: 参考 `../coding-standards/go-style-guide.md`

---

## Web框架

### Go-Zero v1.9+

**选择理由**:
- 微服务架构支持完善
- 内置服务发现、负载均衡
- 丰富的中间件
- 优秀的性能

**目录结构**:
```
api/          # API服务
├── doc/      # API定义
├── internal/ # 内部实现
└── etc/      # 配置文件
```

---

## ORM

### 双ORM支持

**GORM v2.0+**
- 功能丰富
- 适合复杂查询
- 支持关联和预加载

**SQLx**
- 轻量高效
- 更接近原生SQL
- 适合简单CRUD

**使用原则**: 参考 `../architecture/dual-orm-pattern.md`

---

## 数据库

### MySQL 8.0

**用途**: 主数据存储

**规范**:
- 表名: 复数小写 (categories)
- 字段名: 下划线分隔 (created_at)
- 索引: 合理使用，避免过度
- 字符集: utf8mb4

---

## 缓存

### Redis 7.0

**用途**:
- 数据缓存
- 会话存储
- 分布式锁
- 消息队列(简单场景)

**规范**:
- Key命名: `{project}:{module}:{type}:{id}`
- 过期时间: 必须设置
- 序列化: JSON

---

## 消息队列

### Kafka 3.0

**用途**:
- 异步任务
- 事件驱动
- 数据流处理

**Topic命名**: `{project}.{module}.{event}`

---

## 监控和日志

### Prometheus + Grafana

**指标采集和可视化**

### ELK Stack

**日志收集和分析**

### Jaeger

**分布式追踪**

---

## 开发工具

### 必需工具

- **goctl** - Go-Zero代码生成
- **golangci-lint** - 代码检查
- **go test** - 单元测试
- **go-swagger** - API文档

### AI工具

- **Kiro.dev** - 结构化开发
- **Cursor** - IDE集成
- **Claude CLI** - 命令行

---

## 部署

### Docker

**容器化部署**

### Kubernetes

**生产环境编排**

### CI/CD

**GitHub Actions**

---

## 版本要求

| 组件 | 最低版本 | 推荐版本 |
|------|---------|---------|
| Go | 1.21 | 1.21+ |
| Go-Zero | 1.9 | 1.9+ |
| MySQL | 8.0 | 8.0+ |
| Redis | 6.0 | 7.0+ |
| Kafka | 2.8 | 3.0+ |

---

## 依赖管理

使用Go Modules (`go.mod`)

```bash
# 更新依赖
go get -u ./...

# 清理
go mod tidy
```

---

**技术栈选择遵循：稳定性、性能、生态、团队熟悉度**
