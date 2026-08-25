基于 Go 实现的机场行李分拣系统项目，一款后端服务，完成值机行李的条码读取、航班匹配、滑槽分配、分拣、复检、转盘切换与分拣审计。

## 运行

本项目使用文件持久化，启动后提供 JSON API 控制台。

```sh
go run ./cmd/bagsort -addr :8080 -data ./data
```

## 主要接口

- `/api/health` 健康检查
- `/api/flights` 航班列表与登记
- `/api/bags/checkin` 值机收运行李
- `/api/sort` 分拣指令下发
- `/api/chutes` 滑槽分配
- `/api/recheck` 复检队列
- `/api/belt` 转盘批次与传输
- `/api/quota` 分拣配额
- `/api/audit` 分拣审计
