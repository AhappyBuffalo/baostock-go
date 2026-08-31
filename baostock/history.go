package baostock

import (
	"fmt"
	"time"

	"github.com/AhappyBuffalo/baostock-go/internal/request"
	rs "github.com/AhappyBuffalo/baostock-go/internal/result"
)

func (c *Client) QueryHistoryKDataPlus(code, fields, startDate, endDate, frequency, adjustFlag string) (*rs.ResultData, error) {
	code, err := request.NormalizeCode(code)
	if err != nil {
		return nil, err
	}
	if fields == "" {
		return nil, fmt.Errorf("指示简称不能为空，请检查")
	}
	startDate = request.WithDefaultDate(startDate, "2015-01-01")
	endDate = request.WithDefaultDate(endDate, time.Now().Format("2006-01-02"))
	if err := request.ValidateDateRange(startDate, endDate); err != nil {
		return nil, err
	}
	if frequency == "" {
		return nil, fmt.Errorf("数据类型（frequency）不可为空，请检查")
	}
	if adjustFlag == "" {
		return nil, fmt.Errorf("复权类型（adjustflag）不可为空，请检查")
	}
	return c.SendQuery("query_history_k_data_plus", "1", "2000", code, fields, startDate, endDate, frequency, adjustFlag)
}

func (c *Client) QueryDailyHistoryKAStock(date string) (*rs.ResultData, error) {
	if date == "" {
		date = time.Now().Format("2006-01-02")
	}
	if err := request.ValidateDate(date); err != nil {
		return nil, err
	}
	return c.SendQuery("query_daily_history_k_AStock", date)
}

func (c *Client) QueryDailyHistoryKETF(date string) (*rs.ResultData, error) {
	if date == "" {
		date = time.Now().Format("2006-01-02")
	}
	if err := request.ValidateDate(date); err != nil {
		return nil, err
	}
	return c.SendQuery("query_daily_history_k_ETF", date)
}
