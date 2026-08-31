package baostock

import (
	"github.com/example/baostock-go/internal/request"
	rs "github.com/example/baostock-go/internal/result"
)

func (c *Client) QueryDepositRateData(startDate, endDate string) (*rs.ResultData, error) {
	if err := request.ValidateDateRange(startDate, endDate); err != nil {
		return nil, err
	}
	return c.SendQuery("query_deposit_rate_data", "1", "2000", startDate, endDate)
}

func (c *Client) QueryLoanRateData(startDate, endDate string) (*rs.ResultData, error) {
	if err := request.ValidateDateRange(startDate, endDate); err != nil {
		return nil, err
	}
	return c.SendQuery("query_loan_rate_data", "1", "2000", startDate, endDate)
}

func (c *Client) QueryRequiredReserveRatioData(startDate, endDate, yearType string) (*rs.ResultData, error) {
	if err := request.ValidateDateRange(startDate, endDate); err != nil {
		return nil, err
	}
	if yearType == "" {
		yearType = "0"
	}
	return c.SendQuery("query_required_reserve_ratio_data", "1", "2000", startDate, endDate, yearType)
}

func (c *Client) QueryMoneySupplyDataMonth(startDate, endDate string) (*rs.ResultData, error) {
	if err := request.ValidateDateRange(startDate, endDate); err != nil {
		return nil, err
	}
	return c.SendQuery("query_money_supply_data_month", "1", "2000", startDate, endDate)
}

func (c *Client) QueryMoneySupplyDataYear(startDate, endDate string) (*rs.ResultData, error) {
	if err := request.ValidateDateRange(startDate, endDate); err != nil {
		return nil, err
	}
	return c.SendQuery("query_money_supply_data_year", "1", "2000", startDate, endDate)
}

func (c *Client) QueryCPIData(startDate, endDate string) (*rs.ResultData, error) {
	if err := request.ValidateDateRange(startDate, endDate); err != nil {
		return nil, err
	}
	return c.SendQuery("query_cpi_data", "1", "2000", startDate, endDate)
}

func (c *Client) QueryPPIData(startDate, endDate string) (*rs.ResultData, error) {
	if err := request.ValidateDateRange(startDate, endDate); err != nil {
		return nil, err
	}
	return c.SendQuery("query_ppi_data", "1", "2000", startDate, endDate)
}

func (c *Client) QueryPMIData(startDate, endDate string) (*rs.ResultData, error) {
	if err := request.ValidateDateRange(startDate, endDate); err != nil {
		return nil, err
	}
	return c.SendQuery("query_pmi_data", "1", "2000", startDate, endDate)
}
