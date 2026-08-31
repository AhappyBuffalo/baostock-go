package result

type ResultData struct {
	MsgType       string
	MsgBodyLength string

	ErrorCode string
	ErrorMsg  string

	Method       string
	UserID       string
	CurPageNum   string
	PerPageCount string

	Fields  []string
	Data    [][]string
	Index   int

	Code       string
	CodeName   string
	StartDate  string
	EndDate    string
	Frequency  string
	AdjustFlag string
	Year       string
	YearType   string
	Quarter    string
	Day        string
	Date       string
}

func (r *ResultData) Next() bool {
	if r == nil || len(r.Data) == 0 {
		return false
	}
	return r.Index < len(r.Data)
}

func (r *ResultData) GetRowData() []string {
	if r == nil || r.Index >= len(r.Data) {
		return nil
	}
	row := r.Data[r.Index]
	r.Index++
	return row
}

func (r *ResultData) ResetIndex() {
	if r == nil {
		return
	}
	r.Index = 0
}
