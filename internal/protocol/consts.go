package protocol

const (
	clientVersion           = "00.9.30"
	messageSplit            = "\x01"
	attributeSplit          = ","
	messageHeaderBodyLength = 10
	messageHeaderLength     = 21
	defaultPerPageCount     = 2000
)

const (
	msgTypeLoginRequest                         = "00"
	msgTypeLogoutRequest                        = "02"
	msgTypeQueryDividendDataRequest             = "13"
	msgTypeQueryAdjustFactorRequest             = "15"
	msgTypeQueryProfitDataRequest               = "17"
	msgTypeQueryOperationDataRequest            = "19"
	msgTypeQueryGrowthDataRequest               = "21"
	msgTypeQueryDupontDataRequest               = "23"
	msgTypeQueryBalanceDataRequest              = "25"
	msgTypeQueryCashFlowDataRequest             = "27"
	msgTypeQueryPerformanceExpressReportRequest = "29"
	msgTypeQueryForecastReportRequest           = "31"
	msgTypeQueryTradeDatesRequest               = "33"
	msgTypeQueryAllStockRequest                 = "35"
	msgTypeQueryStockBasicRequest               = "45"
	msgTypeQueryDepositRateDataRequest          = "47"
	msgTypeQueryLoanRateDataRequest             = "49"
	msgTypeQueryRequiredReserveRatioDataRequest = "51"
	msgTypeQueryMoneySupplyDataMonthRequest     = "53"
	msgTypeQueryMoneySupplyDataYearRequest      = "55"
	msgTypeQueryStockIndustryRequest            = "59"
	msgTypeQueryHS300StocksRequest              = "61"
	msgTypeQuerySZ50StocksRequest               = "63"
	msgTypeQueryZZ500StocksRequest              = "65"
	msgTypeQueryTerminatedStocksRequest         = "67"
	msgTypeQuerySuspendedStocksRequest          = "69"
	msgTypeQuerySTStocksRequest                 = "71"
	msgTypeQueryStarSTStocksRequest             = "73"
	msgTypeQueryCPIDataRequest                  = "75"
	msgTypeQueryPPIDataRequest                  = "77"
	msgTypeQueryPMIDataRequest                  = "79"
	msgTypeQueryStockConceptRequest             = "81"
	msgTypeQueryStockAreaRequest                = "83"
	msgTypeQueryAMEStockRequest                 = "85"
	msgTypeQueryGEMStockRequest                 = "87"
	msgTypeQuerySHHKStockRequest                = "89"
	msgTypeQuerySZHKStockRequest                = "91"
	msgTypeQueryStocksInRiskRequest             = "93"
	msgTypeGetKDataPlusRequest                  = "95"
	msgTypeGetKDailyDAStockRequest              = "98"
	msgTypeGetKDailyDETFRequest                 = "9A"
	msgTypeGetKDailyAdjustFactorRequest         = "9C"
)

var compressedResponseTypes = map[string]struct{}{
	"96": {},
	"99": {},
	"9B": {},
	"9D": {},
}

func requestTypeByMethod(method string) string {
	switch method {
	case "login":
		return msgTypeLoginRequest
	case "logout":
		return msgTypeLogoutRequest
	case "query_history_k_data_plus":
		return msgTypeGetKDataPlusRequest
	case "query_daily_history_k_AStock":
		return msgTypeGetKDailyDAStockRequest
	case "query_daily_history_k_ETF":
		return msgTypeGetKDailyDETFRequest
	case "query_daily_adjust_factor":
		return msgTypeGetKDailyAdjustFactorRequest
	case "query_trade_dates":
		return msgTypeQueryTradeDatesRequest
	case "query_all_stock":
		return msgTypeQueryAllStockRequest
	case "query_stock_basic":
		return msgTypeQueryStockBasicRequest
	case "query_stock_industry":
		return msgTypeQueryStockIndustryRequest
	case "query_hs300_stocks":
		return msgTypeQueryHS300StocksRequest
	case "query_sz50_stocks":
		return msgTypeQuerySZ50StocksRequest
	case "query_zz500_stocks":
		return msgTypeQueryZZ500StocksRequest
	case "query_terminated_stocks":
		return msgTypeQueryTerminatedStocksRequest
	case "query_suspended_stocks":
		return msgTypeQuerySuspendedStocksRequest
	case "query_st_stocks":
		return msgTypeQuerySTStocksRequest
	case "query_starst_stocks":
		return msgTypeQueryStarSTStocksRequest
	case "query_stock_concept":
		return msgTypeQueryStockConceptRequest
	case "query_stock_area":
		return msgTypeQueryStockAreaRequest
	case "query_ame_stocks":
		return msgTypeQueryAMEStockRequest
	case "query_gem_stocks":
		return msgTypeQueryGEMStockRequest
	case "query_shhk_stocks":
		return msgTypeQuerySHHKStockRequest
	case "query_szhk_stocks":
		return msgTypeQuerySZHKStockRequest
	case "query_stocks_in_risk":
		return msgTypeQueryStocksInRiskRequest
	case "query_dividend_data":
		return msgTypeQueryDividendDataRequest
	case "query_adjust_factor":
		return msgTypeQueryAdjustFactorRequest
	case "query_profit_data":
		return msgTypeQueryProfitDataRequest
	case "query_operation_data":
		return msgTypeQueryOperationDataRequest
	case "query_growth_data":
		return msgTypeQueryGrowthDataRequest
	case "query_dupont_data":
		return msgTypeQueryDupontDataRequest
	case "query_balance_data":
		return msgTypeQueryBalanceDataRequest
	case "query_cash_flow_data":
		return msgTypeQueryCashFlowDataRequest
	case "query_performance_express_report":
		return msgTypeQueryPerformanceExpressReportRequest
	case "query_forecast_report":
		return msgTypeQueryForecastReportRequest
	case "query_deposit_rate_data":
		return msgTypeQueryDepositRateDataRequest
	case "query_loan_rate_data":
		return msgTypeQueryLoanRateDataRequest
	case "query_required_reserve_ratio_data":
		return msgTypeQueryRequiredReserveRatioDataRequest
	case "query_money_supply_data_month":
		return msgTypeQueryMoneySupplyDataMonthRequest
	case "query_money_supply_data_year":
		return msgTypeQueryMoneySupplyDataYearRequest
	case "query_cpi_data":
		return msgTypeQueryCPIDataRequest
	case "query_ppi_data":
		return msgTypeQueryPPIDataRequest
	case "query_pmi_data":
		return msgTypeQueryPMIDataRequest
	default:
		return ""
	}
}

func isCompressedResponseType(msgType string) bool {
	_, ok := compressedResponseTypes[msgType]
	return ok
}
