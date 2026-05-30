# PR 开发计划

本项目采用小粒度 PR 持续开发。每个 PR 只完成一个明确目标，合并后主分支应保持可运行或至少保持文档状态一致。

## PR 规范

每个 PR 描述应包含：

```md
## 功能描述
本 PR 实现了 xxx 功能，用于 xxx。

## 实现思路
- xxx
- xxx

## 测试方式
- [ ] go test ./...
- [ ] npm run dev
- [ ] 手动验证 xxx

## 影响范围
- 后端：xxx
- 前端：xxx
- 文档：xxx
```

## 阶段一：项目骨架与后端核心

### PR 1：完善项目规划文档与基础 README

目标：明确项目方向、MVP 范围、架构目标和后续 PR 拆分。

内容：

- 重写 README。
- 新增架构设计文档。
- 新增 PR 开发计划。

测试方式：

- 人工检查文档是否能说明项目目标和后续计划。

### PR 2：创建 Go 后端基础服务

目标：让后端服务可以启动。

内容：

- 新增 `backend/go.mod`。
- 新增 `backend/cmd/server/main.go`。
- 新增路由模块。
- 新增健康检查接口。

接口：

```text
GET /api/health
```

测试方式：

```bash
cd backend
go mod tidy
go test ./...
go run ./cmd/server
```

### PR 3：定义核心业务模型

目标：定义内容理解、平台草稿和发布任务的数据结构。

内容：

- RawInput
- StructuredContent
- PlatformDraft
- RewriteRequest
- RewriteResponse
- PublishTask

测试方式：

```bash
cd backend
go test ./...
```

### PR 4：实现本地规则版内容理解器

目标：没有大模型 API Key 时也能完成 demo。

内容：

- `ContentAnalyzer` 接口。
- `RuleAnalyzer` 实现。
- `/api/analyze` 接口。
- 基础单元测试。

测试方式：

```bash
cd backend
go test ./...
```

手动验证：

```text
输入一段模糊想法，接口返回主题、核心要点、关键词和推荐标题。
```

### PR 5：实现平台适配器接口与注册器

目标：建立可扩展的平台改写架构。

内容：

- `platform.Adapter` 接口。
- 平台 registry。
- 平台列表接口或内部查询能力。

测试方式：

```bash
cd backend
go test ./...
```

## 阶段二：平台改写能力

### PR 6：实现公众号和知乎适配器

目标：完成两个主要图文平台的改写能力。

内容：

- 公众号适配器。
- 知乎适配器。
- 平台校验规则。
- `/api/rewrite` 接口初版。

测试方式：

```bash
cd backend
go test ./...
```

手动验证：

```text
同一份结构化内容能生成公众号和知乎两种明显不同的稿件。
```

### PR 7：实现 B 站和小红书适配器

目标：补齐四个平台，形成 demo 核心能力。

内容：

- B 站适配器。
- 小红书适配器。
- 标签和互动引导规则。

测试方式：

```bash
cd backend
go test ./...
```

手动验证：

```text
四个平台输出在语气、结构和标签上有明显差异。
```

### PR 8：接入可选大模型 Provider

目标：为模糊输入理解提供增强路径，同时保留本地 fallback。

内容：

- `llm.Provider` 接口。
- `rule` Provider。
- 外部大模型 Provider 占位或初版实现。
- `.env.example`。
- 配置读取模块。

测试方式：

```bash
cd backend
go test ./...
```

手动验证：

```text
无 API Key 时仍然使用规则模式正常运行。
```

## 阶段三：前端工作台

### PR 9：创建 Vue 前端基础项目

目标：前端项目可以启动，并能访问后端健康检查。

内容：

- Vue 3 + Vite + TypeScript 基础结构。
- API client。
- 基础页面布局。

测试方式：

```bash
cd frontend
npm install
npm run dev
```

### PR 10：实现模糊输入工作台

目标：用户可以输入原始想法并选择目标平台。

内容：

- 原始输入组件。
- 平台多选组件。
- 生成按钮和 loading 状态。
- 错误提示。

