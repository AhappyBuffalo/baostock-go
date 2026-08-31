package protocol

import (
	"encoding/json"
	"errors"
	"strings"
)

type ResponseData struct {
	MsgType       string
	MsgBodyLength string

	ErrorCode string
	ErrorMsg  string

	Method       string
	UserID       string
	CurPageNum   string
	PerPageCount string

	Fields     []string
	Data       [][]string

	Code      string
	CodeName  string
	StartDate string
	EndDate   string
	Frequency string
	AdjustFlag string
	Year      string
	YearType  string
	Quarter   string
	Day       string
	Date      string
}

func parseMessage(message string) (*ResponseData, error) {
	if len(message) < messageHeaderLength {
		return nil, errors.New("response message too short")
	}

	header := message[:messageHeaderLength]
	body := strings.TrimSuffix(message[messageHeaderLength:], "\n")

	headerArr := strings.Split(header, messageSplit)
	if len(headerArr) != 3 {
		return nil, errors.New("invalid message header")
	}

	bodyArr := strings.Split(body, messageSplit)
	if len(bodyArr) < 2 {
		return nil, errors.New("invalid message body")
	}

	rd := &ResponseData{
		MsgType:       headerArr[1],
		MsgBodyLength: headerArr[2],
		ErrorCode:     bodyArr[0],
		ErrorMsg:      bodyArr[1],
	}

	if rd.ErrorCode == "0" {
		if len(bodyArr) > 3 {
			rd.Method = bodyArr[2]
			rd.UserID = bodyArr[3]
		}
		if len(bodyArr) > 6 {
			rd.CurPageNum = bodyArr[4]
			rd.PerPageCount = bodyArr[5]
			rd.Data = parseRecordJSON(bodyArr[6])
			if len(bodyArr) > 7 {
				assignQueryMeta(rd, bodyArr)
			}
		} else if len(bodyArr) > 4 {
			rd.Data = parseRecordJSON(bodyArr[4])
			if len(bodyArr) > 5 {
				rd.Fields = parseFields(bodyArr[5])
			}
		}
	}

	return rd, nil
}

func assignQueryMeta(rd *ResponseData, bodyArr []string) {
	switch rd.Method {
	case "query_history_k_data_plus":
		rd.Code = bodyArr[7]
		if len(bodyArr) > 8 {
			rd.Fields = parseFields(bodyArr[8])
		}
		if len(bodyArr) > 9 {
			rd.StartDate = bodyArr[9]
		}
		if len(bodyArr) > 10 {
			rd.EndDate = bodyArr[10]
		}
		if len(bodyArr) > 11 {
			rd.Frequency = bodyArr[11]
		}
		if len(bodyArr) > 12 {
			rd.AdjustFlag = bodyArr[12]
		}
	case "query_trade_dates", "query_deposit_rate_data", "query_loan_rate_data", "query_required_reserve_ratio_data", "query_money_supply_data_month", "query_money_supply_data_year", "query_performance_express_report", "query_forecast_report", "query_adjust_factor":
		if len(bodyArr) > 7 {
			rd.StartDate = bodyArr[7]
		}
		if len(bodyArr) > 8 {
			rd.EndDate = bodyArr[8]
		}
		if len(bodyArr) > 9 {
			rd.Fields = parseFields(bodyArr[9])
		}
	case "query_all_stock":
		if len(bodyArr) > 7 {
			rd.Day = bodyArr[7]
		}
		if len(bodyArr) > 8 {
			rd.Fields = parseFields(bodyArr[8])
		}
	case "query_stock_basic":
		if len(bodyArr) > 7 {
			rd.Code = bodyArr[7]
		}
		if len(bodyArr) > 8 {
			rd.CodeName = bodyArr[8]
		}
		if len(bodyArr) > 9 {
			rd.Fields = parseFields(bodyArr[9])
		}
	case "query_dividend_data":
		if len(bodyArr) > 7 {
			rd.Code = bodyArr[7]
		}
		if len(bodyArr) > 8 {
			rd.Year = bodyArr[8]
		}
		if len(bodyArr) > 9 {
			rd.YearType = bodyArr[9]
		}
		if len(bodyArr) > 10 {
			rd.Fields = parseFields(bodyArr[10])
		}
	case "query_stock_industry", "query_stock_concept", "query_stock_area":
		if len(bodyArr) > 7 {
			rd.Code = bodyArr[7]
		}
		if len(bodyArr) > 8 {
			rd.Date = bodyArr[8]
		}
		if len(bodyArr) > 9 {
			rd.Fields = parseFields(bodyArr[9])
		}
	default:
		if len(bodyArr) > 7 {
			rd.Date = bodyArr[7]
		}
		if len(bodyArr) > 8 {
			rd.Fields = parseFields(bodyArr[8])
		}
	}
}

func parseRecordJSON(raw string) [][]string {
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return nil
	}
	var payload struct {
		Records [][]string `json:"record"`
	}
	if err := json.Unmarshal([]byte(raw), &payload); err != nil {
		return nil
	}
	return payload.Records
}

func parseFields(raw string) []string {
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return nil
	}
	parts := strings.Split(raw, attributeSplit)
	fields := make([]string, 0, len(parts))
	for _, p := range parts {
		if v := strings.TrimSpace(p); v != "" {
			fields = append(fields, v)
		}
	}
	if len(fields) == 0 {
		return nil
	}
	return fields
}
