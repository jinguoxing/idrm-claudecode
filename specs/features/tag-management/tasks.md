# Tasks: Data Tag Management

## Task 1: 创建数据库表
**Status**: Not Started
**Depends on**: -
**Files**:
- `migrations/tag_management/20250129_create_tags.sql`
- `migrations/tag_management/20250129_create_resource_tags.sql`

**Acceptance Criteria**:
- [ ] tags表创建成功，包含所有字段和索引
- [ ] resource_tags表创建成功，包含唯一索引
- [ ] 外键关系正确建立

**Estimated Lines**: 50

---

## Task 2: 实现Tag Model - 基础结构和接口
**Status**: Not Started
**Depends on**: Task 1
**Files**:
- `model/tag_management/tag/interface.go`
- `model/tag_management/tag/types.go`
- `model/tag_management/tag/vars.go`

**Acceptance Criteria**:
- [ ] TagModel接口定义完整
- [ ] Tag结构体定义，包含GORM标签
- [ ] 常量和错误定义
- [ ] TableName方法实现

**Estimated Lines**: 60

---

## Task 3: 实现Tag Model - GORM实现
**Status**: Not Started
**Depends on**: Task 2
**Files**:
- `model/tag_management/tag/factory.go`
- `model/tag_management/tag/gorm_dao.go`

**Acceptance Criteria**:
- [ ] Factory函数NewTagModel
- [ ] Insert方法实现
- [ ] FindOne/FindByName方法实现
- [ ] Update方法实现
- [ ] Delete方法实现
- [ ] List方法（分页）
- [ ] Search方法（关键词搜索）
- [ ] WithTx和Trans方法实现

**Estimated Lines**: 200

---

## Task 4: 实现ResourceTag Model - 基础结构
**Status**: Not Started
**Depends on**: Task 1
**Files**:
- `model/tag_management/resource_tag/interface.go`
- `model/tag_management/resource_tag/types.go`
- `model/tag_management/resource_tag/vars.go`

**Acceptance Criteria**:
- [ ] ResourceTagModel接口定义完整
- [ ] ResourceTag结构体定义
- [ ] 常量和错误定义

**Estimated Lines**: 80

---

## Task 5: 实现ResourceTag Model - GORM实现
**Status**: Not Started
**Depends on**: Task 4
**Files**:
- `model/tag_management/resource_tag/factory.go`
- `model/tag_management/resource_tag/gorm_dao.go`

**Acceptance Criteria**:
- [ ] Assign/Unassign方法
- [ ] BatchAssign/BatchUnassign方法
- [ ] ReplaceTags方法
- [ ] FindByResource/FindByTag方法
- [ ] FindByTags方法（多标签AND查询）
- [ ] CountByTag方法

**Estimated Lines**: 180

---

## Task 6: 定义API Types
**Status**: Not Started
**Depends on**: -
**Files**:
- `api/internal/types/tag_management_types.go`

**Acceptance Criteria**:
- [ ] CreateTagReq/Resp定义
- [ ] GetTagReq/Resp定义
- [ ] ListTagsReq/Resp定义
- [ ] UpdateTagReq/Resp定义
- [ ] DeleteTagReq/Resp定义
- [ ] AssignTagsReq/Resp定义
- [ ] UnassignTagsReq/Resp定义
- [ ] SearchByTagsReq/Resp定义

**Estimated Lines**: 100

---

## Task 7: 实现创建标签Handler和Logic
**Status**: Not Started
**Depends on**: Task 3, Task 6
**Files**:
- `api/internal/handler/tag_management/createtaghandler.go`
- `api/internal/logic/tag_management/createtaglogic.go`

**Acceptance Criteria**:
- [ ] Handler解析请求并调用Logic
- [ ] Logic验证标签名称格式
- [ ] Logic检查名称唯一性
- [ ] Logic设置默认颜色
- [ ] Logic调用Model插入

**Estimated Lines**: 80

---

## Task 8: 实现标签列表Handler和Logic
**Status**: Not Started
**Depends on**: Task 3, Task 6
**Files**:
- `api/internal/handler/tag_management/listtagshandler.go`
- `api/internal/logic/tag_management/listtagslogic.go`

**Acceptance Criteria**:
- [ ] Handler解析分页参数
- [ ] Logic支持关键词搜索
- [ ] Logic支持状态筛选
- [ ] 返回分页结果

**Estimated Lines**: 60

---

## Task 9: 实现标签详情Handler和Logic
**Status**: Not Started
**Depends on**: Task 3, Task 5, Task 6
**Files**:
- `api/internal/handler/tag_management/gettaghandler.go`
- `api/internal/logic/tag_management/gettaglogic.go`

