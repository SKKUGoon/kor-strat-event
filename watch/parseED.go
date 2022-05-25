package watch

import (
	"errors"
	"golang.org/x/net/html"
	"strings"
	"time"
)

func (dartReport *NewReportWatch) parseEventDriven() (interface{}, error) {
	var (
		// load container
		infoBonus = BonusIssue{}
		infoRight = RightsIssue{}
		infoCB    = ConvertibleIssue{}
	)

	var (
		// start parsing?
		ok = false

		// how much waiting loop?, if parsed(>0) what data is this
		wt, c = -1, 0
	)

	tkn := html.NewTokenizer(strings.NewReader(dartReport.RAWHTML))

	for {
		tt := tkn.Next()

		switch {
		// End of parsing
		case tt == html.ErrorToken:
			switch {
			case dartReport.T == 1:
				infoBonus.ReportDate = time.Now()
				return infoBonus, nil
			case dartReport.T == 2:
				infoRight.ReportDate = time.Now()
				return infoRight, nil
			case dartReport.T == 4:
				infoCB.ReportDate = time.Now()
				return infoCB, nil
			default:
				return nil, errors.New("dart report type not supported")
			}

		// ex) <a> is a StartTagToken
		case tt == html.StartTagToken:
			tn := tkn.Token()
			switch {
			case tn.Data == "td":
				tt, tn = tkn.Next(), tkn.Token()
				/*
						ok : if this <td> block is a point to start parsing?
						wt : how much loop time we have to wait for us to start parsing?

					ex)
					  <TD width='196' height='50' rowspan='2'>5. 1주당 신주배정 주식수</TD>
					  <TD width='115' height='20' align='CENTER'>보통주식 (주)&nbsp;</TD>
					  <TD width='289' height='20' align='RIGHT'>1.0000000</TD>

					1주당 신주배정 주식수가 <td> 에서 걸러지면
					2번 루프 뒤에있는 1.00 이 나와야함.
				*/
				if ok {
					if wt != 0 {
						wt -= 1
						continue
					}
					wt = -1
					switch {
					case dartReport.T == 1:
						infoBonus = bonusIssueFillin(tn.Data, c, &infoBonus)
					case dartReport.T == 2:
						infoRight = rightIssueFillin(tn.Data, c, &infoRight)
					case dartReport.T == 4:
						infoCB = cbondIssueFillin(tn.Data, c, &infoCB)
					}

				}

				// sort out whether this is a desired <td> block
				switch {
				case dartReport.T == 1:
					ok, wt, c = bonusIssueSort(tn)
				case dartReport.T == 2:
					ok, wt, c = rightIssueSort(tn)
				case dartReport.T == 4:
					ok, wt, c = cbondIssueSort(tn)
				}

			default:
				continue
			}
		}
	}
}

func bonusIssueFillin(d string, caseNum int, c *BonusIssue) BonusIssue {
	switch {
	case caseNum == 1:
		c.NewStockDist = d
		return *c
	case caseNum == 2:
		c.StkPrice = d
		return *c
	case caseNum == 3:
		c.LockDate = d
		return *c
	default:
		return *c
	}
}

func bonusIssueSort(t html.Token) (bool, int, int) {
	data := strings.ReplaceAll(t.Data, "\u00a0", "") // "\u00a0" is %nbsp
	data = strings.ReplaceAll(data, " ", "")

	switch {
	case strings.Contains(data, bonusIssueStockAdd):
		return true, bonusIssueStockAddP, 1
	case strings.Contains(data, bonusIssueStockPrc):
		return true, bonusIssueStockPrcP, 2
	case strings.Contains(data, bonusIssueLock):
		return true, bonusIssueLockP, 3
	default:
		return false, -1, 0
	}
}

func rightIssueFillin(d string, caseNum int, c *RightsIssue) RightsIssue {
	switch {
	case caseNum == 1:
		c.BefTotalVolume = d
		return *c
	case caseNum == 2:
		c.NewStockPrc = d
		return *c
	case caseNum == 3:
		c.AftTotalVolume = d
		return *c
	default:
		return *c
	}
}

func rightIssueSort(t html.Token) (bool, int, int) {
	data := strings.ReplaceAll(t.Data, "\u00a0", "")
	data = strings.ReplaceAll(data, " ", "")

	switch {
	case strings.Contains(data, rightIssueBefVol):
		return true, rightIssueBefVolP, 1
	case strings.Contains(data, rightIssueStockPrc):
		return true, rightIssueAftVolP, 2
	case strings.Contains(data, rightIssueAftVol):
		return true, rightIssueLockP, 3
	default:
		return false, -1, 0
	}
}

func cbondIssueFillin(d string, caseNum int, c *ConvertibleIssue) ConvertibleIssue {
	switch {
	case caseNum == 1:
		c.ConvertPrc = d
		return *c
	default:
		return *c
	}
}

func cbondIssueSort(t html.Token) (bool, int, int) {
	data := strings.ReplaceAll(t.Data, "\u00a0", "")
	data = strings.ReplaceAll(data, " ", "")

	switch {
	case strings.Contains(data, cBondConvertPrc):
		return true, cBondConvertPrcP, 1
	default:
		return false, -1, 0
	}
}
