# 修复前故障复现（Docker）

## 项目与标准命令
项目使用 Go HTTP 服务。镜像构建后，在容器内执行 `go build ./...`，再执行定向审核请求测试。

## 环境构建与编译
当前平台的镜像构建和容器内 `go build ./...` 已通过。

## 故障触发步骤
向 `POST /api/v1/reviews/obs-missing/approve` 发送 JSON 请求体 `{"reviewer":"qa"}`。

## 实际错误输出
`expected missing observation to return 404, got 400: {"error":"observation not found"}`

## 期望行为
不存在的观察记录应返回 HTTP 404，而不是 HTTP 400。
