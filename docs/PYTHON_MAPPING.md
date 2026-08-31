# Python -> Go 函数映射

| Python 函数 | Go 方法 | 备注 |
| --- | --- | --- |
| `login` | `Client.Login()` | 默认 anonymous，可 WithUser/WithPassword |
| `logout` | `Client.Logout()` | 同时关闭 TCP |
| `set_API_key` | `Client.SetAPIKey()` | 通过 option 或运行时设置 |
| `query_history_k_data_plus` | `Client.QueryHistoryKDataPlus()` | 含字段/频率/复权 |
| `query_daily_history_k_AStock` | `Client.QueryDailyHistoryKAStock()` | 无分页 |
| `query_daily_history_k_ETF` | `Client.QueryDailyHistoryKETF()` | 无分页 |
| 其余 query_* | 对应 Client.QueryXxx() | 参数语义保持一致 |

## 未覆盖项

- 实时订阅相关（Python SDK 中有 SubscibeData，当前未实现）
- 自动翻页并发优化（当前返回当页/当批结果）
