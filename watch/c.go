package watch

const (
	// dcmNo parsing info
	dcmId     = "'dcmNo'"
	dcmLength = 7

	// rcpNo parsing info
	rcpId     = "'rcpNo'"
	rcpLength = 14

	dcmParseIndSrt = len(dcmId) + len("] = '")
	rcpParseIndSrt = len(rcpId) + len("] = '")
)

const (
	sKOSPI  = "유가"
	sKOSDAQ = "코스닥"

	sMusang    = "주요사항보고서(무상증자결정)"
	sYusang    = "주요사항보고서(유상증자결정)"
	sYuMusang  = "주요사항보고서(유무상증자결정)"
	sConvert   = "주요사항보고서(전환사채권발행결정)"
	sNewStock  = "주요사항보고서(신주인수권부사채권발행결정)"
	sExcStock  = "주요사항보고서(교환사채권발행결정)"
	sRedundant = "기재정정"
)

const (
	// title of html <td> token
	// bonus issue (무상증자). typeNum 1
	bonusIssueStockAdd = "1주당신주배정주식수"
	bonusIssueStockPrc = "1주당액면가액"
	bonusIssueLock     = "신주배정기준일"

	// right issue (유상증자). typeNum 2
	rightIssueBefVol   = "발행주식총수"
	rightIssueBefVol0  = "증자전"
	rightIssueAftVol   = "신주의종류와수"
	rightIssueStockPrc = "신주발행가액"

	// convertible bond issue (전환사채발행). typeNum 4
	cBondConvertPrc   = "전환가액(원"
	cBondConvertRatio = "주식총수대비"

	// how many time we should pass after we get a desired information
	bonusIssueStockAddP = 1
	bonusIssueStockPrcP = 0
	bonusIssueLockP     = 0

	rightIssueBefVolP = 1
	rightIssueAftVolP = 1
	rightIssueLockP   = 1

	cBondConvertPrcP   = 0
	cBondConvertRatioP = 0
)

var (
	TestStruct0 = NewReportWatch{
		Ctx: Content{
			Title:    "무상증자 테스트용",
			Link:     "https://dart.fss.or.kr/dsaf001/main.do?rcpNo=20220525000375",
			Category: "TEST",
		},
		T: 1,
	}
	TestStruct1 = NewReportWatch{
		Ctx: Content{
			Title:    "유상증자 테스트용",
			Link:     "https://dart.fss.or.kr/dsaf001/main.do?rcpNo=20220525000311",
			Category: "TEST",
		},
		T: 2,
	}

	TestStruct2 = NewReportWatch{
		Ctx: Content{
			Title:    "전환사채 테스트용",
			Link:     "https://dart.fss.or.kr/dsaf001/main.do?rcpNo=20220525000128",
			Category: "TEST",
		},
		T: 4,
	}
)
