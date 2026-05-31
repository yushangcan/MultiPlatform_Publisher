# MultiPlatform Publisher

多平台内容发布工具，面向需要同时运营公众号、知乎、B 站、小红书等平台的创作者。项目目标是把一段模糊的内容想法整理成结构化稿件，再生成适合不同平台风格的发布内容，并提供预览、复制和模拟发布能力。

## 项目背景

很多创作者在多个平台发布同一主题内容时，需要反复调整标题、正文结构、语气、标签和互动引导。这个过程耗时、容易遗漏平台规则，也不利于维护统一的内容资产。

本项目的核心思路不是简单复制粘贴文章，而是把创作者的输入拆成三步处理：

```text
模糊输入 -> 内容理解与结构化 -> 多平台风格改写 -> 预览 / 复制 / 模拟发布
```

## MVP 范围

第一版 demo 聚焦完整闭环，不追求真实接入所有平台发布 API。

- 支持用户输入一段模糊想法或长文本。
- 自动整理出主题、受众、内容类型、核心要点、关键词和推荐标题。
- 生成公众号、知乎、B 站、小红书四个平台的差异化草稿。
- 支持多平台预览、字数提示、平台校验提示和一键复制。
- 支持模拟发布，生成每个平台的发布状态和 mock 链接。
- 提供可扩展平台适配器架构，后续可以增加更多平台。

## 技术栈规划

后端：

- Go
- Gin
- SQLite
- 分层结构：handler、service、repository、adapter
- 可选大模型 Provider，默认提供本地规则 fallback

前端：

- Vue 3
- TypeScript
- Vite
- Element Plus 或 Naive UI
- 单页工作台交互

## 核心功能流程

1. 用户输入模糊内容，例如一段选题想法、学习笔记或产品介绍。
2. 后端内容理解器将原始输入整理为 `StructuredContent`。
3. 用户可以在前端修改系统理解结果。
4. 平台适配器根据结构化内容生成不同平台的草稿。
5. 用户在前端预览、复制或编辑草稿。
6. 模拟发布服务生成发布任务、平台状态和 mock 链接。
7. 发布历史页展示任务记录和结果。

## 目录规划

```text
backend/
  cmd/server/
  internal/
    analyzer/
    config/
    content/
    database/
    llm/
    platform/
    publish/
    router/

frontend/
  src/
    api/
    components/
    pages/
    stores/

docs/
  architecture.md
  pr-plan.md
  demo-script.md
```

## 扩展更多平台

平台能力通过适配器抽象。新增平台时，优先新增一个平台适配器，而不是修改主流程。

```go
type Adapter interface {
    Platform() string
    Rewrite(content StructuredContent) (PlatformDraft, error)
    Validate(draft PlatformDraft) []ValidationIssue
}
```

后续新增平台的大致步骤：

1. 新增平台配置和适配器实现。
2. 注册到平台 registry。
3. 增加前端平台选项和预览标识。
4. 补充平台规则说明和测试用例。

## 开发方式

本项目按小粒度 PR 持续开发。每个 PR 只完成一个明确目标，合并后主分支应保持可运行。

PR 描述需要包含：

- 功能描述
- 实现思路
- 测试方式
- 影响范围

详细拆分见 [docs/pr-plan.md](docs/pr-plan.md)。

## 运行方式

后端运行方式：

```bash
cd backend
go mod tidy
go run ./cmd/server
```

健康检查：

```bash
curl http://localhost:8080/api/health
```

预期返回：

```json
{"status":"ok"}
```

内容理解接口：

```bash
curl -X POST http://localhost:8080/api/analyze \
  -H "Content-Type: application/json" \
  -d '{"input":"我想写一篇关于大学生暑假提升自己的内容，主要讲学习 Go、做项目和准备简历，语气希望轻松一点。"}'
```

平台列表接口：

```bash
curl http://localhost:8080/api/platforms
```

平台改写接口：

```bash
curl -X POST http://localhost:8080/api/rewrite \
  -H "Content-Type: application/json" \
  -d '{"content":{"topic":"大学生暑假提升自己","audience":"大学生","content_type":"经验分享","tone":"轻松、实用","core_points":["学习 Go","完成项目 demo","准备简历"],"keywords":["大学生","Go","项目","简历"],"suggested_title":"大学生暑假提升自己：一篇适合多平台发布的内容"},"platforms":["wechat","zhihu","bilibili","xiaohongshu"]}'
```

前端运行方式会在后续 PR 补充：

```bash
cd frontend
npm install
npm run dev
```

## 第三方依赖说明

当前阶段尚未引入业务代码依赖。后续引入的 Go、前端和大模型相关依赖会在 README 中持续更新，并注明项目原创功能边界。

## 项目文档

- [架构设计](docs/architecture.md)
- [PR 开发计划](docs/pr-plan.md)
