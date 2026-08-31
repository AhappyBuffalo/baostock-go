package request

import (
	"errors"
	"strings"
	"time"
)

const stockCodeLength = 9

var (
	ErrEmptyCode     = errors.New("股票代码不能为空，请检查")
	ErrInvalidCode   = errors.New("股票代码应为9位，请检查")
	ErrInvalidDate   = errors.New("日期格式不正确，请修改")
	ErrStartAfterEnd = errors.New("起始日期大于终止日期，请修改")
)

func NormalizeCode(code string) (string, error) {
	code = strings.TrimSpace(code)
	if code == "" {
		return "", ErrEmptyCode
	}
	code = strings.ToLower(code)
	if strings.HasSuffix(code, "sh") || strings.HasSuffix(code, "sz") {
		code = code[7:9] + "." + code[0:6]
	}
	if len(code) != stockCodeLength {
		return "", ErrInvalidCode
	}
	return code, nil
}

func NormalizeOptionalCode(code string) (string, error) {
	code = strings.TrimSpace(code)
	if code == "" {
		return "", nil
	}
	return NormalizeCode(code)
}

func ValidateDate(date string) error {
	if date == "" {
		return nil
	}
	_, err := time.Parse("2006-01-02", date)
	if err != nil {
		return ErrInvalidDate
	}
	return nil
}

func WithDefaultDate(date, defaultValue string) string {
	if strings.TrimSpace(date) == "" {
		return defaultValue
	}
	return date
}

func ValidateDateRange(start, end string) error {
	if start == "" || end == "" {
		return nil
	}
	if err := ValidateDate(start); err != nil {
		return err
	}
	if err := ValidateDate(end); err != nil {
		return err
	}
	startTime, _ := time.Parse("2006-01-02", start)
	endTime, _ := time.Parse("2006-01-02", end)
	if endTime.Before(startTime) {
		return ErrStartAfterEnd
	}
	return nil
}
