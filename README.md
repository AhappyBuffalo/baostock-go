# baostock-go

[baostock](http://baostock.com) A 股金融数据接口的 Go 语言 SDK 实现，协议与 Python SDK 完全兼容。

## 安装

```bash
go get github.com/example/baostock-go/baostock
```

## 快速开始

```go
client := baostock.NewClient()
if err := client.Login(); err != nil {
    log.Fatal(err)
}
defer client.Logout()

rs, err := client.QueryHistoryKDataPlus(
    "sh.600000", "date,open,high,low,close,volume",
    "2024-01-01", "2024-01-10", "d", "3",
)
for rs.Next() {
    fmt.Println(rs.GetRowData())
}
```

## 客户端配置

```go
client := baostock.NewClient(
    baostock.WithServer("public-api.baostock.com:10030"),
    baostock.WithTimeout(30*time.Second),
    baostock.WithUser("anonymous"),
    baostock.WithPassword("123456"),
    baostock.WithAPIKey("your-api-key"),
)
```

## API

所有查询返回 `*result.ResultData`，通过 `rs.Next()` / `rs.GetRowData()` 迭代，`rs.Fields` 获取字段名。

### 行情

| 方法 | 说明 |
|------|------|
| `QueryHistoryKDataPlus(code, fields, startDate, endDate, frequency, adjustFlag)` | 历史 K 线 |
| `QueryDailyHistoryKAStock(date)` | 某日全部 A 股日线 |
| `QueryDailyHistoryKETF(date)` | 某日全部 ETF 日线 |

frequency: `d` 日线 / `w` 周线 / `m` 月线 / `5` `15` `30` `60` 分钟线。adjustFlag: `1` 后复权 / `2` 前复权 / `3` 不复权。

### 元数据

| 方法 | 说明 |
|------|------|
| `QueryStockBasic(code, codeName)` | 股票基本信息 |
| `QueryTradeDates(startDate, endDate)` | 交易日历 |
| `QueryAllStock(day)` | 某日所有股票列表 |

### 板块分类

| 方法 | 说明 |
|------|------|
| `QueryStockIndustry(code, date)` | 行业 |
| `QueryStockConcept(code, date)` | 概念 |
| `QueryStockArea(code, date)` | 地域 |
| `QueryHS300Stocks` / `QuerySZ50Stocks` / `QueryZZ500Stocks` | 指数成分股 |
| `QuerySHHKStocks` / `QuerySZHKStocks` | 港股通标的 |
| `QueryAMEStocks` / `QueryGEMStocks` | 主板 / 创业板 |
| `QuerySTStocks` / `QueryStarSTStocks` | ST / *ST |
| `QueryTerminatedStocks` / `QuerySuspendedStocks` / `QueryStocksInRisk` | 退市 / 停牌 / 风险警示 |

### 财务

| 方法 | 说明 |
|------|------|
| `QueryProfitData` / `QueryOperationData` / `QueryGrowthData` | 盈利 / 营运 / 成长 |
| `QueryDupontData` / `QueryBalanceData` / `QueryCashFlowData` | 杜邦 / 偿债 / 现金流 |
| `QueryDividendData(code, year, yearType)` | 分红送股 |
| `QueryPerformanceExpressReport` / `QueryForecastReport` | 业绩快报 / 预告 |
| `QueryAdjustFactor` / `QueryDailyAdjustFactor` | 复权因子 |

### 宏观经济

`QueryDepositRateData` / `QueryLoanRateData` / `QueryRequiredReserveRatioData` / `QueryMoneySupplyDataMonth` / `QueryMoneySupplyDataYear` / `QueryCPIData` / `QueryPPIData` / `QueryPMIData`

## 项目结构

```
baostock/
├── client.go      # 连接管理
├── history.go     # 历史行情
├── metadata.go    # 元数据
├── sector.go      # 板块分类
├── valuation.go   # 财务数据、复权因子
├── report.go      # 业绩快报/预告
└── macro.go       # 宏观经济
internal/
├── protocol/      # TCP 协议编解码、zlib 压缩
├── request/       # 参数校验
└── result/        # 返回体与迭代器
```

## License

MIT
