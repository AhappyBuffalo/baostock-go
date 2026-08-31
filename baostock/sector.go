package baostock

import (
	"github.com/example/baostock-go/internal/request"
	rs "github.com/example/baostock-go/internal/result"
)

func (c *Client) queryDateOnly(method, date string) (*rs.ResultData, error) {
	if err := request.ValidateDate(date); err != nil {
		return nil, err
	}
	return c.SendQuery(method, "1", "2000", date)
}

func (c *Client) queryCodeAndDate(method, code, date string) (*rs.ResultData, error) {
	code, err := request.NormalizeOptionalCode(code)
	if err != nil {
		return nil, err
	}
	if err := request.ValidateDate(date); err != nil {
		return nil, err
	}
	return c.SendQuery(method, "1", "2000", code, date)
}

func (c *Client) QueryStockIndustry(code, date string) (*rs.ResultData, error) {
	return c.queryCodeAndDate("query_stock_industry", code, date)
}

func (c *Client) QueryHS300Stocks(date string) (*rs.ResultData, error) {
	return c.queryDateOnly("query_hs300_stocks", date)
}

func (c *Client) QuerySZ50Stocks(date string) (*rs.ResultData, error) {
	return c.queryDateOnly("query_sz50_stocks", date)
}

func (c *Client) QueryZZ500Stocks(date string) (*rs.ResultData, error) {
	return c.queryDateOnly("query_zz500_stocks", date)
}

func (c *Client) QueryTerminatedStocks(date string) (*rs.ResultData, error) {
	return c.queryDateOnly("query_terminated_stocks", date)
}

func (c *Client) QuerySuspendedStocks(date string) (*rs.ResultData, error) {
	return c.queryDateOnly("query_suspended_stocks", date)
}

func (c *Client) QuerySTStocks(date string) (*rs.ResultData, error) {
	return c.queryDateOnly("query_st_stocks", date)
}

func (c *Client) QueryStarSTStocks(date string) (*rs.ResultData, error) {
	return c.queryDateOnly("query_starst_stocks", date)
}

func (c *Client) QueryStockConcept(code, date string) (*rs.ResultData, error) {
	return c.queryCodeAndDate("query_stock_concept", code, date)
}

func (c *Client) QueryStockArea(code, date string) (*rs.ResultData, error) {
	return c.queryCodeAndDate("query_stock_area", code, date)
}

func (c *Client) QueryAMEStocks(date string) (*rs.ResultData, error) {
	return c.queryDateOnly("query_ame_stocks", date)
}

func (c *Client) QueryGEMStocks(date string) (*rs.ResultData, error) {
	return c.queryDateOnly("query_gem_stocks", date)
}

func (c *Client) QuerySHHKStocks(date string) (*rs.ResultData, error) {
	return c.queryDateOnly("query_shhk_stocks", date)
}

func (c *Client) QuerySZHKStocks(date string) (*rs.ResultData, error) {
	return c.queryDateOnly("query_szhk_stocks", date)
}

func (c *Client) QueryStocksInRisk(date string) (*rs.ResultData, error) {
	return c.queryDateOnly("query_stocks_in_risk", date)
}
