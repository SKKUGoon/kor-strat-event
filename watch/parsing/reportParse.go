package parsing

import (
	"errors"
	"fmt"
	"goKorMktEvent/watch"
	"golang.org/x/net/html"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

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
	bonusIssueStockAdd = "1주당신주배정주식수"
	bonusIssueStockPrc = "1주당액면가액"
	bonusIssueLock     = "신주배정기준일"
)

func GetReportText(u string) (string, error) {
	resp, err := http.Get(u)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", errors.New(resp.Status)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	} else {
		return string(data), nil
	}
}

func getDcmNo(ctx string) (string, string) {
	i := strings.Index(ctx, dcmId)
	dcmNo := ctx[i+dcmParseIndSrt : i+dcmParseIndSrt+dcmLength]
	return "dcmNo", dcmNo
}

func getRcpNo(ctx string) (string, string) {
	i := strings.Index(ctx, rcpId)
	rcpNo := ctx[i+rcpParseIndSrt : i+rcpParseIndSrt+rcpLength]
	return "rcpNo", rcpNo
}

func RawReportURL(ctx string) (string, error) {
	// Create Parameter HashMap
	dcmKey, dcmVal := getDcmNo(ctx)
	rcpKey, rcpVal := getRcpNo(ctx)
	var rawReportURL = "http://dart.fss.or.kr/report/viewer.do?"
	var rawReportParam = watch.ReportMap{
		dcmKey:   dcmVal,
		rcpKey:   rcpVal,
		"eleId":  "0",
		"offset": "0",
		"length": "0",
		"dtd":    "dart3.xsd",
	}

	// Create Request
	resp, err := http.NewRequest("GET", rawReportURL, nil)
	if err != nil {
		log.Println(err)
		return "", err
	}

	// Add Parameter
	qry := resp.URL.Query()
	for key, element := range rawReportParam {
		qry.Add(key, element)
	}

	resp.URL.RawQuery = qry.Encode()

	return resp.URL.String(), nil
}

func RawReportText(source string) {
	c, _ := GetReportText(source)

	u, _ := RawReportURL(c)
	fmt.Println("url", u)

	resp, _ := http.Get(u)
	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)
	_ = ReportContent(string(data))
}

func ReportContent(text string) watch.BonusIssue {
	var container = watch.BonusIssue{}

	tkn := html.NewTokenizer(strings.NewReader(text))
	for {
		tt := tkn.Next()

		switch {
		// End of parsing
		case tt == html.ErrorToken:
			return container

		// ex) <a> is a StartTagToken
		case tt == html.StartTagToken:
			tn := tkn.Token()
			switch {
			case tn.Data == "td":
				tt, tn = tkn.Next(), tkn.Token()
				ok := bonusIssueFill(tn)
				if ok {
					tt, tn = tkn.Next(), tkn.Token()
					fmt.Println(tn.Data)
				}
			default:
				continue
			}
		}
	}
}

func bonusIssueFill(t html.Token) bool {
	data := strings.ReplaceAll(t.Data, "\u00a0", "") // "\u00a0" is %nbsp
	data = strings.ReplaceAll(data, " ", "")

	switch {
	case strings.Contains(data, bonusIssueStockAdd):
		fmt.Println("case1", data)
		return true
	case strings.Contains(data, bonusIssueStockPrc):
		fmt.Println("case2", data)
		return true
	case strings.Contains(data, bonusIssueLock):
		fmt.Println("case3", data)
		return true
	default:
		return false
	}
}
