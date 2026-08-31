# baostock-go 设计说明

本 SDK 复刻 Python baostock 的协议与主要能力，采用 Go 风格接口分层：

- `internal/protocol`: 负责 TCP 消息头、分隔符拼接、压缩解压、请求发送
- `internal/result`: 返回体结构与基础迭代逻辑
- `internal/request`: 参数校验与默认值处理
- `baostock`: 对外暴露的 `Client` 与各业务查询方法

## 重要约定

- 消息分隔符使用 `\x01`，与 Python SDK 保持一致。
- 成功响应解析顺序以 Python 源码为准（method/userid/page/data/meta）。
- 历史行情相关响应会走 zlib 解压。
