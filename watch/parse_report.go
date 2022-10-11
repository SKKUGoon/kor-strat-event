package watch

import (
	"golang.org/x/net/html"
	"strings"
)

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
	case strings.Contains(data, rightIssueBefVol0):
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
	case caseNum == 2:
		c.Ratio = d
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
	case strings.Contains(data, cBondConvertRatio):
		if !strings.Contains(data, "(D=(A+B)/C)") {
			return true, cBondConvertRatioP, 2
		} else {
			return false, -1, 0
		}
	default:
		return false, -1, 0
	}
}

func exBondIssueFillin(d string, caseNum int, c *ExchangeIssue) ExchangeIssue {
	switch {
	case caseNum == 1:
		c.Ratio = d
		return *c
	case caseNum == 2:
		c.Rate = d
		return *c
	default:
		return *c
	}
}

func exBondIssueSort(t html.Token) (bool, int, int) {
	data := strings.ReplaceAll(t.Data, "\u00a0", "")
	data = strings.ReplaceAll(data, " ", "")

	switch {
	case strings.Contains(data, cExcBondRatio):
		return true, cExcBondRatioP, 1
	case strings.Contains(data, cExcBondRate):
		return true, cExcBondRateP, 2
	default:
		return false, -1, 0
	}
}