**Acceptance Criteria**:
- [ ] Handler解析ID参数
- [ ] Logic调用Model查询
- [ ] 统计标签使用次数

**Estimated Lines**: 50

---

## Task 10: 实现更新标签Handler和Logic
**Status**: Not Started
**Depends on**: Task 3, Task 6
**Files**:
- `api/internal/handler/tag_management/updatetaghandler.go`
- `api/internal/logic/tag_management/updatetaglogic.go`

**Acceptance Criteria**:
- [ ] Handler解析请求
- [ ] Logic验证标签存在
- [ ] Logic检查名称唯一性

**Estimated Lines**: 70

---

## Task 11: 实现删除标签Handler和Logic
**Status**: Not Started
**Depends on**: Task 3, Task 5, Task 6
**Files**:
- `api/internal/handler/tag_management/deletetaghandler.go`
- `api/internal/logic/tag_management/deletetaglogic.go`

**Acceptance Criteria**:
- [ ] Handler解析ID参数
- [ ] Logic检查标签是否被使用
- [ ] 被使用时返回错误

**Estimated Lines**: 60

---

## Task 12: 实现为数据打标签Handler和Logic
**Status**: Not Started
**Depends on**: Task 3, Task 5, Task 6
**Files**:
- `api/internal/handler/tag_management/assigntagshandler.go`
- `api/internal/logic/tag_management/assigntagslogic.go`

**Acceptance Criteria**:
- [ ] Handler解析请求
- [ ] Logic验证所有标签ID存在
- [ ] Logic调用BatchAssign

**Estimated Lines**: 70

---

## Task 13: 实现移除标签Handler和Logic
**Status**: Not Started
**Depends on**: Task 5, Task 6
**Files**:
- `api/internal/handler/tag_management/unassigntagshandler.go`
- `api/internal/logic/tag_management/unassigntagslogic.go`

**Acceptance Criteria**:
- [ ] Handler解析请求
- [ ] Logic调用BatchUnassign

**Estimated Lines**: 40

---

## Task 14: 实现按标签搜索Handler和Logic
**Status**: Not Started
**Depends on**: Task 5, Task 6
**Files**:
- `api/internal/handler/tag_management/searchbytagshandler.go`
- `api/internal/logic/tag_management/searchbytagslogic.go`

**Acceptance Criteria**:
- [ ] Handler解析标签ID列表
- [ ] Logic调用FindByTags获取资源ID列表
- [ ] 返回分页结果

**Estimated Lines**: 80

---

## Task 15: 更新ServiceContext
**Status**: Not Started
**Depends on**: Task 3, Task 5
**Files**:
- `api/internal/svc/servicecontext.go`

**Acceptance Criteria**:
- [ ] 添加TagModel字段
- [ ] 添加ResourceTagModel字段
- [ ] 在initialize中初始化Model

**Estimated Lines**: 30

---

## Task 16: 定义错误码
**Status**: Not Started
**Depends on**: -
**Files**:
- `pkg/errorx/tag_errors.go`

**Acceptance Criteria**:
- [ ] 定义标签相关错误码 (31000-31999)
- [ ] 定义关联相关错误码 (32000-32999)

**Estimated Lines**: 30

---

## Task 17: API路由注册
**Status**: Not Started
**Depends on**: Task 7-14, Task 15
**Files**:
- `api/tag_management.api`
- `api/internal/router/routes.go`

**Acceptance Criteria**:
- [ ] 定义API规范
- [ ] 注册所有路由

**Estimated Lines**: 80

---

## Task 18: 编写Model层单元测试
**Status**: Not Started
**Depends on**: Task 3, Task 5
**Files**:
- `model/tag_management/tag/gorm_dao_test.go`
- `model/tag_management/resource_tag/gorm_dao_test.go`

**Acceptance Criteria**:
- [ ] TagModel CRUD测试
- [ ] ResourceTagModel关联测试
- [ ] 覆盖率 > 80%

**Estimated Lines**: 200

---

## Task 19: 编写Logic层单元测试
**Status**: Not Started
**Depends on**: Task 7-14
**Files**:
- `api/internal/logic/tag_management/*_test.go`

**Acceptance Criteria**:
- [ ] 所有Logic方法测试
- [ ] Mock Model层
- [ ] 覆盖率 > 80%

**Estimated Lines**: 250

---

## Task 20: 集成测试
**Status**: Not Started
**Depends on**: Task 1-19
**Files**:
- `tests/integration/tag_management_test.go`

**Acceptance Criteria**:
- [ ] 端到端测试
- [ ] API响应验证

**Estimated Lines**: 150

---

## 任务统计

- 总任务数: 20
- 预估总行数: ~2100行
- 关键路径: Task 1 → Task 2/4 → Task 3/5 → Task 7-14 → Task 17 → Task 20
