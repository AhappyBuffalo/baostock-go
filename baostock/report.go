package baostock

import (
	"time"

	"github.com/AhappyBuffalo/baostock-go/internal/request"
	rs "github.com/AhappyBuffalo/baostock-go/internal/result"
)

func (c *Client) QueryPerformanceExpressReport(code, startDate, endDate string) (*rs.ResultData, error) {
	code, err := request.NormalizeCode(code)
	if err != nil {
		return nil, err
	}
	startDate = request.WithDefaultDate(startDate, "2015-01-01")
	endDate = request.WithDefaultDate(endDate, time.Now().Format("2006-01-02"))
	if err := request.ValidateDateRange(startDate, endDate); err != nil {
		return nil, err
	}
	return c.SendQuery("query_performance_express_report", "1", "2000", code, startDate, endDate)
}

func (c *Client) QueryForecastReport(code, startDate, endDate string) (*rs.ResultData, error) {
	code, err := request.NormalizeCode(code)
	if err != nil {
		return nil, err
	}
	startDate = request.WithDefaultDate(startDate, "2015-01-01")
	endDate = request.WithDefaultDate(endDate, time.Now().Format("2006-01-02"))
	if err := request.ValidateDateRange(startDate, endDate); err != nil {
		return nil, err
	}
	return c.SendQuery("query_forecast_report", "1", "2000", code, startDate, endDate)
}
