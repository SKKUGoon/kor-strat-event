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
	sRedundant = "기재정정"
)

const (
	// title of html <td> token
	// bonus issue (무상증자). typeNum 1
	bonusIssueStockAdd = "1주당신주배정주식수"
	bonusIssueStockPrc = "1주당액면가액"
	bonusIssueLock     = "신주배정기준일"

	// right issue (유상증자). typeNum 2
	rightIssueBefVol   = "발행주식총수(주)"
	rightIssueAftVol   = "신주의종류와수"
	rightIssueStockPrc = "신주발행가액"

	// convertible bond issue (전환사채발행). typeNum 4
	cBondConvertPrc = "전환가액(원"

	// how many time we should pass after we get a desired information
	bonusIssueStockAddP = 1
	bonusIssueStockPrcP = 0
	bonusIssueLockP     = 0

	rightIssueBefVolP = 1
	rightIssueAftVolP = 1
	rightIssueLockP   = 0

	cBondConvertPrcP = 0
)
