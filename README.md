[baostock](http://baostock.com) A 股金融数据接口的 Go 语言 SDK 实现，协议与 Python SDK 完全兼容。

[baostock](http://baostock.com) A 股金融数据接口的 Go 语言实现，协议与 Python SDK 完全兼容。

## 安装

```bash
go get github.com/example/baostock-go/baostock
```

## 快速开始

```go
package main

import (
    "fmt"
    "log"

    baostock "github.com/example/baostock-go/baostock"
)

func main() {
    client := baostock.NewClient()
    if err := client.Login(); err != nil {
        log.Fatal(err)
    }
    defer client.Logout()

    // 查询日 K 线
    rs, err := client.QueryHistoryKDataPlus(
        "sh.600000",
        "date,open,high,low,close,volume",
        "2024-01-01", "2024-01-10",
        "d", "3", // 日线，不复权
    )
    if err != nil {
        log.Fatal(err)
    }
    for rs.Next() {
        fmt.Println(rs.GetRowData())
    }
}
```

## 客户端配置

通过 Option 函数自定义连接参数：

```go
client := baostock.NewClient(
    baostock.WithServer("public-api.baostock.com:10030"),
    baostock.WithTimeout(30*time.Second),
    baostock.WithUser("anonymous"),
    baostock.WithPassword("123456"),
    baostock.WithAPIKey("your-api-key"),
)
```

| Option | 默认值 | 说明 |
|--------|--------|------|
| `WithServer` | `public-api.baostock.com:10030` | 服务器地址 |
| `WithTimeout` | `30s` | 连接超时 |
| `WithUser` | `anonymous` | 用户名 |
| `WithPassword` | `123456` | 密码 |
| `WithAPIKey` | 空 | API Key（可选） |

## API 概览

所有查询方法返回 `*result.ResultData`，通过迭代器读取数据：

```go
rs, err := client.QueryXxx(...)
for rs.Next() {
    row := rs.GetRowData() // []string
    fmt.Println(rs.Fields) // 字段名列表
}
```

### 行情数据

| 方法 | 说明 |
|------|------|
| `QueryHistoryKDataPlus(code, fields, startDate, endDate, frequency, adjustFlag)` | 历史 K 线（日/周/月/分钟） |
| `QueryDailyHistoryKAStock(date)` | 某日全部 A 股日线 |
| `QueryDailyHistoryKETF(date)` | 某日全部 ETF 日线 |

**frequency**: `d` 日线、`w` 周线、`m` 月线、`5`/`15`/`30`/`60` 分钟线

**adjustFlag**: `1` 后复权、`2` 前复权、`3` 不复权

### 股票元数据

| 方法 | 说明 |
|------|------|
| `QueryStockBasic(code, codeName)` | 股票基本信息（代码、名称、上市/退市日期） |
| `QueryTradeDates(startDate, endDate)` | 交易日历 |
| `QueryAllStock(day)` | 某日所有股票列表 |

### 板块与分类

| 方法 | 说明 |
|------|------|
| `QueryStockIndustry(code, date)` | 行业分类 |
| `QueryStockConcept(code, date)` | 概念板块 |
| `QueryStockArea(code, date)` | 地域板块 |
| `QueryHS300Stocks(date)` | 沪深 300 成分股 |
| `QuerySZ50Stocks(date)` | 上证 50 成分股 |
| `QueryZZ500Stocks(date)` | 中证 500 成分股 |
| `QuerySHHKStocks(date)` | 沪港通标的 |
| `QuerySZHKStocks(date)` | 深港通标的 |
| `QueryAMEStocks(date)` | 主板 A 股 |
| `QueryGEMStocks(date)` | 创业板 |

### 特殊状态股票

| 方法 | 说明 |
|------|------|
| `QuerySTStocks(date)` | ST 股票 |
| `QueryStarSTStocks(date)` | *ST 股票 |
| `QueryTerminatedStocks(date)` | 退市股票 |
| `QuerySuspendedStocks(date)` | 停牌股票 |
| `QueryStocksInRisk(date)` | 风险警示股票 |

### 财务数据

| 方法 | 说明 |
|------|------|
| `QueryProfitData(code, year, quarter)` | 盈利能力 |
| `QueryOperationData(code, year, quarter)` | 营运能力 |
| `QueryGrowthData(code, year, quarter)` | 成长能力 |
| `QueryDupontData(code, year, quarter)` | 杜邦分析 |
| `QueryBalanceData(code, year, quarter)` | 偿债能力 |
| `QueryCashFlowData(code, year, quarter)` | 现金流量 |
| `QueryDividendData(code, year, yearType)` | 分红送股 |
| `QueryPerformanceExpressReport(code, startDate, endDate)` | 业绩快报 |
| `QueryForecastReport(code, startDate, endDate)` | 业绩预告 |

### 估值与复权因子

| 方法 | 说明 |
|------|------|
| `QueryAdjustFactor(code, startDate, endDate)` | 复权因子（按股票） |
| `QueryDailyAdjustFactor(date)` | 复权因子（按日期） |

### 宏观经济

| 方法 | 说明 |
|------|------|
| `QueryDepositRateData(startDate, endDate)` | 存款利率 |
| `QueryLoanRateData(startDate, endDate)` | 贷款利率 |
| `QueryRequiredReserveRatioData(startDate, endDate, yearType)` | 存款准备金率 |
| `QueryMoneySupplyDataMonth(startDate, endDate)` | 货币供应量（月） |
| `QueryMoneySupplyDataYear(startDate, endDate)` | 货币供应量（年） |
| `QueryCPIData(startDate, endDate)` | CPI |
| `QueryPPIData(startDate, endDate)` | PPI |
| `QueryPMIData(startDate, endDate)` | PMI |

## 项目结构

```
baostock-go/
├── baostock/          # 对外 API：Client 及各业务查询方法
│   ├── client.go      # 连接管理、Login/Logout/SendQuery
│   ├── auth.go        # API Key
│   ├── options.go     # Option 函数
│   ├── history.go     # 历史行情
│   ├── metadata.go    # 元数据（股票列表、交易日历）
│   ├── sector.go      # 板块分类
│   ├── valuation.go   # 财务数据、分红、复权因子
│   ├── report.go      # 业绩快报/预告
│   └── macro.go       # 宏观经济数据
├── internal/
│   ├── protocol/      # TCP 协议：消息编解码、压缩、CRC
│   ├── request/       # 参数校验与默认值
│   └── result/        # 返回体结构与迭代器
├── cmd/example/       # 使用示例
└── docs/              # 设计文档
```

## 与 Python SDK 的对应关系

| Python | Go |
|--------|-----|
| `bs.login()` | `client.Login()` |
| `bs.logout()` | `client.Logout()` |
| `bs.query_xxx()` | `client.QueryXxx()` |
| `rs.next()` | `rs.Next()` |
| `rs.get_row_data()` | `rs.GetRowData()` |
| `rs.fields` | `rs.Fields` |
| `rs.error_code` | `rs.ErrorCode` |

Python SDK 中的 `set_API_key()` 对应 `client.SetAPIKey()` 或创建时 `WithAPIKey()`。

## 设计说明

- 消息分隔符 `\x01`，与 Python SDK 协议完全兼容
- 历史行情响应自动 zlib 解压
- 参数校验（日期格式、股票代码格式）在客户端侧完成
- 未实现实时订阅（SubscribeData）和自动翻页并发优化

## License

MIT
