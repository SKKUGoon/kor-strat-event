package watch

import "fmt"

func PrettyPrintB(b BonusIssue, nw NewReportWatch) {
	fmt.Println("======================================")
	fmt.Println("<   Event Driven %%% 무상증자 공시   >")
	fmt.Printf("* 시간: %s\n", b.ReportDate.Format("2006-01-02 15:04:05"))
	fmt.Printf("* 제목: %s\n", nw.Ctx.Title)
	fmt.Printf("* 보고서원문: %s\n", nw.RawUrl)
	fmt.Println("-----")
	fmt.Println("* 무상증자 관련 상세 정보")
	fmt.Printf(">>> %s: %s\n", bonusIssueStockPrc, b.StkPrice)
	fmt.Printf(">>> %s: %s(주 / 1주)\n", bonusIssueStockAdd, b.NewStockDist)
	fmt.Printf(">>> %s: %s\n", bonusIssueLock, b.LockDate)
	fmt.Println("======================================")
	fmt.Println()
}

func PrettyPrintR(r RightsIssue, nw NewReportWatch) {
	fmt.Println("======================================")
	fmt.Println("<   Event Driven %%% 유상증자 공시   >")
	fmt.Printf("* 시간: %s\n", r.ReportDate.Format("2006-01-02 15:04:05"))
	fmt.Printf("* 제목: %s\n", nw.Ctx.Title)
	fmt.Printf("* 보고서원문: %s\n", nw.RawUrl)
	fmt.Println("-----")
	fmt.Println("* 유상증자 관련 상세 정보")
	fmt.Printf(">>> %s: %s(주)\n", rightIssueBefVol, r.BefTotalVolume)
	fmt.Printf(">>> %s: %s(주)\n", rightIssueAftVol, r.AftTotalVolume)
	fmt.Printf(">>> %s(원): %s\n", rightIssueStockPrc, r.NewStockPrc)
	fmt.Println("======================================")
	fmt.Println()
}

func PrettyPrintC(c ConvertibleIssue, nw NewReportWatch) {
	fmt.Println("======================================")
	fmt.Println("<   Event Driven %%% 전환사채 공시   >")
	fmt.Printf("* 시간: %s\n", c.ReportDate.Format("2006-01-02 15:04:05"))
	fmt.Printf("* 제목: %s\n", nw.Ctx.Title)
	fmt.Printf("* 보고서원문: %s\n", nw.RawUrl)
	fmt.Println("-----")
	fmt.Println("* 전환사채 관련 상세 정보")
	fmt.Printf(">>> %s): %s\n", cBondConvertPrc, c.ConvertPrc)
	fmt.Printf(">>> %s: %s\n", cBondConvertRatio, c.Ratio)
	fmt.Println("======================================")
	fmt.Println()
}
