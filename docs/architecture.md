# 架构设计

## 目标

MultiPlatform Publisher 的目标是为创作者提供一条从“模糊想法”到“多平台发布稿”的完整 demo 流程。系统优先保证架构清晰、可演示、可扩展，而不是第一阶段就接入真实平台发布接口。

## 总体流程

```text
RawInput
  -> ContentAnalyzer
  -> StructuredContent
  -> PlatformAdapter Registry
  -> PlatformDraft[]
  -> Preview / Copy / Mock Publish
  -> PublishTask
```

核心原则：

- 原始输入只作为用户表达，不直接绑定某个平台。
- 结构化内容是中间资产，后续所有平台生成都基于它。
- 平台差异由独立适配器负责，主流程不写平台特判。
- 大模型能力可插拔，没有 API Key 时仍能使用本地规则完成 demo。

## 后端模块

### router

负责注册 HTTP 路由，连接 handler 和中间件。

计划接口：

```text
GET  /api/health
POST /api/analyze
POST /api/rewrite
POST /api/publish/mock
GET  /api/publish/tasks
GET  /api/publish/tasks/:id
```

### analyzer

负责把用户的模糊输入整理成结构化内容。

输入：

```text
一段自然语言想法或长文本
```

输出：

```text
主题、受众、内容类型、语气、核心要点、关键词、推荐标题
```

第一阶段实现本地规则版，后续接入大模型 Provider。

### llm

负责抽象外部大模型能力。

计划接口：

```go
type Provider interface {
    Analyze(raw string) (StructuredContent, error)
    Rewrite(content StructuredContent, platform string) (PlatformDraft, error)
}
```

默认 Provider 为 `rule`，保证没有外部密钥时也能运行。

### platform

负责平台风格适配。

计划接口：

```go
type Adapter interface {
    Platform() string
    Rewrite(content StructuredContent) (PlatformDraft, error)
    Validate(draft PlatformDraft) []ValidationIssue
}
```

首批平台：

- wechat：公众号长文风格，结构清晰，正式表达。
- zhihu：知乎问答风格，强调问题、分析和经验总结。
- bilibili：B 站动态或视频简介风格，口语化，有互动引导。
- xiaohongshu：小红书经验分享风格，短句、标签、轻量情绪表达。

### content

负责内容相关数据模型和存储。

计划实体：

- RawInput
- StructuredContent
- PlatformDraft

### publish

负责发布任务。MVP 阶段只做模拟发布，不接真实平台 API。

计划实体：

- PublishTask
- PlatformPublishResult

模拟发布结果包含：

- 平台名称
- 发布状态
- mock URL
- 错误信息
- 创建时间

### database

负责数据库连接和迁移。

第一阶段可以用内存存储推进功能；当发布历史和 demo 稳定后，引入 SQLite 持久化。

## 前端模块

### Workspace

主工作台页面，承载完整 demo 流程。

区域：

- 模糊输入区
- 平台选择区
- 结构化内容区
- 平台草稿预览区
- 发布结果区

### RawInputPanel

负责输入原始想法、内容类型和生成按钮。

### StructuredContentPanel

展示并允许修改系统理解后的内容结构。

### PlatformDraftTabs

按平台展示生成结果，包括标题、正文、标签、字数和校验提示。

### PublishResultPanel

展示模拟发布结果和 mock 链接。

## 数据模型草案

```go
type StructuredContent struct {
    Topic          string
    Audience       string
    ContentType    string
    Tone           string
    CorePoints     []string
    Keywords       []string
    SuggestedTitle string
}

type PlatformDraft struct {
    Platform string
    Title    string
    Body     string
    Tags     []string
    Warnings []ValidationIssue
}

type PublishTask struct {
    ID        string
    Status    string
    Results   []PlatformPublishResult
    CreatedAt time.Time
}
```

## 可扩展性设计

### 新增平台

新增平台不需要修改内容理解器，只需要：

1. 实现 `platform.Adapter`。
2. 在 registry 中注册。
3. 为前端添加平台展示名称和图标。
4. 增加平台规则测试。

### 接入真实发布

真实发布能力可以在后续替换 `MockPublisher`：

```go
type Publisher interface {
    Publish(draft PlatformDraft) (PlatformPublishResult, error)
}
```

不同平台可以独立实现自己的 Publisher，并保留 mock 模式用于演示和测试。

### 接入大模型

大模型只作为 Provider 插件接入，不影响主流程。

配置示例：

```text
LLM_PROVIDER=rule
LLM_API_KEY=
LLM_MODEL=
```

没有 API Key 时使用本地规则；有 API Key 时使用外部模型增强内容理解和改写。

## 非目标

MVP 阶段暂不做：

- 真实登录公众号、知乎、B 站、小红书账号。
- 绕过平台风控或自动化网页发布。
- 多用户权限系统。
- 复杂素材审核和内容合规系统。
- 完整富文本编辑器。

这些能力可以在 demo 稳定后作为扩展方向。
