package watch

import "time"

type ReportMap map[string]string

type BonusIssue struct {
	Company      string
	StkPrice     string
	NewStockDist string // distribute n stocks for every 1 stock. 1주당 n개 배분
	ReportDate   time.Time
	LockDate     string // snapShot. 권리락일
}

type RightsIssue struct {
	Company        string
	NewStockPrc    string // 신주발행가액
	BefTotalVolume string // 유상증자 전 총 주식 수
	AftTotalVolume string // 신주의 수
	ReportDate     time.Time
}

type ConvertibleIssue struct {
	Company    string
	ReportDate time.Time
	ConvertPrc string
}

type NewReportWatch struct {
	Ctx     Content
	RawUrl  string // url to raw report
	T       int    // report type
	OrgHTML string // original HTML in string
	RAWHTML string // inner HTML in string
}
