# field-observation-ledger__005 Docker 交付说明

## 项目概览
- Field Observation Ledger is an offline-friendly Go HTTP service for ecological field teams. It keeps observation records connected to species and field sites, then supports review 
- Go module: `example.com/field-observation-ledger`

## 标准命令

```bash
go build ./...
go test ./...
```

## 实际启动入口

```bash
go run ./cmd/observer
```

## Docker 构建

```bash
./build_benzhi_docker.sh field-observation-ledger__005-benzhi linux/amd64
docker run --rm -it field-observation-ledger__005-benzhi bash
```

## 环境

- 基础镜像: `golang:1.26.5`
- 依赖在镜像构建阶段预下载，容器内可直接执行 Go 构建和测试命令。
- 代码目录: `/app`
- 源码中检测到的服务端口: `30`, `18080`
