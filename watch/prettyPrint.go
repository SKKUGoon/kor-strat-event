package watch

import "fmt"

func PrettyPrintB(b BonusIssue, nw NewReportWatch) {
	fmt.Println("==================================")
	fmt.Println("<   Event Driven %%% 무상증자 공시   >")
	fmt.Printf("* 시간: %s\n", b.ReportDate.Format("2006-01-02 15:04:05"))
	fmt.Printf("* 제목: %s\n", nw.Ctx.Title)
	fmt.Println("* 무상증자 관련 상세 정보")
	fmt.Printf(">>> %s: %s\n", bonusIssueStockPrc, b.StkPrice)
	fmt.Printf(">>> %s: %s\n", bonusIssueStockAdd, b.NewStockDist)
	fmt.Printf(">>> %s: %s\n", bonusIssueLock, b.LockDate)
	fmt.Println("==================================")
}

func PrettyPrintR(r RightsIssue, nw NewReportWatch) {
	fmt.Println("==================================")
	fmt.Println("<   Event Driven %%% 유상증자 공시   >")
	fmt.Printf("* 시간: %s\n", r.ReportDate.Format("2006-01-02 15:04:05"))
	fmt.Printf("* 제목: %s\n", nw.Ctx.Title)
	fmt.Println("* 유상증자 관련 상세 정보")
	fmt.Printf(">>> %s: %s\n", rightIssueBefVol, r.BefTotalVolume)
	fmt.Printf(">>> %s: %s\n", rightIssueAftVol, r.AftTotalVolume)
	fmt.Printf(">>> %s: %s\n", rightIssueStockPrc, r.NewStockPrc)
	fmt.Println("==================================")
}

func PrettyPrintC(c ConvertibleIssue, nw NewReportWatch) {
	fmt.Println("==================================")
	fmt.Println("<   Event Driven %%% 전환사채 공시   >")
	fmt.Printf("* 시간: %s\n", c.ReportDate.Format("2006-01-02 15:04:05"))
	fmt.Printf("* 제목: %s\n", nw.Ctx.Title)
	fmt.Println("* 전환사채 관련 상세 정보")
	fmt.Printf(">>> %s: %s\n", cBondConvertPrc, c.ConvertPrc)
	fmt.Println("==================================")
}