测试方式：

```text
浏览器输入内容，点击生成按钮能调用后端接口。
```

### PR 11：实现结构化内容展示与编辑

目标：展示系统理解结果，并允许用户调整。

内容：

- 主题、受众、内容类型、语气展示。
- 核心要点和关键词展示。
- 字段编辑能力。

测试方式：

```text
输入模糊内容后，页面能展示结构化结果；修改字段后可继续生成平台稿件。
```

### PR 12：实现多平台稿件预览与复制

目标：完成核心 demo 展示。

内容：

- 平台 tab 或分栏预览。
- 标题、正文、标签展示。
- 字数和校验提示。
- 一键复制。

测试方式：

```text
同一份输入能生成四个平台草稿，并可复制内容。
```

## 阶段四：发布闭环与完善

### PR 13：实现模拟发布后端

目标：为 demo 提供发布闭环。

内容：

- `Publisher` 接口。
- `MockPublisher` 实现。
- 发布任务模型。
- 发布任务接口。

接口：

```text
POST /api/publish/mock
GET  /api/publish/tasks
GET  /api/publish/tasks/:id
```

测试方式：

```bash
cd backend
go test ./...
```

### PR 14：前端接入模拟发布结果

目标：用户可以一键模拟发布并查看结果。

内容：

- 发布按钮。
- 平台发布状态展示。
- mock 链接展示。
- 失败状态展示。

测试方式：

```text
生成稿件后点击模拟发布，页面展示每个平台发布结果。
```

### PR 15：实现发布历史记录

目标：提升作品完整度。

内容：

- 发布历史列表。
- 发布详情。
- 任务状态和平台结果展示。

测试方式：

```text
发布多次后，历史记录中可以查看之前的任务。
```

### PR 16：增加 SQLite 持久化

目标：让数据在后端重启后仍可保留。

内容：

- 数据库连接。
- 内容记录表。
- 发布任务表。
- 发布结果表。

测试方式：

```bash
cd backend
go test ./...
go run ./cmd/server
```

手动验证：

```text
发布记录在服务重启后仍存在。
```

### PR 17：优化前端交互和视觉

目标：让 demo 更像完整产品。

内容：

- 顶部导航。
- 工作台布局优化。
- loading、empty、error 状态。
- 平台标识和响应式布局。

测试方式：

```text
在桌面浏览器完整走通输入、生成、预览、发布流程。
```

### PR 18：补充测试、示例输入和演示脚本

目标：方便评委复现。

内容：

- 后端核心单元测试补充。
- 示例输入文档。
- demo 视频脚本。

测试方式：

```bash
cd backend
go test ./...
```

### PR 19：完善 README 和依赖说明

目标：满足提交规范。

内容：

- 完整运行说明。
- 第三方依赖说明。
- 原创功能边界说明。
- 扩展平台说明。
- demo 视频链接占位。

测试方式：

```text
按照 README 从零启动项目。
```

### PR 20：最终演示版本整理

目标：保证主分支稳定可演示。

内容：

- 修复 demo 流程 bug。
- 统一端口和配置。
- 补充最终截图或说明。
- 检查 README、docs 和代码一致性。

测试方式：

```bash
cd backend
go test ./...
go run ./cmd/server

cd frontend
npm install
npm run dev
```

手动验证完整流程：

```text
模糊输入 -> 结构化理解 -> 多平台生成 -> 复制 -> 模拟发布 -> 历史记录
```

## 推荐分支命名

```text
docs/project-plan
backend/bootstrap
backend/core-models
backend/rule-analyzer
backend/platform-adapter
backend/wechat-zhihu
backend/bilibili-xiaohongshu
backend/llm-provider
frontend/bootstrap
frontend/workspace-input
frontend/structured-content
frontend/platform-preview
backend/mock-publish
frontend/mock-publish
frontend/publish-history
backend/sqlite
frontend/polish
docs/demo-assets
docs/final-readme
release/demo-ready
```
