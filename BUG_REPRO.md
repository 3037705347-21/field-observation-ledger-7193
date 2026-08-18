# 修复前故障复现（Docker）

## 项目与标准命令
项目使用 Go HTTP 服务。镜像构建后，在容器内执行 `go build ./...`，再请求摘要接口。

## 环境构建与编译
当前平台的镜像构建和容器内 `go build ./...` 已通过。

## 故障触发步骤
发送 `GET /api/v1/summary` 请求，服务处理已有观察记录。

## 实际错误输出
`panic: assignment to entry in nil map`

## 期望行为
摘要接口应返回 HTTP 200，并将 `by_species` 和 `by_status` 序列化为 JSON 对象。
