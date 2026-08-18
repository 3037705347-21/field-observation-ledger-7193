# Field Observation Ledger

## 项目目标

Field Observation Ledger 是一个面向生态调查团队的本地 HTTP 服务，用于整理野外观测记录、物种目录、调查地点和审核状态。系统使用内存存储，方便在培训、演示和离线整理场景中快速启动。

## 用户角色

- 观察员：提交观察记录并补充数量、时间和环境信息。
- 分类员：维护物种目录和别名，帮助统一记录口径。
- 审核员：检查记录证据并批准或退回记录。
- 研究协调员：查看地点概览和按物种聚合的观察摘要。

## 核心实体

- Observation：一次在调查地点发现某物种的记录。
- Species：物种名称、分类级别和保护标签。
- Site：调查地点及其生态区域。
- Review：观察记录的审核动作和意见。
- Summary：按地点、物种和状态聚合的统计结果。

## 业务流程

### 记录观察

观察员通过 `POST /api/v1/observations` 提交地点、物种、数量和观察时间。服务校验地点与物种存在后保存草稿记录，并返回 `draft` 状态。

### 审核观察

审核员通过 `POST /api/v1/reviews/{observation-id}/approve` 或 `.../reject` 对草稿记录做决定。服务检查记录当前状态、保存审核事件并更新记录状态。

### 查询与摘要

研究协调员通过 `GET /api/v1/observations` 结合地点、物种和状态筛选记录，通过 `GET /api/v1/summary` 查看总数、已批准数和物种分布。

### 目录维护

分类员通过 `POST /api/v1/species` 和 `POST /api/v1/sites` 创建目录项，查询接口用于确认名称和区域是否可用于新记录。

## 状态与规则

- 观察记录状态只能是 `draft`、`approved` 或 `rejected`。
- 数量必须为正整数，观察时间必须是 RFC3339 时间。
- 已批准或已退回的记录不能重复审核。
- 物种和地点标识符在各自目录中唯一。
- 摘要默认只统计当前可见的观察记录，并分别返回状态计数。

## 接口与验证

- `GET /health` 返回服务健康状态。
- `GET /api/v1/sites`、`GET /api/v1/species` 返回目录。
- `GET /api/v1/observations` 支持筛选。
- `POST /api/v1/observations` 创建草稿记录。
- `POST /api/v1/reviews/{id}/approve` 批准记录。
- `POST /api/v1/reviews/{id}/reject` 退回记录。
- `GET /api/v1/summary` 返回聚合摘要。

健康项目使用 `go build ./...`、`go test ./...` 和 `.go-annotation-runtime.json` 中的 HTTP 检查验证。
