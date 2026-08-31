package main

import (
	"fmt"
	"log"

	baostock "github.com/example/baostock-go/baostock"
)

func main() {
	client := baostock.NewClient()
	if err := client.Login(); err != nil {
		log.Fatal(err)
	}
	defer client.Logout()

	// 示例1：查询历史K线
	rs, err := client.QueryHistoryKDataPlus("sh.600000", "date,open,high,low,close,volume", "2024-01-01", "2024-01-10", "d", "3")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("=== 历史K线 ===")
	fmt.Println("fields:", rs.Fields)
	for rs.Next() {
		fmt.Println(rs.GetRowData())
	}

	// 示例2：查询交易日历
	rs2, err := client.QueryTradeDates("2024-01-01", "2024-01-10")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("\n=== 交易日历 ===")
	fmt.Println("fields:", rs2.Fields)
	for rs2.Next() {
		fmt.Println(rs2.GetRowData())
	}
}
