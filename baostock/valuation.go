package baostock

import (
	"time"

	"github.com/AhappyBuffalo/baostock-go/internal/request"
	rs "github.com/AhappyBuffalo/baostock-go/internal/result"
)

func (c *Client) QueryDividendData(code, year, yearType string) (*rs.ResultData, error) {
	code, err := request.NormalizeCode(code)
	if err != nil {
		return nil, err
	}
	if year == "" {
		year = time.Now().Format("2006")
	}
	if yearType == "" {
		yearType = "report"
	}
	return c.SendQuery("query_dividend_data", "1", "2000", code, year, yearType)
}

func (c *Client) QueryAdjustFactor(code, startDate, endDate string) (*rs.ResultData, error) {
	code, err := request.NormalizeCode(code)
	if err != nil {
		return nil, err
	}
	startDate = request.WithDefaultDate(startDate, "2015-01-01")
	endDate = request.WithDefaultDate(endDate, time.Now().Format("2006-01-02"))
	return c.SendQuery("query_adjust_factor", "1", "2000", code, startDate, endDate)
}

func (c *Client) QueryDailyAdjustFactor(date string) (*rs.ResultData, error) {
	if date == "" {
		date = time.Now().Format("2006-01-02")
	}
	return c.SendQuery("query_daily_adjust_factor", date)
}

func (c *Client) QueryProfitData(code, year, quarter string) (*rs.ResultData, error) {
	code, err := request.NormalizeCode(code)
	if err != nil {
		return nil, err
	}
	if year == "" {
		year = time.Now().Format("2006")
	}
	if quarter == "" {
		quarter = currentQuarter()
	}
	return c.SendQuery("query_profit_data", "1", "2000", code, year, quarter)
}

func (c *Client) QueryOperationData(code, year, quarter string) (*rs.ResultData, error) {
	code, err := request.NormalizeCode(code)
	if err != nil {
		return nil, err
	}
	if year == "" {
		year = time.Now().Format("2006")
	}
	if quarter == "" {
		quarter = currentQuarter()
	}
	return c.SendQuery("query_operation_data", "1", "2000", code, year, quarter)
}

func (c *Client) QueryGrowthData(code, year, quarter string) (*rs.ResultData, error) {
	code, err := request.NormalizeCode(code)
	if err != nil {
		return nil, err
	}
	if year == "" {
		year = time.Now().Format("2006")
	}
	if quarter == "" {
		quarter = currentQuarter()
	}
	return c.SendQuery("query_growth_data", "1", "2000", code, year, quarter)
}

func (c *Client) QueryDupontData(code, year, quarter string) (*rs.ResultData, error) {
	code, err := request.NormalizeCode(code)
	if err != nil {
		return nil, err
	}
	if year == "" {
		year = time.Now().Format("2006")
	}
	if quarter == "" {
		quarter = currentQuarter()
	}
	return c.SendQuery("query_dupont_data", "1", "2000", code, year, quarter)
}

func (c *Client) QueryBalanceData(code, year, quarter string) (*rs.ResultData, error) {
	code, err := request.NormalizeCode(code)
	if err != nil {
		return nil, err
	}
	if year == "" {
		year = time.Now().Format("2006")
	}
	if quarter == "" {
		quarter = currentQuarter()
	}
	return c.SendQuery("query_balance_data", "1", "2000", code, year, quarter)
}

func (c *Client) QueryCashFlowData(code, year, quarter string) (*rs.ResultData, error) {
	code, err := request.NormalizeCode(code)
	if err != nil {
		return nil, err
	}
	if year == "" {
		year = time.Now().Format("2006")
	}
	if quarter == "" {
		quarter = currentQuarter()
	}
	return c.SendQuery("query_cash_flow_data", "1", "2000", code, year, quarter)
}

func currentQuarter() string {
	m := time.Now().Month()
	switch {
	case m <= 3:
		return "1"
	case m <= 6:
		return "2"
	case m <= 9:
		return "3"
	default:
		return "4"
	}
}
