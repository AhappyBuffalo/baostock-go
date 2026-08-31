package baostock

import (
	"time"

	"github.com/example/baostock-go/internal/request"
	rs "github.com/example/baostock-go/internal/result"
)

func (c *Client) QueryTradeDates(startDate, endDate string) (*rs.ResultData, error) {
	startDate = request.WithDefaultDate(startDate, "2015-01-01")
	endDate = request.WithDefaultDate(endDate, time.Now().Format("2006-01-02"))
	if err := request.ValidateDateRange(startDate, endDate); err != nil {
		return nil, err
	}
	return c.SendQuery("query_trade_dates", "1", "2000", startDate, endDate)
}

func (c *Client) QueryAllStock(day string) (*rs.ResultData, error) {
	if day == "" {
		day = time.Now().Format("2006-01-02")
	}
	if err := request.ValidateDate(day); err != nil {
		return nil, err
	}
	return c.SendQuery("query_all_stock", "1", "2000", day)
}

func (c *Client) QueryStockBasic(code, codeName string) (*rs.ResultData, error) {
	code, err := request.NormalizeOptionalCode(code)
	if err != nil {
		return nil, err
	}
	return c.SendQuery("query_stock_basic", "1", "2000", code, codeName)
}
